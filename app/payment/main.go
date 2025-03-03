package main

import (
	"log"
	"net"

	payment "github.com/ChaoJiCaiNiao3/dymall/app/payment/kitex_gen/payment/paymentservice"
	"github.com/cloudwego/kitex/server"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", ":8890")
	if err != nil {
		log.Fatal(err)
	}

	svr := payment.NewServer(
		new(PaymentServiceImpl),
		server.WithServiceAddr(addr),
	)

	if err := svr.Run(); err != nil {
		log.Fatal("服务运行失败:", err)
	}
}
