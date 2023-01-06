package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kvaliant/online-memo/config"
	"github.com/kvaliant/online-memo/models"
)

func PostMemo(c *gin.Context){
	var userInput struct {
		Title	string	`json:"title" gorm:"required"`
		Body	string	`json:"body" gorm:"required"`
	}

	if err := c.Bind(&userInput); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	
	var memo = &models.Memo{}
	memo.Title = userInput.Title
	memo.Body = userInput.Body

	user := &models.User{}
	config.DB.Find(&user, c.GetInt("user_id"))
	memo.User = *user
	
	result := config.DB.Create(&memo)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": result.Error})
		return
	}

	c.JSON(200, gin.H{
		"memo": memo,
	})
}
