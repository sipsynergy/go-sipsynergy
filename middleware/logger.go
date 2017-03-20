package middleware

import (
    "fmt"
    "gopkg.in/gin-gonic/gin.v1"
)

func Logger() gin.HandlerFunc {
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