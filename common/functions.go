package common

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"random-quote-picture/common/define"
)

func HttpGet(url string) (res string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	res = string(body)
	return
}

func CommonJSON(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, define.Response{
		Code:    code,
		Message: message,
	})
}

func BadRequest(c *gin.Context, message string) {
	CommonJSON(c, http.StatusBadRequest, message)
}

func InternalServerError(c *gin.Context, message string) {
	CommonJSON(c, http.StatusInternalServerError, message)
}

func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, define.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}
