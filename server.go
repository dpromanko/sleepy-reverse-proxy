package main

import (
	"fmt"
	"net/http/httputil"
	"net/url"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func startServer(port string, proxyURL string, sleep time.Duration) error {
	u, err := url.Parse(proxyURL)
	if err != nil {
		return err
	}

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	rp := httputil.NewSingleHostReverseProxy(u)

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))
	r.Use(sleepyMiddleware(sleep))

	r.NoRoute(proxyHandler(u.Host, rp))

	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		return err
	}

	return nil
}

func sleepyMiddleware(d time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		time.Sleep(d)
		c.Next()
	}
}

func proxyHandler(host string, rp *httputil.ReverseProxy) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Host = host
		rp.ServeHTTP(c.Writer, c.Request)
	}
}
