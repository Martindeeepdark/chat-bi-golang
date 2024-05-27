package services

import (
	"fmt"
	"strings"
)

// processEchat 生成 ECharts 图表配置
func ProcessEchat(data [][]string) (string, error) {
	// 将数据转换为 CSV 格式的字符串
	var sb strings.Builder
	for _, row := range data {
		sb.WriteString(strings.Join(row, ","))
		sb.WriteString("\n")
	}
	csvData := sb.String()

	// 构建 prompt
	prompt := fmt.Sprintf("生成需求：根据下方原始数据生成echat图表代码 原始数据：%s", csvData)
	prompt = strings.ReplaceAll(prompt, "\n", " ")

	// 调用 AI 进行处理
	answer, err := xfChat(prompt)
	if err != nil {
		return "", err
	}
	return answer, nil
}
