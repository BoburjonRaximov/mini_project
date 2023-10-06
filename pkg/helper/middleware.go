package helper

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func StartMiddleware(c *gin.Context) {
	// before request
	fmt.Printf("%s request start path: %s time %v\n", c.Request.Method, c.Request.URL.Path, time.Now())
	c.Next()
}
func EndMiddleware(c *gin.Context) {
	c.Next()
	// after request
	fmt.Printf("%s request end path: %s time %v\n", c.Request.Method, c.Request.URL.Path, time.Now())
}
func LoggerAllInOne(c *gin.Context) {
	// before request
	t := time.Now()
	c.Next()
	// after request
	latency := time.Since(t)

	// access the status we are sending
	status := c.Writer.Status()

	fmt.Printf("Completed %s %s with status code %d in %v\n", c.Request.Method, c.Request.URL.Path, status, latency)
}
func Logger(c *gin.Context) {
	// before request
	beforeRequest(c)
	c.Next()
	// after request
	afterRequest(c)
}
func beforeRequest(c *gin.Context) {
	// before request
	t := time.Now()
	c.Set("start", t)
	c.Next()
}
func afterRequest(c *gin.Context) {
	// Get the start time from the request context
	startTime, exists := c.Get("start")
	if !exists {
		startTime = time.Now()
	}

	// Calculate the request duration
	duration := time.Since(startTime.(time.Time))

	// Log the request completion time and duration
	fmt.Printf("Completed %s %s in %v\n", c.Request.Method, c.Request.URL.Path, duration)
}
func AuthMiddleware(c *gin.Context) {

}
