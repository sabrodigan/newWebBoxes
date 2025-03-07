package middleware

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sabrodigan/newWebBoxes/server/config"
)

func ErrorHandling() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) > 0 {
			ctx.JSON(400, gin.H{"error": ctx.Errors})
			err := ctx.Errors[0]
			errMap := make(map[string]any)
			if err != nil {
				splitErr := strings.Split(err.Error(), "::")
				isDebugEnabled, _ := config.GetEnvProperty("debug")
				isDebugEnabled = strings.TrimSpace(strings.ToLower(isDebugEnabled))
				if isDebugEnabled == "yes" && len(splitErr) > 0 {
					fmt.Printf("Method name: %s, Status Code:%s, Client Message:%s, Error:%v\n", splitErr[1], splitErr[0], splitErr[2], splitErr[3])
				}
				statusCode, err := strconv.Atoi(splitErr[0])
				if err != nil {
					statusCode = 500
				}

				errMap["status"] = statusCode
				errMap["message"] = splitErr[1]
				errMap["error"] = splitErr[2]
				ctx.JSON(statusCode, errMap)
			}
		}
	}
}
