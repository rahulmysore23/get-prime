package http_server

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rahulmysore23/get-prime/pkg/utilities"
)

// Start will use a serve config pkg to setup the server and it's dependencies
func Start() error {
	return Run()
}

// Run will inject 3rd party resources and create REST endpoints
func Run() error {
	routes := gin.New()
	routes.Use(
		CORS(),
		gin.Recovery(),
		gin.Logger(),
	)

	// v1 := routes.Group("/v1",
	// 	gin.Logger(),
	// )

	// v1.GET("")
	// v1.GET("")

	// For test - REMOVE AFTER ADDING UI
	routes.GET("/", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/html")
		c.Writer.WriteHeader(http.StatusOK)
		io.WriteString(c.Writer, utilities.DefaultHtml)
	})

	fmt.Printf("webserver listening on port %v\n", "6060") // TODO - Move to env
	err := http.ListenAndServe(":6060", routes)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Writer.Header()
		header.Set("Access-Control-Allow-Origin", "*")
		header.Set("Access-Control-Allow-Methods", "*")
		header.Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
		header.Set("Access-Control-Max-Age", "86400")

		if c.Request.Method == http.MethodOptions {
			c.Writer.WriteHeader(http.StatusOK)
			return
		}

		c.Next()
	}
}
