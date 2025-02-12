package main

import (
	payment "github.com/ChaoJiCaiNiao3/dymall/app/payment/kitex_gen/payment/paymentservice"
	"log"
)

func main() {
	svr := payment.NewServer(new(PaymentServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
