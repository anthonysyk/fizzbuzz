package middleware

import (
	"github.com/anthonysyk/fizzbuzz/store"
	"github.com/gin-gonic/gin"
	"net/url"
)

func Stats(st store.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := getQueryParams(*c.Request.URL)
		st.AddQueryHit(key)
		c.Next()
	}
}

func getQueryParams(u url.URL) string {
	return u.Query().Encode()
}
