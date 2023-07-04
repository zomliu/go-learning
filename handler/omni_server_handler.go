package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"gorm.io/gorm"
)

var speciaficOrderNo = []string{
	"6230271048056418",
	"62302b1048054066",
	"6230211048056236",
	"6230241048056417",
}

// 处理指定订单号记录
func QuerySpecificOrder(db *gorm.DB) {
	var orderDetails []OrderDetail
	err := db.Table("recharge_log_6").Where("trade_no in ?", speciaficOrderNo).Find(&orderDetails).Error
	if err != nil {
		panic("query from db error")
	}

	//var totalResult []string
	for i := range orderDetails {
		orderDetails[i].PrintCreateTime = orderDetails[i].CreateTime.UnixMilli()
		orderDetails[i].PrintPaidTime = orderDetails[i].PaidTime.UnixMilli()
		orderDetails[i].PrintFinishTime = orderDetails[i].FinishTime.UnixMilli()
		var b []byte
		if b, err = json.Marshal(orderDetails[i]); err != nil {
			fmt.Println(err)
			panic("marshal error")
		}
		//totalResult = append(totalResult, string(b))
		fmt.Println(string(b))
	}
}

// write order detail to json file
func WriteOrderToFile(db *gorm.DB) {
	tableName := "recharge_log_4"
	pageSize := 200
	currentLogId := int64(1047723935)

	fmt.Print(currentLogId)

	var totalResult []string
	for {
		var orderDetails []OrderDetail
		err := db.Table(tableName).
			Where("finish_time>'2023-04-27 23:59:59' and finish_time<'2023-04-28 10:10:00' and charge_log_id>?", currentLogId).
			Order("charge_log_id asc").Limit(pageSize).Find(&orderDetails).Error

		if err != nil {
			panic("query from db error")
		}

		for i := range orderDetails {
			orderDetails[i].PrintCreateTime = orderDetails[i].CreateTime.UnixMilli()
			orderDetails[i].PrintPaidTime = orderDetails[i].PaidTime.UnixMilli()
			orderDetails[i].PrintFinishTime = orderDetails[i].FinishTime.UnixMilli()
			var b []byte
			if b, err = json.Marshal(orderDetails[i]); err != nil {
				fmt.Println(err)
				panic("marshal error")
			}
			totalResult = append(totalResult, string(b))

		}

		if cap(orderDetails) < pageSize {
			break
		}
		currentLogId = orderDetails[pageSize-1].ChargeLogId
	}

	writeSomethingToFile(totalResult)
}

// 以行追加的方式写文件
func writeSomethingToFile(orders []string) {
	fmt.Printf("result size: %d\n", cap(orders))
	target_file := "recharge_2023-04-28_06-00.log"
	f, err := os.OpenFile(target_file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("open file error")
	}
	defer f.Close()

	for i := range orders {
		f.WriteString(orders[i] + "\n")
	}
}

type OrderDetail struct {
	ChargeLogId             int64     `json:"chargeLogId" gorm:"column:charge_log_id"`
	AdChannelId             string    `json:"adChannelId" gorm:"column:ad_channel_id"`
	BuildNumber             string    `json:"buildNumber" gorm:"column:build_number"`
	ChannelAppId            string    `json:"channelAppId" gorm:"column:channel_app_id"`
	ChannelBonusAmount      int64     `json:"channelBonusAmount" gorm:"column:channel_bonus_amount"`
	ChannelId               string    `json:"channelId" gorm:"column:channel_id"`
	ChannelTradeNo          string    `json:"channelTradeNo" gorm:"column:channel_trade_no"`
	CreateTime              time.Time `json:"-" gorm:"column:create_time"`
	PrintCreateTime         int64     `json:"createTime"`
	CurrencyName            string    `json:"currencyName" gorm:"column:currency_name"`
	CustomInfo              string    `json:"customInfo" gorm:"column:custom_info"`
	DeviceBrand             string    `json:"deviceBrand" gorm:"column:device_brand"`
	DeviceId                string    `json:"deviceId" gorm:"column:device_id"`
	DeviceIp                string    `json:"deviceIp"	gorm:"column:device_ip"`
	DeviceModel             string    `json:"deviceModel" gorm:"column:device_model"`
	FinishTime              time.Time `json:"-" gorm:"column:finish_time"`
	PrintFinishTime         int64     `json:"finishTime"`
	GameNotifyTimes         int       `json:"gameNotifyTimes" gorm:"column:game_notify_times"`
	GameTradeNo             string    `json:"gameTradeNo" gorm:"column:game_trade_no"`
	IsSandbox               bool      `json:"isSandbox" gorm:"column:is_sandbox"`
	IsTrialPeriod           bool      `json:"isTrialPeriod" gorm:"column:is_trial_period"`
	ManualOrderFlag         int       `json:"manualOrderFlag" gorm:"column:manual_order_flag"`
	PaidAmount              int       `json:"paidAmount" gorm:"column:paid_amount"`
	PaidTime                time.Time `json:"-" gorm:"column:paid_time"`
	PrintPaidTime           int64     `json:"paidTime"`
	PayType                 string    `json:"payType" gorm:"column:pay_type"`
	PlanId                  string    `json:"planId" gorm:"column:plan_id"`
	ProductDesc             string    `json:"productDesc" gorm:"column:product_desc"`
	ProductId               string    `json:"productId" gorm:"column:product_id"`
	ProductName             string    `json:"productName" gorm:"column:product_name"`
	ProductQuantity         int       `json:"productQuantity" gorm:"column:product_quantity"`
	ProductUnitPrice        int       `json:"productUnitPrice" gorm:"column:product_unit_price"`
	RoleId                  string    `json:"roleId" gorm:"column:role_id"`
	RoleLevel               string    `json:"roleLevel" gorm:"column:role_level"`
	RoleName                string    `json:"roleName" gorm:"column:role_name"`
	RoleVipLevel            string    `json:"roleVipLevel" gorm:"column:role_vip_level"`
	SearchChannelOrderTimes int       `json:"searchChannelOrderTimes" gorm:"column:search_channel_order_times"`
	ServerId                string    `json:"serverId" gorm:"column:server_id"`
	Status                  int       `json:"status" gorm:"column:status"`
	TableNamePostfix        string    `json:"tableNamePostfix" gorm:"column:table_name_postfix"`
	TimeZoneId              string    `json:"timeZoneId" gorm:"column:time_zone_id"`
	TokenId                 string    `json:"tokenId" gorm:"column:token_id"`
	TotalAmount             int       `json:"totalAmount" gorm:"column:total_amount"`
	TradeNo                 string    `json:"tradeNo" gorm:"column:trade_no"`
	Uid                     string    `json:"uid" gorm:"column:uid"`
	VoucherAmount           int       `json:"voucherAmount" gorm:"column:voucher_amount"`
	XgAppId                 string    `json:"xgAppId" gorm:"column:xg_app_id"`
	XgCustomInfo            string    `json:"xgCustomInfo" gorm:"column:xg_custom_info"`
	ZoneId                  string    `json:"zoneId" gorm:"column:zone_id"`
}

const (
	timeFormart = "2006-01-02 15:04:05"
	zone        = "Asia/Shanghai"
)

type LocalTime time.Time

func (t LocalTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormart)
	b = append(b, '"')
	return b, nil
}

func Marshal(v any, escapeHTML bool) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(escapeHTML)
	err := encoder.Encode(v)
	buffer.Truncate(buffer.Len() - 1) // remove newline
	return buffer.Bytes(), err
}
