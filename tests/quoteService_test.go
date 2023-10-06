package tests

import (
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"random-quote-picture/common/define"
	"random-quote-picture/config"
	"random-quote-picture/service"
	"testing"
)

func init() {
	config.Init("../")
}

func TestGetRandomQuote(t *testing.T) {
	r := gin.Default()
	r.GET("/api/quote", service.GetRandomQuote)

	// create http request
	req, err := http.NewRequest("GET", "/api/quote", nil)
	if err != nil {
		t.Fatalf("create http request fail,%s\n", err)
	}
	// create a response recorder and send request to the gin engine
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)

	assert.Equal(t, recorder.Code, http.StatusOK)

	var resp define.Response
	err = jsoniter.UnmarshalFromString(recorder.Body.String(), &resp)
	if err != nil {
		t.Errorf("GetRandomQuote response error: %s\n", err.Error())
	}
	assert.Equal(t, resp.Code, define.SUCCESSCODE)

	bytes, err := jsoniter.Marshal(resp.Data)
	if err != nil {
		t.Error(err)
	}
	var quote service.Quote
	err = jsoniter.Unmarshal(bytes, &quote)
	if err != nil {
		t.Error(err)
	}
	if len(quote.QuoteText) < 0 {
		t.Errorf("quoteText is empty")
	}

	t.Log("GetRandomQuote PASSED.")
}

func TestGetOneQuote(t *testing.T) {
	key := ""
	quote, err := service.GetOneQuote(key)
	if err != nil {
		t.Errorf("GetOneQuote Failed, error: %s\n", err.Error())
		return
	}
	if quote.QuoteText == "" {
		t.Errorf("GetOneQuote Failed, quoteText is empty!\n")
		return
	}
	t.Log("GetOneQuote PASSED.")
}

func BenchmarkGetOneQuote(b *testing.B) {
	key := ""
	for i := 0; i < b.N; i++ {
		service.GetOneQuote(key)
	}
}
