package api

import (
	"github.com/anthonysyk/fizzbuzz/middleware"
	"github.com/anthonysyk/fizzbuzz/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRouter(st store.Store) http.Handler {
	r := gin.Default()
	r.GET("/health", GetHealth)
	r.GET("/fizzbuzz", middleware.Stats(st), GetFizzBuzz)
	r.GET("/stats", GetStats(st))
	return r.Handler()
}
