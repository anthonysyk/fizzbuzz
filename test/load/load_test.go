package load

import (
	"fmt"
	"github.com/anthonysyk/fizzbuzz/api"
	"github.com/anthonysyk/fizzbuzz/store"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"sync"
	"testing"
	"time"

	"go.uber.org/ratelimit"
)

func TestLoad(t *testing.T) {
	requests := make(map[string]int)
	requests["/fizzbuzz?int1=4&int2=8&limit=50&str1=fizz&str2=buzz"] = 19
	requests["/fizzbuzz?int1=8&int2=12&limit=50&str1=fizz&str2=buzz"] = 12
	requests["/fizzbuzz?int1=10&int2=87&limit=1000&str1=fizz&str2=buzz"] = 95
	requests["/fizzbuzz?int1=2&int2=7&limit=200&str1=fizz&str2=buzz"] = 29
	requests["/fizzbuzz?int1=3&int5=8&limit=100&str1=fizz&str2=buzz"] = 31

	total := 0
	for _, count := range requests {
		total += count
	}

	st := store.NewMemoryStore()
	router := api.GetRouter(st)
	go func() { http.ListenAndServe(":8081", router) }()
	client := resty.New()

	rl := ratelimit.New(100, ratelimit.Per(time.Second))
	var wg sync.WaitGroup
	wg.Add(total)
	for url, hits := range requests {
		for i := 0; i < hits; i++ {
			rl.Take()
			go client.R().Get(fmt.Sprintf("http://localhost:8081%s", url))
			wg.Done()
		}
	}
	wg.Wait()

	res, err := client.R().Get(fmt.Sprintf("http://localhost:8081/stats"))
	assert.NoError(t, err)
	expected := `{"query":{"params":"int1=10&int2=87&limit=1000&str1=fizz&str2=buzz","hits":95}}`
	assert.Equal(t, expected, res.String())
}
