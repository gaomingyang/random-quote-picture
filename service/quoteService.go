package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/viper"
	"log"
	"random-quote-picture/common"
	"strconv"
	"strings"
)

type Quote struct {
	QuoteText   string `json:"quoteText"`
	QuoteAuthor string `json:"quoteAuthor"`
	SenderName  string `json:"senderName"`
	SenderLink  string `json:"senderLink"`
	QuoteLink   string `json:"quoteLink"`
}

func Test(c *gin.Context) {
	q := Quote{
		QuoteText: "Do you want to know who you are? Don't ask. Act! Action will delineate and define you.",
	}
	common.OK(c, q)
}

func GetRandomQuote(c *gin.Context) {
	key := c.DefaultQuery("key", "")
	if key != "" {
		if len(key) > 6 {
			common.BadRequest(c, "key length can't more than 6")
			return
		}
		if _, err := strconv.Atoi(key); err != nil {
			common.BadRequest(c, "key is not illegal integer")
			return
		}
	}

	quote, err := getOneQuote(key)
	if err != nil {
		log.Println("get quote error", err)
		common.InternalServerError(c, "get quote err")
		return
	}
	common.OK(c, quote)
}

func getOneQuote(key string) (quote Quote, err error) {
	quoteServerUrl := viper.GetString("quoteServerUrl")
	quoteMethod := viper.GetString("quoteMethod")
	quoteFormat := viper.GetString("quoteFormat")
	quoteLang := viper.GetString("quoteLang")
	if quoteServerUrl == "" || quoteMethod == "" || quoteFormat == "" || quoteLang == "" {
		err = errors.New("quote config error")
		return
	}
	getQuoteUrl := fmt.Sprintf("%s?method=%s&format=%s&lang=%s", quoteServerUrl, quoteMethod, quoteFormat, quoteLang)
	if key != "" {
		getQuoteUrl += "&key=" + key
	}
	fmt.Println("quoteServerUrl:", getQuoteUrl)
	res, err := common.HttpGet(getQuoteUrl)
	// _, err = strconv.Unquote(res) // doesn't work
	err = jsoniter.UnmarshalFromString(res, &quote)
	if err != nil {
		// fix the "\'" problem in quoteText
		if strings.Contains(err.Error(), "readEscapedChar") {
			res = strings.ReplaceAll(res, "\\", "")
		}
		fmt.Println("new res:", res)
		err = jsoniter.UnmarshalFromString(res, &quote)
	}
	return
}
