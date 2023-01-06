package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kvaliant/online-memo/config"
	"github.com/kvaliant/online-memo/models"
)

func GetMemos(c *gin.Context){
	var memos []models.Memo
	result := config.DB.Where("user_id = ?", c.GetInt("user_id")).Find(&memos)
	if result.Error != nil {
		c.Status(400)
		return
	}
	
	c.JSON(200, gin.H{
		"memos": memos,
	})
}