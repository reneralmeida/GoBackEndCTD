package main

import (
	"fmt"
	"time"
)

func main() {

	refund := make(chan string)
	payment := make(chan string)

	go func() {
		for {
			time.Sleep(time.Second / 2) // 500ms
			payment <- "pagamento recebido"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second) // 1s
			refund <- "reembolso realizado"
		}
	}()

	// mensagens de pagamento
	go func() {
		for message := range payment {
			fmt.Println(message)
		}
	}()

	// mensagens de reembolso
	go func() {
		for message := range refund {
			fmt.Println(message)
		}
	}()

	for {
	}
}
