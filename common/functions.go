package common

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"random-quote-picture/common/define"
	"strings"
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

func GetHash(key string) string {
	hasher := sha256.New()
	hasher.Write([]byte(key))
	hashBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}

func MakeUuid() string {
	id := uuid.New()
	s := id.String()
	s = strings.ReplaceAll(s, "-", "")
	return s
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
		Code:    define.SUCCESSCODE,
		Message: "success",
		Data:    data,
	})
}
