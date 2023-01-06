package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kvaliant/online-memo/config"
	"github.com/kvaliant/online-memo/models"
)

func ValidateAccess(c *gin.Context, memo models.Memo) {
	var user models.User
	if err := config.DB.First(&user, c.GetInt("user_id")); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Invalid"})
		return 
	}
	
	if memo.UserID != user.ID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return 
	}

	return
}