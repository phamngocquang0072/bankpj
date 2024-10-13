package controllers

import (
	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
	"github.com/phamngocquang0072/bankpj/internal/models"
)

func GetUser(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	c.JSON(200, user)
}