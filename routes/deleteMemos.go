package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kvaliant/online-memo/config"
	"github.com/kvaliant/online-memo/models"
	"gorm.io/gorm"
)

func DeleteMemo(c *gin.Context){
	id := c.Param("memoid")
	result := config.DB.Delete(&models.Memo{}, id)
	if result.Error != nil {
		switch result.Error {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Record not found"})
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Invalid"})
		}
		return
	}
		
	c.JSON(http.StatusOK, gin.H{"message": "Delete successful"})
}