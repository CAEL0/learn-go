package main

import (
	"fmt"

	"github.com/CAEL0/learn-go/accounts"
)

func main() {
	account := accounts.NewAccount("CAELO")
	fmt.Println(account)
}