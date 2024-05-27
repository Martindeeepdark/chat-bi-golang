// services/sessionmanager/sessionmanager.go
package sessionmanager

import (
	"github.com/Martindeeepdark/xf-golang-sdk/chat"
	"log"
	"sync"
)

var (
	Session *chat.Session
	once    sync.Once
)

// InitSession 初始化 WebSocket 会话
func InitSession(appID, apiKey, apiSecret, hostURL string) {
	once.Do(func() {
		var err error
		s := chat.NewServer(appID, apiKey, apiSecret, hostURL)
		Session, err = s.GetSession("123456789")
		if err != nil {
			log.Fatalf("Failed to initialize chat session: %v", err)
		}
	})
}

// GetSession 获取当前的 WebSocket 会话
func GetSession() *chat.Session {
	return Session
}
