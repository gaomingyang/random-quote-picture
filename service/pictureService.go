package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"random-quote-picture/common"
)

type PictureUrlResponse struct {
	PictureUrl string `json:"pictureUrl"`
}

func GetRandomPicture(c *gin.Context) {
	grayscale := c.DefaultQuery("grayscale", "")
	pictureUrl, err := GetPictureUrl(grayscale)
	if err != nil {
		log.Println("get picture error", err)
		common.InternalServerError(c, "get picture err")
		return
	}
	res := PictureUrlResponse{
		PictureUrl: pictureUrl,
	}
	common.OK(c, res)
}

func GetPictureUrl(grayscale string) (pictureUrl string, err error) {
	pictureServerUrl := viper.GetString("pictureServerUrl")
	pictureWidth := viper.GetString("pictureWidth")
	pictureHeight := viper.GetString("pictureHeight")
	if pictureServerUrl == "" || pictureWidth == "" || pictureHeight == "" {
		err = errors.New("picture config error")
		return
	}
	pictureUrl = fmt.Sprintf("%s/%s/%s", pictureServerUrl, pictureWidth, pictureHeight)
	if grayscale == "1" {
		pictureUrl += "?grayscale"
	}
	return
}
