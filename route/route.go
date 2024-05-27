// route/setroute.go
package route

import (
	"chat-bi-golang/apis"
	"chat-bi-golang/configs"
	"chat-bi-golang/handler"
	"github.com/gin-gonic/gin"
)

func Setroute(config configs.Config) *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("tmpl/*")
	r.Static("/static", "./static")

	r.GET("/login", handlers.ShowLoginPage)
	r.POST("/login", handlers.HandleLoginPost)
	r.GET("/upload", handlers.ShowUploadPage)
	r.POST("/process", apis.ProcessRequestHandler)
	r.POST("/upload_data", apis.UploadDataHandler)
	r.POST("/echat", apis.GetEchat)
	dataAnalysisGroup := r.Group("/data_analysis")
	dataAnalysisGroup.POST("/upload_data", apis.UploadDataHandler)

	return r
}
