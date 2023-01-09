package api

import (
	"github.com/anthonysyk/fizzbuzz/model"
	"github.com/anthonysyk/fizzbuzz/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func GetFizzBuzz(c *gin.Context) {
	var query model.FizzBuzzQuery
	if err := c.BindQuery(&query); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "could not bind query parameters"})
		return
	}
	if query.IsEmpty() {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "all query params are empty"})
		return
	}
	res := FizzBuzz(query.Int1, query.Int2, query.Limit, query.Str1, query.Str2)
	c.JSON(http.StatusOK, res)
}

func GetStats(store store.Store) func(c *gin.Context) {
	return func(c *gin.Context) {
		muq := store.GetMostUsedQuery()
		if muq.Params == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "all query params are empty"})
			return
		}
		c.PureJSON(http.StatusOK, gin.H{"query": muq})
	}
}
