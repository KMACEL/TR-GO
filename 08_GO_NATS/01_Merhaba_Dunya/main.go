package main

import (
	"fmt"

	"github.com/nats-io/go-nats"
)

var (
	// Benim sunucuma kurduğum NATS server. Kurması çok kolay. Siz kendi bilgisayarınızda da kurabilirsiniz.
	// Kurmak için ;
	// 		go get github.com/nats-io/gnatsd
	// Çalıştırmak için ;
	//		gnatsd
	// Uyarı : Eğer kendi bilgisayarınız üzerinden test yapacaksanız, mke.systems yerine localhost yazmanız gerekmektedir
	url     = "nats://mke.systems:4222"
	subject = "test"
)

func main() {
	nc, _ := nats.Connect(url)
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()

	// Publisher - Gönderen
	c.Publish(subject, "Merhaba Dunya")

	// Async Subscriber - Asenkron Alıcı
	c.Subscribe(subject, func(s string) {
		fmt.Printf("Gelen Mesaj: %s\n", s)
	})

	fmt.Scanln()
}
