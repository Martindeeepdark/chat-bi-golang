package services

import (
	"chat-bi-golang/services/sessionmanager"
	"fmt"
	"strings"
)

// Process AI 处理逻辑
// Todo 要求根据文件后缀，选择处理的方式
func Process(data [][]string, request string) (string, error) {
	// 将数据转换为 CSV 格式的字符串
	var sb strings.Builder
	for _, row := range data {
		sb.WriteString(strings.Join(row, ","))
		sb.WriteString("\n")
	}
	csvData := sb.String()

	// 构建 prompt
	prompt := fmt.Sprintf("分析需求：%s 原始数据：%s", request, csvData)
	prompt = strings.ReplaceAll(prompt, "\n", " ")

	// 调用 AI 进行处理
	answer, err := xfChat(prompt)
	if err != nil {
		return "", err
	}
	return answer, nil
}

// xfChat 调用 AI 进行处理
func xfChat(prompt string) (string, error) {
	if sessionmanager.Session == nil {
		return "", fmt.Errorf("session is not initialized")
	}

	// 发送 prompt 并获取 AI 的回答
	answer, err := sessionmanager.Session.Send(prompt)
	if err != nil {
		return "", err
	}
	return answer, nil
}
