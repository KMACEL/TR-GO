package main

import (
	"fmt"

	"github.com/vmihailenco/msgpack"
)

type Item struct {
	Foo string
}

func main() {
	b, err := msgpack.Marshal(&Item{Foo: "bar"})
	if err != nil {
		panic(err)
	}

	fmt.Println("Marshal Hali : ", string(b))

	var item Item
	err = msgpack.Unmarshal(b, &item)
	if err != nil {
		panic(err)
	}
	fmt.Println(item.Foo)
}
