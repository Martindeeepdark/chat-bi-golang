// handlers/login.go
package handlers

import (
	"github.com/gin-gonic/gin"
)

func ShowLoginPage(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{
		"title": "Login",
	})
}

func HandleLoginPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "admin" && password == "password" {
		c.Redirect(302, "/upload")
	} else {
		c.HTML(200, "login.html", gin.H{
			"title": "Login",
			"error": "Invalid username or password",
		})
	}
}

func ShowUploadPage(c *gin.Context) {
	c.HTML(200, "upload.html", gin.H{"title": "Upload Data"})
}
