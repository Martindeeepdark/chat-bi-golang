package apis

import (
	"chat-bi-golang/configs"
	"chat-bi-golang/services"
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
)

// 假设您有一个全局变量 db 用于数据库连接
var db *gorm.DB

func init() {
	// 初始化数据库连接
	var err error
	db, err = configs.SetupDatabase()
	if err != nil {
		panic("failed to connect database")
	}
}

// LoginHandler 处理登录请求
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	err := services.Authenticate(db, username, password)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	response := map[string]string{"message": "Login successful"}
	json.NewEncoder(w).Encode(response)
}
