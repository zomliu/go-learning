package handler

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

func GetPassportDetail(passportId, env string) (string, error) {
	apiUrl := "http://img-pass.seasungame.com:8080/pass/query/user-detail"
	authAppId := "xigua_bj"
	signKey := "V9vF7qe5uag0p50WlhM"
	if env == "dev" {
		apiUrl = "http://125.88.194.72:8080/pass/query/user-detail"
		signKey = "hEJWTQBPVgjyhemDXrn"
	}

	timeNow := time.Now().Format("20060102150405")
	originStr := fmt.Sprintf("%s%s%s%s", passportId, authAppId, timeNow, signKey)
	signStr := SignWithMD5(originStr)

	values := map[string]string{"accountOpenId": passportId, "authAppId": authAppId, "timestamp": timeNow, "sign": signStr}
	jsonValue, _ := json.Marshal(values)

	// 设置默认 httpclient 的超时时间
	http.DefaultClient.Timeout = 5 * time.Second
	resp, err := http.Post(apiUrl, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	infoByte, _ := io.ReadAll(resp.Body)
	infoStr := string(infoByte)
	fmt.Println(infoStr)
	var result struct {
		Code    int64  `json:"code"`
		Message string `json:"message"`
		Data    struct {
			AccountOpenId string `json:"accountOpenId"`
			Account       string `json:"account"`
		}
	}
	if err = json.Unmarshal([]byte(infoStr), &result); err != nil {
		fmt.Printf("json error: %v", err)
		return "", err
	}
	if result.Code == 1 {
		return result.Data.AccountOpenId, nil
	}
	fmt.Printf("Code: %d, Message: %s", result.Code, result.Message)
	return result.Message, nil
}

func GetPassportOpenId(passportId, env string) (string, error) {
	time.Sleep(2 * time.Millisecond)
	apiUrl := "http://img-pass.seasungame.com:8080/pass/query/final-accountOpenId"
	authAppId := "xigua_bj"
	signKey := "V9vF7qe5uag0p50WlhM"
	if env == "dev" {
		authAppId = "xigua_bj_1"
		apiUrl = "http://125.88.194.72:8080/pass/query/final-accountOpenId"
		signKey = "pQdbcgmrQ6mqM"
	}

	timeNow := time.Now().Format("20060102150405")
	originStr := fmt.Sprintf("%s%s%s%s", passportId, authAppId, timeNow, signKey)
	signStr := SignWithMD5(originStr)

	values := map[string]string{"account": passportId, "authAppId": authAppId, "timestamp": timeNow, "sign": signStr}
	jsonValue, _ := json.Marshal(values)

	// 设置默认 httpclient 的超时时间
	http.DefaultClient.Timeout = 5 * time.Second
	resp, err := http.Post(apiUrl, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	infoByte, _ := io.ReadAll(resp.Body)
	infoStr := string(infoByte)

	var result struct {
		Code    int64  `json:"code"`
		Message string `json:"message"`
		Data    struct {
			AccountOpenId string `json:"accountOpenId"`
			Account       string `json:"account"`
		}
	}
	if err = json.Unmarshal([]byte(infoStr), &result); err != nil {
		fmt.Printf("json error: %v", err)
		return "", err
	}
	if result.Code == 1 {
		return result.Data.AccountOpenId, nil
	}
	fmt.Printf("Code: %d, Message: %s", result.Code, result.Message)
	return "", errors.New(result.Message)
}

func GetPassportVerifyInfo(passportId, env string) (string, error) {
	apiUrl := "http://img-pass.seasungame.com:8080/pass/query/verify-info"
	authAppId := "xigua_bj"
	signKey := "V9vF7qe5uag0p50WlhM"
	if env == "dev" {
		apiUrl = "http://125.88.194.72:8080/pass/query/verify-info"
		signKey = "hEJWTQBPVgjyhemDXrn"
	}

	clientIp := "36.112.24.16"

	timeNow := time.Now().Format("20060102150405")
	originStr := fmt.Sprintf("%s%s%s%s%s", passportId, clientIp, authAppId, timeNow, signKey)
	signStr := SignWithMD5(originStr)

	values := map[string]string{"accountOpenId": passportId, "clientIP": clientIp, "authAppId": authAppId, "timestamp": timeNow, "sign": signStr}
	jsonValue, _ := json.Marshal(values)

	// 设置默认 httpclient 的超时时间
	http.DefaultClient.Timeout = 5 * time.Second
	resp, err := http.Post(apiUrl, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	infoByte, _ := io.ReadAll(resp.Body)
	infoStr := string(infoByte)
	return infoStr, nil
}

func SignWithMD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	ret := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return ret
}
