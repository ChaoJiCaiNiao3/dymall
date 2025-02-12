package main

import (
	cart "github.com/ChaoJiCaiNiao3/dymall/app/cart/kitex_gen/cart/cartservice"
	"log"
)

func main() {
	svr := cart.NewServer(new(CartServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
