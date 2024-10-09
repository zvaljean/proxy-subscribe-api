package data

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"valjean/proxy/subscribe/pkg/log"
)

// ParseCsvForMap 解析csv文件，生成map数据，key,value为要获取的列数,从1开始
func ParseCsvForMap(filepath string, key int, value int) *map[string]string {

	// 打开CSV文件
	file, err := os.Open(filepath)
	log.FatalCheck(err, "open file fail!")
	defer file.Close()
	// 创建一个新的CSV reader
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 4
	context := make(map[string]string)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// 解析中有注释行，预见不符合的列数，跳过
		if er, ok := err.(*csv.ParseError); ok {
			if errors.As(er.Err, &csv.ErrFieldCount) {
				continue
			}
		}

		log.FatalCheck(err, "read error!")
		if key == value {
			log.Fatal("key: %d or value: %d column equal ", key, value)
		}

		if len(record) < 4 {
			log.Warn("the record is error: %s", reader)
			continue
		}
		// 跳过#注释
		// https://www.ascii-code.com/35
		if strings.TrimSpace(record[0])[0] == 35 {
			continue
		}

		if key > len(record) || value > len(record) {
			log.Fatal("key: %d or value: %d column bigger than total ", key, value)
		}

		// key: 2,3 列作为key, 4: 作为value
		k := fmt.Sprintf("%s-%s", record[key-1], record[key])
		v := record[value-1]

		if len(k) == 0 || len(v) == 0 {
			log.Info("key: %s, value: %s have empty value.", k, v)
			continue
		}
		context[k] = v
	}
	return &context
}

// ParseCsvForList 解析csv文件，生产list数据, column为要获取的列数,从1开始
func ParseCsvForList(filepath string, column int) *[]string {
	// 打开CSV文件
	file, err := os.Open(filepath)
	log.FatalCheck(err, "open file fail!")
	defer file.Close()

	// 创建一个新的CSV reader
	reader := csv.NewReader(file)

	// 可选：设置字段数量（如果确定只有两列）
	//reader.FieldsPerRecord = 2

	strings := make([]string, 0)

	// 逐行读取CSV文件
	for {
		// 读取一行
		record, err := reader.Read()
		if err == io.EOF {
			break // 文件结束
		}
		log.FatalCheck(err, "write error!")

		// 确保这行有两列
		if column > len(record) {
			log.Fatal("the column %s empty!", column)
		}

		// 处理这行数据
		column1 := record[column-1]
		if len(column1) == 0 {
			continue
		}

		strings = append(strings, column1)
	}

	return &strings
}
