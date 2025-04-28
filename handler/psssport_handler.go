package handler

import (
	"bufio"
	"demo/handler/common"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"gorm.io/gorm"
)

// 文件格式只有一个 omnisdk uid 的场景

const (
	source_file = "/Users/leon/Downloads/moyu/2025-03-10/20250310.csv"
	target_file = "/Users/leon/Downloads/moyu/2025-03-10/20250310-result.csv" // create manually if not exists
)

func ReadFileAndQueryExtData(db *gorm.DB) {
	fs, err := os.Open(source_file)
	if err != nil {
		panic("read file error")
	}
	defer fs.Close()

	scanner := bufio.NewScanner(fs)
	loop := 0
	batchSize := 10 // 批量查询 size
	var result []string
	//针对大文件，一行一行的读取文件
	for scanner.Scan() {
		line := scanner.Text()
		// Clean the line if it contains quotes
		if strings.Contains(line, "\"") {
			line = strings.ReplaceAll(line, "\"", "")
		}
		result = append(result, line)

		loop++
		if loop >= batchSize {
			processPassportData(db, result)
			result = nil
			loop = 0
		}
	}

	// Check for any scanner errors
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	if len(result) > 0 {
		processPassportData(db, result)
	}
}

func processPassportData(db *gorm.DB, result []string) {
	retMap := make(map[string]string)

	originPlayerIds := make(map[string]string)
	var playerIds []string
	for i := range result {
		playerId := result[i]
		if playerId == "" {
			continue
		}

		index := strings.Index(playerId, "__")
		realPlayerId := playerId[index+2:]
		if strings.Contains(playerId, "ios_jinshanApple") {
			index = strings.Index(playerId, "__")
			realPlayerId = playerId[index+2:]
		}
		playerIds = append(playerIds, realPlayerId)
		originPlayerIds[realPlayerId] = playerId
	}

	var resultList []Passport
	err := db.Table("SDK_EXP_PLAYER").Where("exp_player_id in ?", playerIds).Select("passport_id", "exp_player_id").Find(&resultList).Error
	if err != nil {
		fmt.Printf("query from db error: %v", err)
		return
	}

	for i := range resultList {
		if common.IsDigit(resultList[i].PassportId) {
			retMap[resultList[i].PassportId] = originPlayerIds[resultList[i].ExpPlayerId]
		}
	}

	if len(retMap) > 0 {
		writeFile(retMap)
	} else {
		fmt.Println("--------- No Result ------------")
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
		if v == "" || k == "" {
			continue
		}
		r := []string{k, v}
		data = append(data, r)
	}

	//写入数据
	w.WriteAll(data)
	w.Flush()
}

type Passport struct {
	PassportId  string `gorm:"column:passport_id"`
	ExpPlayerId string `gorm:"column:exp_player_id"`
}
