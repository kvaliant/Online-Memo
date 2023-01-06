package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kvaliant/online-memo/config"
	"github.com/kvaliant/online-memo/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context){
	type UserInput struct {
		models.User
		InputPassword string `json:"password" binding:"required,min=8"`
	 }

	var userInput = &UserInput{}
	if err := c.Bind(userInput); err != nil {
		resp := map[string]string{"message": "Malformed input"}
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.InputPassword), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)


	var user = &models.User{}
	user.Username = userInput.Username
	user.FullName = userInput.FullName
	user.Password = userInput.Password

	if err := config.DB.Create(&user).Error; err != nil {
		resp := map[string]string{"message": err.Error()}
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp)
		return
	}

	c.JSON(http.StatusOK, gin.H{"user" : user})
}