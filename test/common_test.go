package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"
	"unicode"

	"github.com/spf13/cast"
)

type (
	People struct {
		Name string
		Age  int64
	}
)

func TestString(t *testing.T) {

	p1 := &People{Name: "aa", Age: 20}
	fmt.Printf("p1[%p] is %v \n", p1, p1)

	p1 = newPeople()
	fmt.Printf("p1[%p] is %v", p1, p1)
}

func newPeople() *People {
	return &People{Name: "bb", Age: 120}
}

func TestSringLength(t *testing.T) {
	str := "hello world"
	t.Logf("str length is %d", len(str))
}

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

func TestHttp(t *testing.T) {
	apiURl := "https://sdk-store.mlinkapp.com/game/order/query"

	appID := "3161875"
	cpOrderID := "3240271049849243"
	sign := "03e3a0b4f84973f0f0151d46709dec09"
	signType := "md5"
	ts := "1711490695269"

	// 创建表单数据
	formData := url.Values{}
	formData.Add("app_id", appID)
	formData.Add("cp_order_id", cpOrderID)
	formData.Add("sign", sign)
	formData.Add("sign_type", signType)
	formData.Add("ts", ts)

	// 将表单数据编码成 'application/x-www-form-urlencoded' 格式
	body := bytes.NewBufferString(formData.Encode())

	// 创建请求
	req, err := http.NewRequest(http.MethodPost, apiURl, body)
	if err != nil {
		panic(err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 发送请求
	client := &http.Client{Timeout: time.Second * 10} // 设置超时时间
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var respBody []byte
	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	log.Print(string(respBody))
}

func TestMap(t *testing.T) {
	// 21.3652155730043
	var data = struct {
		LatencyMs    int64
		Ecpm         float64
		EcpmCurrency string
	}{
		LatencyMs:    110,
		Ecpm:         0,
		EcpmCurrency: "CNY",
	}

	properties := map[string]interface{}{
		"ecpm":          cast.ToFloat64(strconv.FormatFloat(data.Ecpm, 'f', 3, 64)),
		"ecpm_currency": data.EcpmCurrency,
		"latency_ms":    data.LatencyMs,
	}

	for k, v := range properties {
		if t, ok := v.(string); ok && t == "" {
			delete(properties, k)
		} else if t, ok := v.(float64); ok && t == 0 {
			delete(properties, k)
		} else if t, ok := v.(int64); ok && t == 0 {
			delete(properties, k)
		}

	}
	t.Log(properties)
}

func TestT2(t *testing.T) {
	for i := range 10 {
		ttt(i + 1)
	}
}

func ttt(idx int) {
	fmt.Println("start one: ", idx)

	if idx%2 == 0 {
		return
	}
	defer func() {
		fmt.Println("end one", idx)
	}()

	fmt.Println("start two", idx)
}

func TestT3(t *testing.T) {
	uri := "/ads/v1/topon?user_id=1240425162510014&trans_id=42a2f14f12e0562e6387955f010ae4bc_4947086_1715054701108&reward_amount=1&reward_name=%E6%BF%80%E5%8A%B1%E5%A5%96%E5%8A%B1&placement_id=b65d874487d9f3&extra_data=018f5132b29c7f3b9de1520a79930281&network_firm_id=8&adsource_id=4947086&scenario_id=&sign=1ff110d93b9a696ff8cc56003747c0ba&ilrd=%7B%22id%22%3A%2242a2f14f12e0562e6387955f010ae4bc_4947086_1715054701108%22%2C%22publisher_revenue%22%3A0.007797289577999999%2C%22currency%22%3A%22CNY%22%2C%22country%22%3A%22CN%22%2C%22adunit_id%22%3A%22b65d874487d9f3%22%2C%22adunit_format%22%3A%22RewardedVideo%22%2C%22precision%22%3A%22exact%22%2C%22network_type%22%3A%22Network%22%2C%22network_placement_id%22%3A%224059031216112767%22%2C%22ecpm_level%22%3A0%2C%22segment_id%22%3A0%2C%22scenario_reward_name%22%3A%22reward_item%22%2C%22scenario_reward_number%22%3A1%2C%22custom_rule%22%3A%7B%22user_id%22%3A%223fc28b1b70e87ab0%22%7D%2C%22network_firm_id%22%3A8%2C%22adsource_id%22%3A%224947086%22%2C%22adsource_index%22%3A30%2C%22adsource_price%22%3A7.797289577999999%2C%22adsource_isheaderbidding%22%3A1%2C%22ext_info%22%3A%7B%22mp%22%3A-1%2C%22is_reward_ad%22%3Afalse%2C%22request_id%22%3A%22hjrevp7cm5nb401%22%2C%22gdt_trans_id%22%3A%2269fac0cebc1491f1c8c4386bbac8fb30%22%2C%22token%22%3A%22%22%7D%2C%22reward_custom_data%22%3A%22018f5132b29c7f3b9de1520a79930281%22%2C%22abtest_id%22%3A486335%2C%22user_load_extra_data%22%3A%7B%22user_id%22%3A%221240425162510014%22%2C%22user_custom_data%22%3A%22018f5132b29c7f3b9de1520a79930281%22%7D%2C%22placement_type%22%3A1%2C%22bid_floor%22%3A0%2C%22ad_source_type%22%3A1%2C%22ad_source_custom_ext%22%3A%22%22%2C%22network_name%22%3A%22Tencent+Ads%22%2C%22show_custom_ext%22%3A%22018f5139efd47230b040adeed56ea132%22%7D"

	n, e := url.QueryUnescape(uri)
	if e != nil {
		t.Error(e)
	}
	t.Log(n)
	n2, e2 := url.QueryUnescape(n)
	if e2 != nil {
		t.Error(e2)
	}
	t.Log(n2)

}
