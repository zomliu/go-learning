package test

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"testing"
)

// Oversea
func TestGeO(t *testing.T) {
	const (
		// 武林闲侠 Prod
		appId  = "1239"
		appKey = "6682570acb5647d3a5a9611b53955783"
		planId = "233"

		// 武林闲侠 Dev
		//appKey = "0e42a253d75c45bf8c0e13884d80f34d"
		//planId = "232"

		// 音舞 Prod
		//appId  = "1240"
		//appKey = "35a88661c1d55f76782d623c440e4869"
		//planId = "212"
	)

	const (
		channelCode = "apple"
		region      = "oversea"
		host        = "https://omni-build.seayoo.io/omni/build/app/release/channel/ios?app_id=%s&channel_code=%s&debug_mode=true&region=%s&release_version_id=%s&sign=%s"
		//host = "https://omni-build.dev.seayoo.io/omni/build/app/release/channel/ios?app_id=%s&channel_code=%s&debug_mode=true&region=%s&release_version_id=%s&sign=%s"
	)

	paramMap := map[string]string{
		"app_id":             appId,
		"channel_code":       channelCode,
		"debug_mode":         "true",
		"region":             region,
		"release_version_id": planId,
	}
	keys := []string{"app_id", "channel_code", "debug_mode", "region", "release_version_id", "key"}

	// 参数名称排序
	sort.Strings(keys)

	// 拼接签名字符串
	var r strings.Builder
	for _, key := range keys {
		r.WriteString(key)
		r.WriteString("=")
		if key == "key" {
			r.WriteString(appKey)
		} else {
			r.WriteString(paramMap[key])
		}
		r.WriteString(";")
	}
	ret := r.String()
	// 计算 md5 值
	signWithMD5 := signWithMD5(ret)

	url := fmt.Sprintf(host, appId, channelCode, region, planId, signWithMD5)
	//fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		t.Error("request error", err)
		return
	}
	defer resp.Body.Close()

	b, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		t.Error("read http response error ", err2)
		return
	}
	fmt.Println("_____________ Result _________________")
	fmt.Println(string(b))
}

// Domestic
func TestGe(t *testing.T) {
	const (
		// is dev env or not
		devEnv      = false
		channelCode = "ios_jinshanApple"
		//channelCode = "jinshan_prom"
		region = "domestic"
	)
	host := "https://omni-build.xgsdk.seayoo.com/omni/build/app/release/channel?app_id=%s&channel_code=%s&debug_mode=true&region=domestic&release_version_id=%s&sign=%s"
	if devEnv {
		host = "https://omni-build.xgsdk.dev.seayoo.com/omni/build/app/release/channel?app_id=%s&channel_code=%s&debug_mode=true&region=domestic&release_version_id=%s&sign=%s"
	}

	const (
		// paopao-dev
		//appId  = "111111650"
		//appKey = "007b2e0fe58f4b798399ffbb6ce45ea8"
		//planId = "200011"

		// paopao
		//appId  = "111111650"
		//appKey = "007b2e0fe58f4b798399ffbb6ce45ea8"
		//planId = "200170"

		// moyu
		//appId  = "15985"
		//appKey = "25c9b673be984a78844a6a14417e6311"
		//planId = "200178"

		// 风暴魔域2
		//appId  = "111111624"
		//appKey = "bd4da789356a405dbbac6bf438345f98"
		//planId = "200016"

		// 西瓜测试--渠道测试and开发使用
		//appId  = "91000184"
		//appKey = "4f61d0f2339d4faca8730af702c28d2f"
		//planId = "200033"

		// 幻影战争
		//appId  = "111111634"
		//appKey = "e1856655015d4e8b973f8ea4cb040511"
		//planId = "200147"

		// 魔域2
		appId  = "200000000"
		appKey = "55a4fd2aa48c4273872eeeabb0680a73"
		planId = "200187"
	)

	paramMap := map[string]string{
		"app_id":             appId,
		"channel_code":       channelCode,
		"debug_mode":         "true",
		"region":             region,
		"release_version_id": planId,
	}
	keys := []string{"app_id", "channel_code", "debug_mode", "region", "release_version_id", "key"}

	// 参数名称排序
	sort.Strings(keys)

	// 拼接签名字符串
	var r strings.Builder
	for _, key := range keys {
		r.WriteString(key)
		r.WriteString("=")
		if key == "key" {
			r.WriteString(appKey)
		} else {
			r.WriteString(paramMap[key])
		}
		r.WriteString(";")
	}
	ret := r.String()
	// 计算 md5 值
	signWithMD5 := signWithMD5(ret)

	url := fmt.Sprintf(host, appId, channelCode, planId, signWithMD5)

	resp, err := http.Get(url)
	if err != nil {
		t.Error("request error", err)
		return
	}
	defer resp.Body.Close()

	b, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		t.Error("read http response error ", err2)
		return
	}
	fmt.Println("_____________ Result _________________")
	fmt.Println(string(b))
}

func signWithMD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	ret := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return ret
}
