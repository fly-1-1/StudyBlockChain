package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

const reward = 12.5

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

// NewCoinbaseTX  提供创建交易方法 (挖矿交易)
func NewCoinbaseTX(address string, data string) *Transaction {
	//挖矿交易的特点 1只有一个input 2 无需引用交易ID 3 无需引用index 4 矿工由于挖矿时无需指定签名 这个sig字段一般填写矿池的名字
	input := TXInput{[]byte{}, -1, data}
	output := TXOutput{reward, address}
	//对于挖矿交易来说 只有一个input和output
	tx := Transaction{[]byte{}, []TXInput{input}, []TXOutput{output}}
	tx.SetHash()
	return &tx
}

//3 创建挖矿交易

//4 根据交易调整程序
