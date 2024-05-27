package services

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	ID           uint
	Username     string
	PasswordHash string
}

func Authenticate(db *gorm.DB, username, password string) error {
	var user User
	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return result.Error
	}

	// 直接比较密码（注意：实际生产环境中建议使用哈希密码）
	if user.PasswordHash != password {
		return errors.New("invalid credentials")
	}

	return nil
}
