package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)


// Transaction 1 定义交易结构
type Transaction struct {
	TXID      []byte     //交易ID
	TXInputs  []TXInput  //交易输入数组
	TXOutputs []TXOutput //交易输出数组
}

// TXInput 定义交易输入
type TXInput struct {
	//引用的交易ID
	TXid []byte
	//引用输出的output的索引
	Index int64
	//解锁脚本 用地址模拟
	Sig string
}

// TXOutput 定义交易输出
type TXOutput struct {
	//转账金额
	value float64
	//锁定脚本,用地址模拟
	PukKeyHash string
}

// SetHash 设置交易ID
func (tx *Transaction) SetHash() {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	data := buffer.Bytes()
	hash := sha256.Sum256(data)
	tx.TXID = hash[:]
}

//2 提供创建交易方法

//3 创建挖矿交易

//4 根据交易调整程序
