package services

import (
	"encoding/csv"
	"errors"
	"github.com/xuri/excelize/v2"
	"os"
	"path/filepath"
)

// ParseFile 解析文件，根据文件扩展名调用相应的解析函数
func ParseFile(filePath string) ([][]string, error) {
	fileExt := filepath.Ext(filePath)
	switch fileExt {
	case ".csv":
		return ParseCSV(filePath)
	case ".xlsx", ".xls":
		return ParseExcel(filePath)
	default:
		return nil, errors.New("不支持的文件格式")
	}
}

// ParseCSV 解析 CSV 文件
func ParseCSV(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

// ParseExcel 解析 Excel 文件
func ParseExcel(filePath string) ([][]string, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var records [][]string
	sheetName := f.GetSheetName(1)
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, err
	}

	for _, row := range rows {
		records = append(records, row)
	}

	return records, nil
}
