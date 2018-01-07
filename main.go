package main

import (
	"fmt"
	"github.com/yfujita/sample-block-chain/block_chain"
	"math/rand"
	"time"
)


func main() {
	for i := 0; i<1000000; i++ {
		rand.Seed(time.Now().UnixNano())
		nonce := rand.Intn(10000000000000)
		fmt.Printf("nonce:%d\n", nonce)
		block, err := block_chain.CreateBlock("aaaa", "aa", nonce)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Hash=" + block_chain.HashString(block.Hash))
			break
		}
	}
}