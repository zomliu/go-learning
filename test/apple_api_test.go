package test

import (
	"context"
	"testing"

	"github.com/richzw/appstore"
)

// https://developer.apple.com/documentation/appstoreserverapi

// [魔域1]的参数
func TestOrderLookUp(t *testing.T) {
	const accountPrivateKey = `
-----BEGIN PRIVATE KEY-----
MIGTAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBHkwdwIBAQQgwNP+18hC21FWalnKVd+GYXnoAk08fusicP5wdG3qOhagCgYIKoZIzj0DAQehRANCAAR4dtiOvs3Ct4bliGiCJqGUnfv4Hg7L3eZeIDUN6qMHtSrj2Q5qXngRlyis997yw37N3U0AdkCnVGrFwI5CIRKH
-----END PRIVATE KEY-----    
	`
	c := &appstore.StoreConfig{
		KeyContent: []byte(accountPrivateKey),
		KeyID:      "88SW88F499",
		BundleID:   "com.mysy.cn",
		Issuer:     "1a4d9eac-cbca-4e53-aeb2-bf72ec56f272",
		Sandbox:    false,
	}
	a := appstore.NewStoreClient(c)

	orderList := []string{
		"ML4H9STD3X",
		"ML4H9STBJW",
		"ML4H9ST9ZX",
		"MTFHTHH4XJ",
		"MTFHTHH4K8",
		"MTFHTHH4GB",
		"MTFHTHH334",
		"MTFHTHH2YJ",
		"MNK1VMo9VY",
		"MNK1VM09VT",
	}

	for i := range orderList {
		rsp, err := a.LookupOrderID(context.TODO(), orderList[i])

		if err != nil {
			t.Fatal(err)
		}

		orders, err := a.ParseSignedTransactions(rsp.SignedTransactions)
		if err != nil {
			t.Fatal(err)
		}
		for _, order := range orders {
			t.Logf("%s - %s", orderList[i], order.AppAccountToken)
		}
	}
}
