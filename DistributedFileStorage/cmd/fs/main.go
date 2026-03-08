package main

import (
	"distributed_file_storage/p2p"
	"log"
)

func main() {
	tr := p2p.NewTCPTransport(":8080")
	err := tr.ListenAndAccept()
	if err != nil {
		log.Fatal(err)
		return
	}
}
