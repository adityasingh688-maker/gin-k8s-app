package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// setupRouter is separate from main() so tests can use it.
func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello from Gin on Kubernetes!"})
	})

	// Used by Kubernetes liveness/readiness probes
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.GET("/api/hello/:name", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": c.Param("name")})
	})

	return r
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// TLS is terminated at the Ingress (see k8s/ingress.yaml)
	if err := setupRouter().Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
