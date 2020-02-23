package main

import (
	"fmt"
	"net/http/httputil"
	"net/url"
	"os"
	"time"

	ginzap "github.com/gin-contrib/zap"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	gin.SetMode(gin.ReleaseMode)

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	u, _ := url.Parse("http://localhost:8090")
	rp := httputil.NewSingleHostReverseProxy(u)

	r := gin.New()
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))
	r.Use(sleepyMiddleware)
	r.NoRoute(proxyHandler(u.Host, rp))
	return r.Run()
}

func sleepyMiddleware(c *gin.Context) {
	time.Sleep(10000 * time.Millisecond)
	c.Next()
}

func proxyHandler(host string, rp *httputil.ReverseProxy) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Host = host
		rp.ServeHTTP(c.Writer, c.Request)
	}
}
