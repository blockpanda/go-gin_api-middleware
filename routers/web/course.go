package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Create",
	})

}
func Edit(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Edit",
	})

}
func Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get",
	})

}
func Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete",
	})

}
func Create2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Create2",
	})

}
