package main

import (
	"fmt"
)

func main() {

	//block := NewBlock("A->B 1 BTC", []byte{})

	bc := NewBlockChain()
	bc.AddBlock("A->B 20 BTC")
	bc.AddBlock("A->B 30 BTC")
	for i, block := range bc.blocks {
		fmt.Printf("前区块height:%d\n", i)
		fmt.Printf("前区块Hash:%x\n", block.PrevHash)
		fmt.Printf("当前区块Hash:%x\n", block.Hash)
		fmt.Printf("区块数据:%s\n", block.Data)
	}

}
