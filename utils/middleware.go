package utils

import (
    "fmt"

    "gopkg.in/gin-gonic/gin.v1"
    "net/http"
)

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*.sipsynergy.co.uk")
        c.Writer.Header().Set("Access-Control-Max-Age", "86400")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PATCH, DELETE, UPDATE")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
        c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

        if c.Request.Method == "OPTIONS" {
            fmt.Println("OPTIONS")

            c.JSON(http.StatusInternalServerError, gin.H{
                "message": "Cores error. Option not allowed.",
            })

            return

        } else {
            c.Next()
        }
    }
}

func HeaderKeyMiddleware(key string) gin.HandlerFunc {
    return func (c *gin.Context) {
        if c.Request.Header.Get("X-API-KEY") != key {

            c.JSON(http.StatusUnauthorized, gin.H{
                "message": "Not Authorized",
            })

            return
        } else {
            c.Next()
        }
    }
}

func LoggerMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Header("Content-Type", "application/json")
        c.Next()
        errors := c.Errors.ByType(gin.ErrorTypeAny)
        if len(errors) > 0 {
            fmt.Println(errors)
            c.JSON(-1, errors.JSON())
            return
        }
    }
}