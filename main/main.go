package main

import (
	"fmt"

	"github.com/elfincafe/gogetssl"
)

func main() {

	// Token
	// api := &gogetssl.Auth{LiveAPI: "https://my.gogetssl.com/api/"}
	// res, err := api.Auth("register@m.tsukinoha.jp", "7f30a82db2dc58a39bffcd4cd8662441")

	token := "b5df270e1b167b585a968f12935cfdf6526dfc2b"

	// WebServers
	api := &gogetssl.WebServers{LiveAPI: "https://my.gogetssl.com/api/", Key: token}
	res, err := api.Get(2)

	// Status

	fmt.Println(res, err)
}
