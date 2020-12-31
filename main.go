package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryulitaro/crawler"
)

func getting(c *gin.Context) {
	url := c.DefaultQuery("url", "news.naver.com")
	crawler := &crawler.MyCrawler{
		BaseURL: url,
		Depth:   2,
	}
	result, err := crawler.Start()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}
func main() {
	router := gin.Default()

	router.GET("/crawler", getting)
	router.Run(":3004") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
