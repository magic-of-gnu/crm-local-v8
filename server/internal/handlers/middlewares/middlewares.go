package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/services"
)

type authMiddleware struct {
	tokensService services.TokensService
}

type AuthMiddleware interface {
	IsAuthorized() gin.HandlerFunc
}

func NewAuthMiddleware(
	tokensService services.TokensService,
) *authMiddleware {
	return &authMiddleware{
		tokensService: tokensService,
	}
}

// type TokenStruct struct {
// 	token string `json:"token"`
// }

func (mm *authMiddleware) IsAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {

		// before
		authHeader := c.GetHeader("Authorization")
		fmt.Println("authToken: ", authHeader)
		// token := strings.Split(authHeader, "")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			fmt.Println("broken authHeader, authHeader: ", authHeader)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
			return
		}

		splitted := strings.Split(authHeader, " ")
		if len(splitted) != 2 {
			fmt.Println("broken authHeader, authHeader: ", authHeader)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
			return
		}

		jwt_token_string := splitted[1]

		_, err := mm.tokensService.ParseSignedToken(jwt_token_string)
		if err != nil {
			fmt.Println("broken token, token: ", jwt_token_string, " err: ", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
			return
		}

		c.Next()

	}
}
