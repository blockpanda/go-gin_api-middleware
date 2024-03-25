package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Registe(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "registe",
	})
}
