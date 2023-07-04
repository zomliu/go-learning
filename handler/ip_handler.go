package handler

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var ip_file_path = "/Users/leon/Downloads/temp/ip_list.csv"
var ip_file_output_path = "/Users/leon/Downloads/temp/ip_list_output.csv"

func QueryIPLocal() {
	ipList := readIPData()
	for _, ip := range ipList {
		//fmt.Println(ip)
		// fmt.Println(queryIPLocal(ip))
		writeIPToFile(ip, queryIPLocal(ip))
	}
}

func queryIPLocal(ip string) string {
	t := "http://opendata.baidu.com/api.php?query=%s&co=&resource_id=6006&oe=utf8"
	url := fmt.Sprintf(t, ip)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var ipLocation IPLocation
	if err := json.NewDecoder(resp.Body).Decode(&ipLocation); err != nil {
		panic(err)
	}

	return ipLocation.Data[0].Location

}

// Read ip data from ip_file_path
func readIPData() []string {
	file, err := os.Open(ip_file_path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var ipList []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ipList = append(ipList, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return ipList
}

func writeIPToFile(ip, location string) {
	file, err := os.OpenFile(ip_file_output_path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err := file.WriteString(ip + "," + location + "\n"); err != nil {
		panic(err)
	}
}

/*
	{
    "status": "0",
    "t": "",
    "set_cache_time": "",
    "data": [
        {
            "ExtendedLocation": "",
            "OriginQuery": "117.136.12.79",
            "appinfo": "",
            "disp_type": 0,
            "fetchkey": "117.136.12.79",
            "location": "广东省广州市 移动",
            "origip": "117.136.12.79",
            "origipquery": "117.136.12.79",
            "resourceid": "6006",
            "role_id": 0,
            "shareImage": 1,
            "showLikeShare": 1,
            "showlamp": "1",
            "titlecont": "IP地址查询",
            "tplt": "ip"
        }
    ]
}
*/
// generate a struct from json
type IPLocation struct {
	Status       string `json:"status"`
	T            string `json:"t"`
	SetCacheTime string `json:"set_cache_time"`
	Data         []struct {
		ExtendedLocation string `json:"ExtendedLocation"`
		OriginQuery      string `json:"OriginQuery"`
		Appinfo          string `json:"appinfo"`
		DispType         int    `json:"disp_type"`
		Fetchkey         string `json:"fetchkey"`
		Location         string `json:"location"`
		Origip           string `json:"origip"`
		Origipquery      string `json:"origipquery"`
		Resourceid       string `json:"resourceid"`
		RoleId           int    `json:"role_id"`
		ShareImage       int    `json:"shareImage"`
		ShowLikeShare    int    `json:"showLikeShare"`
		Showlamp         string `json:"showlamp"`
		Titlecont        string `json:"titlecont"`
		Tplt             string `json:"tplt"`
	} `json:"data"`
}
