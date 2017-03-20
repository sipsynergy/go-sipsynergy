package storage

import "github.com/jinzhu/gorm"

// InitDB creates and migrates the database
func InitDB(dbDialect string, dbDsn string, logMode bool) (*gorm.DB, error) {
	db, err := gorm.Open(dbDialect, dbDsn)
	if err != nil {
		return nil, err
	}

	if logMode == true {
		db.LogMode(true)
	}

	return db, nil
}
