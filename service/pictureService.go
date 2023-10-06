package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"math/rand"
	"random-quote-picture/common"
)

type PictureUrlResponse struct {
	PictureUrl string `json:"pictureUrl"`
}

// GetRandomPicture api
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

// GetPictureUrl - get picture from picsum (https://picsum.photos/)
func GetPictureUrl(grayscale string) (pictureUrl string, err error) {
	pictureServerUrl := viper.GetString("pictureServerUrl")
	pictureWidth := viper.GetString("pictureWidth")
	pictureHeight := viper.GetString("pictureHeight")
	if pictureServerUrl == "" || pictureWidth == "" || pictureHeight == "" {
		err = errors.New("picture config error")
		return
	}
	seed := common.MakeUuid()
	pictureUrl = fmt.Sprintf("%s/seed/%s/%s/%s", pictureServerUrl, seed, pictureWidth, pictureHeight)
	if grayscale == "1" {
		pictureUrl += "?grayscale"
	}
	return
}

func GetBackupPicture(c *gin.Context) {
	backupPictureUrls := viper.GetStringSlice("pictureBackupUrl")
	index := rand.Intn(len(backupPictureUrls))
	backupPictureUrl := backupPictureUrls[index]
	res := PictureUrlResponse{
		PictureUrl: backupPictureUrl,
	}
	common.OK(c, res)
}
