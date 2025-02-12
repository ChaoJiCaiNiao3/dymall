package main

import (
	order "github.com/ChaoJiCaiNiao3/dymall/app/order/kitex_gen/order/orderservice"
	"log"
)

func main() {
	svr := order.NewServer(new(OrderServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
