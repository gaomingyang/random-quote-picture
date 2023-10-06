package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/viper"
	"log"
	"math/rand"
	"random-quote-picture/common"
	"strconv"
	"strings"
	"time"
)

type Quote struct {
	QuoteText   string `json:"quoteText"`
	QuoteAuthor string `json:"quoteAuthor"`
	SenderName  string `json:"senderName,omitempty"`
	SenderLink  string `json:"senderLink,omitempty"`
	QuoteLink   string `json:"quoteLink"`
}

var CachedQuotes map[string]Quote
var CachedQuoteKeys []string

func init() {
	CachedQuotes = make(map[string]Quote)
}

// func Test(c *gin.Context) {
// 	q := Quote{
// 		QuoteText: "Do you want to know who you are? Don't ask. Act! Action will delineate and define you.",
// 	}
// 	common.OK(c, q)
// }

func GetRandomQuote(c *gin.Context) {
	key := c.DefaultQuery("key", "")
	// check parameter
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

	// get quote data
	quote, err := GetOneQuote(key)
	if err != nil {
		log.Println("get quote from forismatic error", err)

		// If forismatic is unvailable, get data from cache.
		quote, err = GetQuoteFromCache()
	} else {
		// store quote data to cache
		hashKey := common.GetHash(quote.QuoteText)
		// length < 100 and key not exist
		if len(CachedQuoteKeys) < 100 {
			if _, ok := CachedQuotes[hashKey]; !ok {
				CachedQuotes[hashKey] = quote
				CachedQuoteKeys = append(CachedQuoteKeys, hashKey)
			}
		}
	}

	if err != nil {
		common.InternalServerError(c, "get quote err")
	}

	common.OK(c, quote)
}

// GetOneQuote - get quote from forismatic (https://forismatic.com/en/api/)
func GetOneQuote(key string) (quote Quote, err error) {
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
	// log.Println("getQuoteUrl:", getQuoteUrl)
	res, err := common.HttpGet(getQuoteUrl)
	// _, err = strconv.Unquote(res) // doesn't work
	err = jsoniter.UnmarshalFromString(res, &quote)
	if err != nil {
		// fix the "\'" problem in quoteText
		if strings.Contains(err.Error(), "readEscapedChar") {
			res = strings.ReplaceAll(res, "\\", "")
		}
		err = jsoniter.UnmarshalFromString(res, &quote)
	}
	return
}

// GetQuoteFromCache - this function will be called when forismatic api is unavailable
func GetQuoteFromCache() (quote Quote, err error) {
	n := len(CachedQuoteKeys)
	if n < 1 {
		err = errors.New("no cache quote")
		return
	}
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(n)
	key := CachedQuoteKeys[index]
	if key == "" {
		err = errors.New("cached quote data error")
		return
	}
	quote = CachedQuotes[key]
	return
}
