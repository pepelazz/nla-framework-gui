package webServer

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const PORT = 3076

func Start() {
	r := gin.New()

	r.Use(LiberalCORS)
	r.Static("/static", "./webServer/static")

	r.POST("/getProject", GetProject)
	r.POST("/saveProject", SaveProject)
	r.POST("/projectTemplateGenerate", ProjectTemplateGenerate)

	r.Run(fmt.Sprintf(":%v", PORT))
}

// LiberalCORS is a very allowing CORS middleware.
func LiberalCORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	if c.Request.Method == "OPTIONS" {
		if len(c.Request.Header["Access-Control-Request-Headers"]) > 0 {
			c.Header("Access-Control-Allow-Headers", c.Request.Header["Access-Control-Request-Headers"][0])
		}
		c.AbortWithStatus(http.StatusOK)
	}
}