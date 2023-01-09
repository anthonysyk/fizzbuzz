package middleware

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestGenerateUniqueKey(t *testing.T) {
	expectedResult := "int1=12&int2=123&limit=1000&str1=fizz&str2=buzz"
	url1, _ := url.Parse("/fizzbuzz?int1=12&int2=123&limit=1000&str1=fizz&str2=buzz")
	url2, _ := url.Parse("/fizzbuzz?str2=buzz&int2=123&int1=12&str1=fizz&limit=1000")
	uniqueKey1 := getQueryParams(*url1)
	uniqueKey2 := getQueryParams(*url2)
	assert.Equal(t, expectedResult, uniqueKey1)
	assert.Equal(t, expectedResult, uniqueKey2)
}
