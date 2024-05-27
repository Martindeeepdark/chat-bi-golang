package apis

import (
	"chat-bi-golang/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ProcessRequestHandler 处理数据分析请求
func ProcessRequestHandler(c *gin.Context) {
	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文件上传失败"})
		return
	}

	// 获取请求内容
	request := c.PostForm("request")

	// 保存上传的文件到服务器临时目录
	filePath := "/tmp/" + file.Filename
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		fmt.Println("文件保存失败:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文件保存失败"})
		return
	}

	// 解析文件内容
	parseResult, err := services.ParseFile(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文件解析失败"})
		return
	}
	// 对解析结果进行AI处理，并记录到数据库
	aiResult, err := services.Process(parseResult, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "AI处理失败"})
		return
	}

	// 返回处理结果，确保返回的数据格式正确
	c.JSON(http.StatusOK, gin.H{
		"message":         "分析处理成功",
		"analysis_report": aiResult,
	})
}
