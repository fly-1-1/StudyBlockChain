package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

type ProofOfWork struct {
	block *Block
	//目标值
	target *big.Int
}

// NewProofOfWork 创建pow函数
func NewProofOfWork(block *Block) *ProofOfWork {
	pow := ProofOfWork{
		block: block,
	}
	//指定难度值 string类型 需要进行转换
	targetStr := "0000f00000000000000000000000000000000000000000000000000000000000"
	temInt := big.Int{}
	temInt.SetString(targetStr, 16)
	pow.target = &temInt
	return &pow
}

// Run 提供不断计算哈希的函数
func (pow *ProofOfWork) Run() ([]byte, uint64) {

	var nonce uint64
	block := pow.block
	var hash [32]byte
	fmt.Println("开始挖矿...")
	for {
		// 拼装数据 (区块数据 变化的随机数)
		tmp := [][]byte{
			Uint64ToByte(block.Version),
			block.PrevHash,
			block.MerkelRoot,
			Uint64ToByte(block.TimeStamp),
			Uint64ToByte(block.Difficulty),
			Uint64ToByte(nonce),
			//只对区块头做哈希值 区块体通过merkelRoot产生影响
			//block.Data,
		}
		//二维切片数组连接 返回一位数组切片
		blockInfo := bytes.Join(tmp, []byte{})
		// 哈希运算
		hash = sha256.Sum256(blockInfo)
		// 与pow中的target进行比较
		temInt := big.Int{}
		temInt.SetBytes(hash[:])
		//比较当前的hash值与目标哈希值,如果当前hash值小于目标hash说明找到了,否则继续
		if temInt.Cmp(pow.target) == -1 {
			//a 找到了 退出返回
			fmt.Printf("挖矿成功! hash: %x,nonce: %d\n", hash, nonce)
			//break
			return hash[:], nonce
		} else {
			//b 没找到,继续找,随机数+1
			nonce++
		}

	}
	//TODO
	//return []byte("HelloWorld"), 10

}
