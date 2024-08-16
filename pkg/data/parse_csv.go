package data

import (
	"encoding/csv"
	"io"
	"os"
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
	context := make(map[string]string)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		log.FatalCheck(err, "write error!")
		if key == value {
			log.Fatal("key: %d or value: %d column equal ", key, value)
		}

		if key > len(record) || value > len(record) {
			log.Fatal("key: %d or value: %d column bigger than total ", key, value)
		}

		k := record[key-1]
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
