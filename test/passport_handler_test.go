package test

import (
	"crypto/sha256"
	"demo/handler"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestGetPassportDetail(t *testing.T) {
	//s, err := handler.DoCallZhuHaiAPIForCertificate("zhaoweiyan5233@163.com", "dev")//ksa_nrf8knwg395p3nyp3tkdsiwt3so
	s, _, err := handler.GetPassportDetail("ksa_ccy6knwg395p3nyp3tkdzix93no", "dev")
	if err != nil {
		t.Error(err)
	}
	t.Log(s)
}

func TestGetPassportOpenId(t *testing.T) {
	//account := "zhaoweiyan5233@163.com"
	account := "18171242620"
	s, err := handler.GetPassportOpenId(account, "dev")

	if err != nil {
		t.Error(err)
	}
	t.Log(s)
}

func TestGetPassportVerifyInfo(t *testing.T) {
	info, err := handler.GetPassportVerifyInfo("ksa_ccy6knwg395p3nyp3tkdzix93no", "dev")
	if err != nil {
		t.Error(err)
	}
	t.Log(info)
}

func TestSha256(t *testing.T) {
	input := "3db5dc8580c24f0e8543c2b71a5cb9cc"
	hash := sha256.Sum256([]byte(input))
	signature := hex.EncodeToString(hash[:])
	fmt.Println(signature)
}
