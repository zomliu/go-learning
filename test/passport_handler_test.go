package test

import (
	"demo/handler"
	"testing"
)

func TestGetPassportDetail(t *testing.T) {
	//s, err := handler.DoCallZhuHaiAPIForCertificate("zhaoweiyan5233@163.com", "dev")//ksa_nrf8knwg395p3nyp3tkdsiwt3so
	s, err := handler.GetPassportDetail("ksa_ccy6knwg395p3nyp3tkdzix93no", "dev")
	if err != nil {
		t.Error(err)
	}
	t.Log(s)
}

func TestGetPassportOpenId(t *testing.T) {
	//account := "zhaoweiyan5233@163.com"
	account := "18641932345"
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
