package main

import (
	"log"

	"github.com/sharjil07/distributed-file-system/p2p"
)

func main() {
	tr := p2p.NewTCPtransport(":3000")

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
