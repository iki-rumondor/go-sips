package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/go-monev/internal/http/response"
	"github.com/iki-rumondor/go-monev/internal/utils"
)

func IsRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		mc := c.MustGet("map_claims")
		mapClaims := mc.(jwt.MapClaims)

		roleJwt := mapClaims["role"].(string)
		if roleJwt != role {
			utils.HandleError(c, response.UNAUTH_ERR("Hak Akses Dibatasi"))
			return
		}
		c.Next()
	}
}
