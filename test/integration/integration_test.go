package integration

import (
	"github.com/anthonysyk/fizzbuzz/api"
	"github.com/anthonysyk/fizzbuzz/store"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestIntegration(t *testing.T) {
	st := store.NewMemoryStore()
	router := api.GetRouter(st)
	go func() { http.ListenAndServe(":8081", router) }()
	client := resty.New()
	client.R().Get("http://localhost:8081/fizzbuzz?int1=3&int2=5&limit=100&str1=fizz&str2=buzz")
	res, err := client.R().Get("http://localhost:8081/stats")
	assert.NoError(t, err)
	expected := `{"query":{"params":"int1=3&int2=5&limit=100&str1=fizz&str2=buzz","hits":1}}`
	assert.Equal(t, expected, res.String())
}
