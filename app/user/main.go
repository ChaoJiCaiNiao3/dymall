package main

import (
	user "github.com/ChaoJiCaiNiao3/dymall/app/user/kitex_gen/user/userservice"
	"log"
)

func main() {
	svr := user.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
