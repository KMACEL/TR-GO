package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	listen = flag.String("listen", ":8081", "listen address")
)

func main() {
	flag.Parse()
	log.Printf("listening on %q...", *listen)
	log.Fatal(http.ListenAndServe(*listen, http.FileServer(http.Dir("./wasm"))))
}
