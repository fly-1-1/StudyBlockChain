package main

import "fmt"

func main() {

	//block := NewBlock("A->B 1 BTC", []byte{})

	bc := NewBlockChain()
	bc.AddBlock("A->B 20 BTC")
	bc.AddBlock("A->B 30 BTC")
	it := bc.NewIterator()
	//调用迭代器 返回每一个区块数据
	for {
		//返回区块 左移
		block := it.Next()
		fmt.Printf("=========================================\n\n")
		fmt.Printf("前区块Hash:%x\n", block.PrevHash)
		fmt.Printf("当前区块Hash:%x\n", block.Hash)
		fmt.Printf("区块数据:%s\n", block.Data)
		if len(block.PrevHash) == 0 {
			break
		}
	}
}
