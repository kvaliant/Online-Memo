package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kvaliant/online-memo/config"
	"github.com/kvaliant/online-memo/models"
)

func GetMemo(c *gin.Context){
	id := c.Param("memoid")

	var memo models.Memo
	config.DB.Find(&memo, id)

	c.JSON(200, gin.H{
		"memos": memo,
	})
}
