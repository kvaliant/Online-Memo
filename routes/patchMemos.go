package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kvaliant/online-memo/config"
	"github.com/kvaliant/online-memo/helper"
	"github.com/kvaliant/online-memo/models"
)

func PatchMemo(c *gin.Context){
	id := c.Param("memoid")
	var body = &models.Memo{}
	c.Bind(body)
	
	var memo models.Memo
	config.DB.First(&memo, id)

	helper.ValidateAccess(c, memo)

	config.DB.Model(&memo).Updates(models.Memo{
		Title: body.Title,
		Body: body.Body,
	})

	c.JSON(200, gin.H{
		"memos": memo,
	})
}
