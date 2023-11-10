package main

import "fmt"

func (cli *CLI) PrintBlockChain() {

	bc := cli.bc
	it := bc.NewIterator()
	for {
		//返回区块 左移
		block := it.Next()
		fmt.Printf("=========================================\n\n")
		fmt.Printf("version:%d\n", block.Version)
		fmt.Printf("前区块Hash:%x\n", block.PrevHash)
		fmt.Printf("MerkelRoot:%x\n", block.MerkelRoot)
		fmt.Printf("时间戳:%d\n", block.TimeStamp)
		fmt.Printf("难度值:%d\n", block.Difficulty)
		fmt.Printf("随机数:%d\n", block.Nonce)
		fmt.Printf("当前区块Hash:%x\n", block.Hash)
		fmt.Printf("区块数据:%s\n", block.Data)
		if len(block.PrevHash) == 0 {
			break
		}
	}
}

func (cli *CLI) AddBlock(data string) {
	cli.bc.AddBlock(data)
	fmt.Println("添加区块成功")
}
