package main

import (
	"fmt"

	"github.com/libp2p/go-libp2p-core/peer"
)

func main() {
	stringToParse := "bafykbzacecesdmus2x37xyvtxdwlpx3vu5cjecxed473f66fezh6tdoccqhqq"
	fmt.Println("Parsing", stringToParse, "...")
	peerIdStruct, err := peer.Decode(stringToParse)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println(peerIdStruct.Pretty())
	}
}
