package main

import (
	"crypto/sha256"
	"fmt"
)

// Block 定义结构
type Block struct {
	//前区块哈希
	PrevHash []byte
	//当前区块哈希
	Hash []byte
	//数据
	Data []byte
}

// NewBlock 创建区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		PrevHash: prevBlockHash,
		Hash:     []byte{},
		Data:     []byte(data),
	}
	block.SetHash()

	return &block
}

// SetHash 生成哈希
func (block *Block) SetHash() {
	//1 拼装数据
	blockInfo := append(block.PrevHash, block.Data...)
	//2 sha256
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}

func main() {
	block := NewBlock("A->B 1 BTC", []byte{})
	fmt.Printf("前区块Hash:%x\n", block.PrevHash)
	fmt.Printf("当前区块Hash:%x\n", block.Hash)
	fmt.Printf("区块数据:%s\n", block.Data)
}
