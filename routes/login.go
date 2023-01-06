package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kvaliant/online-memo/config"
	"github.com/kvaliant/online-memo/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context){
	var userInput struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.BindJSON(&userInput); err != nil {
		resp := map[string]string{"message": "Unauthorized"}
		c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
		return
	}

	var user models.User
	result := config.DB.Find(&user, "username = ?", userInput.Username)
	if result.Error != nil {
		resp := map[string]string{"message": "Username not found"}
		c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		resp := map[string]string{"message": "Password invalid"}
		c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
		return
	}

	expTime := time.Now().Add(time.Minute * 1)
	claims := &config.JWTClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "online-memo",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		resp := map[string]string{"message": err.Error()}
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}

	c.SetCookie(
		"token",
		token,
		3600,
		"/",
		"localhost",
		false,
		true,
	)
	resp := map[string]string{"message": "Login successful"}
	c.JSON(http.StatusOK, resp)
}