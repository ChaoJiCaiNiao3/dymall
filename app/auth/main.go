package main

import (
	"log"

	auth "github.com/ChaoJiCaiNiao3/dymall/app/auth/kitex_gen/auth/authservice"
)

func main() {
	svr := auth.NewServer(new(AuthServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
