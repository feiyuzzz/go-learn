package main

import (
	"log"

	"github.com/gin-gonic/gin"

	midleware "github.com/feiyuzzz/go-learn/go-gin/midleware/log"
)

func main() {
	r := gin.New()
	r.Use(midleware.Logger())
	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)
		// it would print: "12345"
		log.Println(example)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
