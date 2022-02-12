package api

import (
	"fmt"
	"strconv"
	"testing"
)

const token = "a5b78192-a188-45e6-bcfb-8c096fb02914"

func TestRequestOtp(t *testing.T) {
	req, _ := RequestOtp("083897007021")
	fmt.Println(req)
}
func TestLoginOtp(t *testing.T) {
	req, _ := LoginOtp("083897007022", "ABPXSI")
	//_ = req
	fmt.Println(req)
}
func TestPackageBuy(t *testing.T) {
	req, _ := BuyPackage(token, "3212251")
	_ = req
	//fmt.Println(req )
}
func TestClaim(t *testing.T) {

	r := make(chan string)

	for i := 0000000; i < 9999999; i++ {

		str := strconv.Itoa(i)
		req, _ := BuyPackage(token, str)
		fmt.Println(req)

	}
	fmt.Println(r)
}
