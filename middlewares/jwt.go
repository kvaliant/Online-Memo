package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kvaliant/online-memo/config"
	"github.com/kvaliant/online-memo/models"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				resp := map[string]string{"message": "Unauthorized"}
				c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
				return
			}
		}

		tokenString := cookie.Value

		claims := &config.JWTClaim{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error){
			return config.JWT_KEY, nil
		})
		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				resp := map[string]string{"message": "Unauthorized"}
				c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			case jwt.ValidationErrorExpired:
				resp := map[string]string{"message": "Unauthorized, Token expored!"}
				c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			default:
				resp := map[string]string{"message": "Unauthorized"}
				c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			}
			return
		}

		if !token.Valid {
			resp := map[string]string{"message": "Unauthorized"}
			c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			return
		}
		user := &models.User{}
		result := config.DB.First(user, "username = ?", claims.Username)
		if result.Error != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}

		c.Set("user_id", user.ID)
		c.Next()

	}
}