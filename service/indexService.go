package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexPage(c *gin.Context) {
	data := gin.H{"title": "Random Quote & Picture"}
	c.HTML(http.StatusOK, "index.html", data)
}
