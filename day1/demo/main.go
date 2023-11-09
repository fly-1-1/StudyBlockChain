package main

import "fmt"

// Block 定义结构
type Block struct {
	//前区块哈希
	PrevHash []byte
	//当前区块哈希
	Hash []byte
	//数据
	Date []byte
}

// NewBlock 创建区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		PrevHash: prevBlockHash,
		Hash:     []byte{}, //TODO
		Date:     []byte(data),
	}
	return &block
}

func main() {
	block := NewBlock("A->B 1 BTC", []byte{})
	fmt.Printf("前区块Hash:%x\n", block.PrevHash)
	fmt.Printf("当前区块Hash:%x\n", block.Hash)
	fmt.Printf("区块数据:%s\n", block.Date)
}
