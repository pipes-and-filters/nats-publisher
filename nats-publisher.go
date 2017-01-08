package main

import (
	"log"

	nats "github.com/nats-io/go-nats"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Print(err)
	}
	err = nc.Publish("footest", []byte("Hello World"))
	if err != nil {
		log.Print(err)
	}
	nc.Close()
}
