package main

import (
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

// ReverseProxy ...
func ReverseProxy() gin.HandlerFunc {
	return func(c *gin.Context) {
		director := func(req *http.Request) {
			// r := c.Request
			// req = r
			req.URL.Scheme = "http"
			req.URL.Host = "localhost:35000"
			// req.Header["X-Forwarded-For"] = []string{r.Header.Get("my-header")}
			// Golang camelcases headers
			delete(req.Header, "My-Header")
		}
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Any("/", ReverseProxy())
	r.Run()
}
