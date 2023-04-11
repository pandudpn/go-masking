package main

import (
	"fmt"
	"github.com/pandudpn/go-masking"
)

func main() {
	var pwd = "123456"
	
	result := masking.String(masking.Password, pwd)
	fmt.Println(result)
	
	var cc = "1234343204033013"
	result = masking.String(masking.CreditCard, cc)
	fmt.Println(result)
	
	var email = "pandu@pandudpn.id"
	result = masking.String(masking.Email, email)
	fmt.Println(result)
	
	var phone = "6283875181609"
	result = masking.String(masking.PhoneNumber, phone)
	fmt.Println(result)
	
	var d = "ab"
	result = masking.String(5, d)
	fmt.Println(result)
}
