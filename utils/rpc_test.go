package utils

import (
	"encoding/json"
	"fmt"
	"testing"
)

var (
	walletUrl = "http://192.168.1.165:50301"
)

func TestGetAssetInfo(t *testing.T) {
	resp, err := GetAssetInfo(walletUrl, "TT")
	if err != nil {
		t.Fatal(err.Error())
	}

	respBody, err := json.Marshal(resp)
	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println(string(respBody))
}

func TestGetAddressBalance(t *testing.T) {
	resp, err := GetAddressBalance(walletUrl, "mfhGJnP5T7A5kYDJNxnHozxrVzC7WKHzKs", "1.3.0")
	if err != nil {
		t.Fatal(err.Error())
	}

	respBody, err := json.Marshal(resp)
	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println(string(respBody))
}

func TestGetChainId(t *testing.T) {
	resp, err := GetChainId(walletUrl)
	if err != nil {
		t.Fatal(err.Error())
	}

	respBody, err := json.Marshal(resp)
	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println(string(respBody))
}

func TestGetRefBlockInfo(t *testing.T) {
	refBlockNum, refBlockPrefix, err := GetRefBlockInfo(walletUrl)
	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println(refBlockNum)
	fmt.Println(refBlockPrefix)
}

func TestGetBindingAccount(t *testing.T) {
	result, err := GetBindingAccount(walletUrl, "mnUbdaJcTiBUARHGMZpQ5dVkrcj1XUMame", "BTC")
	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println(result)
}

func TestGetDepositAddress(t *testing.T) {
	result, err := GetDepositAddress(walletUrl, "BTC")
	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println(*result)
}

func TestGetAddressMintInfo(t *testing.T) {
	result, err := GetAddressMintInfo(walletUrl, "mnUbdaJcTiBUARHGMZpQ5dVkrcj1XUMame", "OO")
	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println(*result)
}
