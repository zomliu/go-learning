package handler

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"gorm.io/gorm"
)

const (
	source_file = "/Users/leon/Downloads/moyu2.csv"
	target_file = "/Users/leon/Downloads/moyu2-result.csv"
)

func processPassportData(db *gorm.DB, result []string) {
	retMap := make(map[string]string)
	for i := range result {
		playerId := result[i]
		if playerId == "" {
			continue
		}
		var ret struct {
			PassportId string `gorm:"column:passport_id"`
		}
		err := db.Table("SDK_EXP_PLAYER").Where("exp_player_id=?", playerId).Select("passport_id").First(&ret).Error
		if err != nil {
			if errors.Is(gorm.ErrRecordNotFound, err) {
				fmt.Printf("No record found: %s", playerId)
				continue
			}
			fmt.Printf("query from db error: %v", err)
			return
		}
		retMap[playerId] = ret.PassportId
	}
	if len(retMap) > 0 {
		writeFile(retMap)
	} else {
		fmt.Println("--------- No Result ------------")
	}
}

func ReadFileAndQueryExtData(db *gorm.DB) {
	fs, err := os.Open(source_file)
	if err != nil {
		panic("read file error")
	}
	defer fs.Close()
	r := csv.NewReader(fs)
	loop := 0
	batchSize := 1000 // 批量查询 size
	var result []string
	//针对大文件，一行一行的读取文件
	for {
		var row []string
		row, err = r.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
		}
		if err == io.EOF {
			break
		}

		tmp := row[0]
		index := strings.Index(tmp, "__")
		val := tmp[index+2:]
		result = append(result, val)

		loop += 1
		if loop >= batchSize {
			processPassportData(db, result)
			result = nil
			loop = 0
		}
	}
	if len(result) > 0 {
		processPassportData(db, result)
	}
}

func writeFile(retMap map[string]string) {
	//追加写文件
	f, err := os.OpenFile(target_file, os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// 写入UTF-8 BOM
	//f.WriteString("\xEF\xBB\xBF")
	//创建一个新的写入文件流
	w := csv.NewWriter(f)
	// data := [][]string{
	//     {"1", "刘备"},
	//     {"2", "张飞"},
	// }
	var data [][]string
	for k, v := range retMap {
		if v == "" {
			continue
		}
		r := []string{k, v}
		data = append(data, r)
	}
	//写入数据
	w.WriteAll(data)
	w.Flush()
}
