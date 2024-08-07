package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Any("/account/*proxyPath", func(c *gin.Context) {
		proxyRequest(c, "http://localhost:8080")
	})

	r.Any("/profile/*proxyPath", func(c *gin.Context) {
		proxyRequest(c, "http://localhost:8080")
	})

	r.Any("/project/*proxyPath", func(c *gin.Context) {
		proxyRequest(c, "http://localhost:7777")
	})

	r.Any("/posts/*proxyPath", func(c *gin.Context) {
		proxyRequest(c, "http://localhost:8888")
	})

	r.Run(":4444")
}

func proxyRequest(c *gin.Context, target string) {
	targetURL, err := url.Parse(target)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid target URL"})
		return
	}

	// Create a new reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	// Get the path prefix from the route
	prefix := c.Param("proxyPath")
	if strings.HasPrefix(prefix, "/") {
		prefix = prefix[1:]
	}

	// Remove the prefix from the request path
	newPath := strings.TrimPrefix(c.Request.URL.Path, "/"+prefix)

	// Remove the trailing slash if it exists
	if strings.HasSuffix(newPath, "/") && newPath != "/" {
		newPath = newPath[:len(newPath)-1]
	}

	c.Request.URL.Scheme = targetURL.Scheme
	c.Request.URL.Host = targetURL.Host
	c.Request.URL.Path = newPath
	c.Request.Host = targetURL.Host

	// Print debugging information
	fmt.Printf("Original Path: %s\n", c.Request.URL.Path)
	fmt.Printf("Forwarding to: %s%s\n", targetURL.String(), c.Request.URL.Path)
	fmt.Printf("Full URL: %s\n", targetURL.ResolveReference(c.Request.URL).String())

	// Serve the HTTP request
	proxy.ServeHTTP(c.Writer, c.Request)
}
