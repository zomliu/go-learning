package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"unicode"
)

type (
	People struct {
		Name string
		Age  int64
	}
)

func TestSwitch(t *testing.T) {
	// get me a httptest.NewServer expmple
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// wirete http 200 to response
		w.WriteHeader(http.StatusOK)
		// write hello world to response
		w.Write([]byte("Hello World"))
	}))
	defer server.Close()

	server.Client()
}

func TestMapModify(t *testing.T) {
	m := make(map[string]People)

	p1 := People{Name: "aa", Age: 20}

	fmt.Printf("%p \n", &p1)

	m[p1.Name] = p1

	p2 := m[p1.Name]
	fmt.Printf("%p \n", &p2)
}

func TestCurrencyCode(t *testing.T) {
	str := `
	[{
		"id": 34,
		"region": "瑞士",
		"currencyCode": "CHF",
		"currencyName": "法郎",
		"createTime": 1450343830000,
		"appleMatchRegex": null
	}, {
		"id": 4,
		"region": "墨西哥",
		"currencyCode": "MXN",
		"currencyName": "墨西哥比索",
		"createTime": 1450343829000,
		"appleMatchRegex": null
	}, {
		"id": 44,
		"region": "沙特阿拉伯",
		"currencyCode": "SAR",
		"currencyName": "里亚尔",
		"createTime": 1450343830000,
		"appleMatchRegex": null
	}, {
		"id": 45,
		"region": "南非",
		"currencyCode": "ZAR",
		"currencyName": "南非兰特",
		"createTime": 1450343830000,
		"appleMatchRegex": null
	}, {
		"id": 40,
		"region": "印度",
		"currencyCode": "INR",
		"currencyName": "卢比",
		"createTime": 1450343830000,
		"appleMatchRegex": null
	}, {
		"id": 49,
		"region": "越南",
		"currencyCode": "VND",
		"currencyName": "越南盾",
		"createTime": 1472214725000,
		"appleMatchRegex": null
	}, {
		"id": 1,
		"region": "中国",
		"currencyCode": "CNY",
		"currencyName": "人民币",
		"createTime": 1450343829000,
		"appleMatchRegex": null
	}, {
		"id": 50,
		"region": "泰国",
		"currencyCode": "THB",
		"currencyName": "泰铢",
		"createTime": 1476843601000,
		"appleMatchRegex": null
	}, {
		"id": 35,
		"region": "澳大利亚",
		"currencyCode": "AUD",
		"currencyName": "澳元",
		"createTime": 1450343830000,
		"appleMatchRegex": null
	}, {
		"id": 3,
		"region": "韩国",
		"currencyCode": "KRW",
		"currencyName": "韩元",
		"createTime": 1450343829000,
		"appleMatchRegex": null
	}, {
		"id": 42,
		"region": "以色列",
		"currencyCode": "ILS",
		"currencyName": "新谢克尔",
		"createTime": 1450343830000,
		"appleMatchRegex": null
	}, {
		"id": 37,
		"region": "日本",
		"currencyCode": "JPY",
		"currencyName": "日元",
		"createTime": 1450343830000,
		"appleMatchRegex": null
	}, {
		"id": 6,
		"region": "英国",
		"currencyCode": "GBP",
		"currencyName": "英镑",
		"createTime": 1450343829000,
		"appleMatchRegex": null
	}, {
		"id": 41,
		"region": "印度尼西亚",
		"currencyCode": "IDR",
		"currencyName": "盾",
		"createTime": 1450343830000,
		"appleMatchRegex": null
	}, {
		"id": 46,
		"region": "土耳其",
		"currencyCode": "TRY",
		"currencyName": "新里拉",
		"createTime": 1450343830000,
		"appleMatchRegex": null
	}, {
		"id": 43,
		"region": "俄罗斯",
		"currencyCode": "RUB",
		"currencyName": "卢布",
		"createTime": 1450343830000,
		"appleMatchRegex": null
	}, {
		"id": 38,
		"region": "香港",
		"currencyCode": "HKD",
		"currencyName": "港币",
		"createTime": 1450343830000,
		"appleMatchRegex": null
	}, {
		"id": 39,
		"region": "台湾",
		"currencyCode": "TWD",
		"currencyName": "台币",
		"createTime": 1450343830000,
		"appleMatchRegex": null
	}, {
		"id": 47,
		"region": "阿拉伯联合酋长国",
		"currencyCode": "AED",
		"currencyName": "迪拉姆",
		"createTime": 1450343830000,
		"appleMatchRegex": null
	}, {
		"id": 7,
		"region": "塞浦路斯,爱沙尼亚,匈牙利,荷兰,立陶宛,希腊,斯洛伐克,奥地利,芬兰,德国,波兰,马耳他,葡萄牙,罗马尼亚,捷克共和国,斯洛文尼亚,拉脱维亚,法国,西班牙,意大利,爱尔兰,保加利亚,比利时,卢森堡",
		"currencyCode": "EUR",
		"currencyName": "欧元",
		"createTime": 1450343829000,
		"appleMatchRegex": null
	}, {
		"id": 32,
		"region": "丹麦",
		"currencyCode": "DKK",
		"currencyName": "克朗",
		"createTime": 1450343830000,
		"appleMatchRegex": null
	}, {
		"id": 2,
		"region": "美国",
		"currencyCode": "USD",
		"currencyName": "美元",
		"createTime": 1450343829000,
		"appleMatchRegex": null
	}, {
		"id": 5,
		"region": "加拿大",
		"currencyCode": "CAD",
		"currencyName": "加拿大元",
		"createTime": 1450343829000,
		"appleMatchRegex": null
	}, {
		"id": 51,
		"region": "马来西亚",
		"currencyCode": "MYR",
		"currencyName": "马来西亚林吉特",
		"createTime": 1536634278000,
		"appleMatchRegex": null
	}, {
		"id": 33,
		"region": "挪威",
		"currencyCode": "NOK",
		"currencyName": "克朗",
		"createTime": 1450343830000,
		"appleMatchRegex": null
	}, {
		"id": 48,
		"region": "新加坡",
		"currencyCode": "SGD",
		"currencyName": "新加坡元",
		"createTime": 1450343831000,
		"appleMatchRegex": null
	}, {
		"id": 31,
		"region": "瑞典",
		"currencyCode": "SEK",
		"currencyName": "克朗",
		"createTime": 1450343830000,
		"appleMatchRegex": null
	}, {
		"id": 36,
		"region": "新西兰",
		"currencyCode": "NZD",
		"currencyName": "新西兰币",
		"createTime": 1450343830000,
		"appleMatchRegex": null
	}]	
	`

	var list []struct {
		Id           int64
		Region       string
		CurrencyCode string `json:"currencyCode"`
		CurrencyName string `json:"currencyName"`
	}
	err := json.Unmarshal([]byte(str), &list)
	if err != nil {
		fmt.Printf("json error :%v", err)
		return
	}

	for i := range list {
		fmt.Printf("%s,", list[i].CurrencyCode)
	}
	t.Log("OK")
}

func BenchmarkDemo(b *testing.B) {
	for i := 1; i < b.N; i++ {
		demo1()
	}
}

func demo1() {
	ip := "192 168 1 1"

	s := strings.Fields(ip)
	fmt.Printf("lenght: %d, value: %v", len(s), s)
}

func TestSlice(t *testing.T) {

	list := make([]int, 10)

	for i := range list {
		list[i] = i
	}
	fmt.Printf("the size is : %d , the cap is %d \n", len(list), cap(list))
	fmt.Printf("the print address is : %p \n", list)
	list = list[:0]
	fmt.Printf("the size is : %d, the cap is %d  \n", len(list), cap(list))
	fmt.Printf("the print address is : %p \n", list)
	list = append(list, 100)
	fmt.Printf("the size is : %d, the cap is %d  \n", len(list), cap(list))
	fmt.Printf("the print address is : %p \n", list)

}

func TestNumber(t *testing.T) {
	str := "13716742307@163.321"

	fmt.Println()

	if IsDigit(str) {
		t.Log(str + " is a number")
	} else {
		t.Log(str + " is not a number")
	}
	fmt.Println()

}

func IsDigit(str string) bool {
	for _, r := range str {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
