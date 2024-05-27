package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"os"
	"path/filepath"
)

// FileInfo 结构体用于存储文件信息
type FileInfo struct {
	ID       string
	Filename string
	Request  string
}

var files = make(map[string]FileInfo)

// UploadDataHandler 处理上传数据请求
func UploadDataHandler(c *gin.Context) {
	// 从请求中获取文件内容和文件头信息
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无法获取文件"})
		return
	}
	defer file.Close()

	// 检查文件大小，限制为20MB
	const maxFileSize = 20 << 20 // 20MB
	if header.Size > maxFileSize {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文件大小不能超过20MB"})
		return
	}

	// 检查文件格式（扩展名）
	fileExt := filepath.Ext(header.Filename)
	switch fileExt {
	case ".xlsx", ".xls", ".csv":
		// 文件格式符合要求，继续处理
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "不支持的文件格式"})
		return
	}

	// 在保存文件之前，确保file目录存在
	uploadDir := "file" // 指定的文件夹名称
	if err := ensureUploadsDirExists(uploadDir); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法创建上传目录"})
		return
	}

	// 生成唯一文件名
	uniqueID := uuid.New().String()
	savePath := filepath.Join(uploadDir, uniqueID+fileExt)

	// 现在可以安全地保存文件了
	if err := c.SaveUploadedFile(header, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 记录文件信息
	files[uniqueID] = FileInfo{
		ID:       uniqueID,
		Filename: savePath,
	}

	c.JSON(http.StatusOK, gin.H{"message": "文件上传成功", "file_id": uniqueID})
}

// ensureUploadsDirExists 确保上传目录存在
func ensureUploadsDirExists(uploadDir string) error {
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		// 如果目录不存在，则创建目录
		return os.MkdirAll(uploadDir, os.ModePerm) // 使用os.ModePerm提供足够的权限
	}
	return nil
}
