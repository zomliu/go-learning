package passport

import (
	"demo/handler/common"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cast"
	"gorm.io/gorm"
)

var appIds = []int{10037, 10038}

type PassportHandler struct {
	DB *gorm.DB
	SourceFile string // 传入的文件要符合 AccountVo 的结构，列的顺序： 0: RoleId, 1: ComboId, 2: SquenceId
	TargetFile string // create manually if not exists
}

type AccountVo struct {
	RoleId      string
	ComboId     string
	ExpPlayerId string
	SquenceId   string
	PassportId  string // 输出结果
}

type ResultDto struct {
	PassportId  string `gorm:"column:passport_id"`
	ExpPlayerId string `gorm:"column:exp_player_id"`
}

func (p *PassportHandler) ReadFileAndQuery() {
	fs, err := os.Open(p.SourceFile)
	if err != nil {
		panic("read file error")
	}
	defer fs.Close()
	r := csv.NewReader(fs)
	loop := 0
	batchSize := 50 // 批量查询 size
	var dataMap = make(map[string]*AccountVo, batchSize)
	var playerIds []string
	//一行一行的读取文件
	for {
		var row []string
		row, err = r.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
		}
		if err == io.EOF {
			break
		}

		vo := &AccountVo{RoleId: row[0], ComboId: row[1], SquenceId: row[2]}
		index := strings.Index(vo.ComboId, "__")
		vo.ExpPlayerId = strings.ToLower(vo.ComboId[index+2:])

		dataMap[vo.ExpPlayerId] = vo
		playerIds = append(playerIds, vo.ExpPlayerId)

		loop += 1
		if loop >= batchSize {
			if err := p.processPassportData(dataMap, playerIds); err != nil {
				log.Fatalf("process passport data error: %+v", err)
				break
			}
			dataMap = make(map[string]*AccountVo, batchSize)
			playerIds = make([]string, 0)
			loop = 0
		}
	}
	if len(dataMap) > 0 {
		if err := p.processPassportData(dataMap, playerIds); err != nil {
			log.Fatalf("process passport data error: %+v", err)
		}
	}
}

func (p *PassportHandler) processPassportData(resultMap map[string]*AccountVo, playerIds []string) error {
	var ret []ResultDto
	err := p.DB.Table("SDK_EXP_PLAYER").Where("app_id in (?) AND exp_player_id in (?)", appIds, playerIds).Select("passport_id", "exp_player_id").Find(&ret).Error
	if err != nil {
		fmt.Printf("query from db error: %v", err)
		return err
	}
	for i := range ret {
		if common.IsDigit(ret[i].PassportId) {
			pId := strings.ToLower(ret[i].ExpPlayerId)
			rr := resultMap[pId]
			if rr == nil {
				continue
			}
			rr.PassportId = ret[i].PassportId
		}
	}

	// order by SequenceId
	var resultList = make([]*AccountVo, 0)
	for _, v := range resultMap {
		resultList = append(resultList, v)
	}
	// sort by SequenceId
	sort.SliceStable(resultList, func(i, j int) bool {
		return cast.ToInt64(resultList[i].SquenceId) < cast.ToInt64(resultList[j].SquenceId)  
	})

	p.writeFile(resultList)
	return nil
}

func (p *PassportHandler) writeFile(resultList []*AccountVo) {
	//追加写文件
	f, err := os.OpenFile(p.TargetFile, os.O_APPEND|os.O_RDWR, 0666)
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
	for _, v := range resultList {
		r := []string{v.SquenceId, v.RoleId, v.ComboId, v.PassportId}
		data = append(data, r)
	}
	//写入数据
	w.WriteAll(data)
	w.Flush()
}
