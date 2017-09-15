package auth

import (
	"context"
	"database/sql"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Jeffail/gabs"
	"github.com/buger/jsonparser"
	"github.com/micro/go-micro/metadata"
	"github.com/parnurzeal/gorequest"
	"github.com/sipsynergy/proto-go/accounts_users"
	"github.com/sipsynergy/proto-go/stypes/common"
)

// IntrospectToken reads the 'subject' from the 'session_data' for this token
// via direct communication with the hydra db.
// In our case the subject is the humanID of the authenticated user.
// If the token has expired an error will be returned.
func IntrospectToken(token string, db *sql.DB) (subject string, expires time.Time, err error) {
	segments := strings.Split(token, "."); if len(segments) != 2 {
		err = errors.New("invalid hydra token")
		return
	}

	signature := segments[1]

	row := db.QueryRow("SELECT session_data FROM hydra_oauth2_access WHERE signature = ?", signature); if row == nil {
		err = errors.New("unable to build token query")
		return
	}

	var d string

	scanErr := row.Scan(&d); if scanErr != nil {
		err = errors.New("unable to query token: " + scanErr.Error())
		return
	}

	expiresAt, expiryValueErr := jsonparser.GetString([]byte(d), "idToken", "Claims", "ExpiresAt"); if expiryValueErr != nil {
		err = errors.New("unable to parse token expiry from session_data: " + expiryValueErr.Error())
		return
	}

	expires, expiryParseErr := time.Parse("2006-01-02T15:04:05.999999999Z", expiresAt); if expiryParseErr != nil {
		err = errors.New("unable to parse token expiry time: " + expiresAt + " " + expiryParseErr.Error())
		return
	}

	if expires.Before(time.Now()) {
		err = errors.New("token has expired")
		return
	}

	subject, subjectParseErr := jsonparser.GetString([]byte(d), "idToken", "Subject"); if subjectParseErr != nil {
		err = errors.New("unable to parse userID from session_data: " + subjectParseErr.Error())
		return
	}

	return subject, expires, nil
}

// GetUserFromToken will use both current (hydra and accounts ms) and
// legacy (monolith) methods of returning the user from an access token.
func GetUserFromToken(
	token string,
	legacyAPIRootURL string,
	db *sql.DB,
	uc accounts_users.UsersClient,
) (
	user *accounts_users.User,
	expires time.Time,
	err error,
) {
	user, expires, err = GetUserFromHydraToken(token, uc, db); if err == nil {
		return
	}

	user, expires, err = GetUserFromLegacyToken(token, legacyAPIRootURL, uc); if err == nil {
		return
	}

	return
}

// GetUserFromToken contacts the legacy api (monolith) to determine the
// authenticated user.
func GetUserFromLegacyToken(
	token string,
	legacyAPIRootURL string,
	uc accounts_users.UsersClient,
) (
	user *accounts_users.User,
	expires time.Time,
	err error,
) {
	resp, body, errs := gorequest.New().
		Get(legacyAPIRootURL).
		Set("Authorization", "Bearer " + token).
		End()

	if len(errs) > 0 {
		return nil, time.Time{}, errors.New("failed to retrieve user from legacy api: " + errs[0].Error())
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, time.Time{}, errors.New(
			"query failed with status: [" + strconv.Itoa(resp.StatusCode) + "] " +
				resp.Status + " body: " + string(body),
		)
	}

	parsedJson, err := gabs.ParseJSON([]byte(body)); if err != nil {
		return nil, time.Time{}, errors.New("failed to parse v1 response json")
	}

	humanIdData := parsedJson.Path("user.humanId").Data()
	if humanIdData == nil {
		return nil, time.Time{}, errors.New("failed to parse v1 response user.humanId field")
	}

	req := common.SimpleEntityRequest{
		EntityID: humanIdData.(string),
	}

	rsp, err := uc.Get(context.Background(), &req); if err != nil {
		return nil, time.Time{}, err
	}

	expiresOn := time.Now().Add(time.Duration(time.Hour))

	return rsp.User, expiresOn, nil
}

// GetUserFromToken uses the current method (hydra and accounts service) to
// determine the authenticated user.
func GetUserFromHydraToken(
	token string,
	uc accounts_users.UsersClient,
	db *sql.DB,
) (
	user *accounts_users.User,
	expiresOn time.Time,
	err error,
) {
	userID, expiresOn, err := IntrospectToken(token, db); if err != nil {
		return
	}

	req := common.SimpleEntityRequest{EntityID: userID}
	rsp, err := uc.Get(context.Background(), &req); if err != nil {
		return
	}

	return rsp.User, expiresOn, nil
}


// GetAccessTokenFromRequest returns the bearer access token from the
// 'Authorization' http header.
func GetAccessTokenFromRequest(r *http.Request) (accessToken string, ok bool) {
	a, ok := r.Header["Authorization"]; if !ok || len(a) < 1 {
		return "", false
	}

	authSegments := strings.Split(a[0], " "); if len(authSegments) != 2 || authSegments[0] != "Bearer" {
		return "", false
	}

	return authSegments[1], true
}

// GetAccessTokenFromContext returns the access token by parsing the
// 'Authorization' metadata from the context.
func GetAccessTokenFromContext(ctx context.Context) (apiKey string, ok bool) {
	meta, ok := metadata.FromContext(ctx); if !ok {
		return "", false
	}

	a, ok := meta["X-Api-Key"]; if !ok || len(a) < 1 {
		return "", false
	}

	return a, true
}

// GetAPIKeyFromRequest returns the value of the 'X-Api-Key' request header.
func GetAPIKeyFromRequest(r *http.Request) (apiKey string, ok bool) {
	key, ok := r.Header["X-Api-Key"]; if !ok || len(key) != 1 {
		return "", false
	}

	return key[0], true
}

// GetAPIKeyFromContext returns the value of the 'Authorization' header stored
// in the micro context.
func GetAPIKeyFromContext(ctx context.Context) (apiKey string, ok bool) {
	meta, ok := metadata.FromContext(ctx); if !ok {
		return "", false
	}

	a, ok := meta["Authorization"]; if !ok || len(a) < 1 {
		return "", false
	}

	authSegments := strings.Split(a, " "); if len(authSegments) != 2 || authSegments[0] != "Bearer" {
		return "", false
	}

	return authSegments[1], true
}
