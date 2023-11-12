package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
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
	Value float64
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

// IsCoinBase 判断当前交易是否为挖矿交易
func (tx *Transaction) IsCoinBase() bool {
	//交易的input只有一个
	//交易id为空
	//交易的index为-1
	if len(tx.TXInputs) == 1 {
		input := tx.TXInputs[0]
		if bytes.Equal(input.TXid, []byte{}) && input.Index == -1 {
			return true
		}
	}
	return false
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

//创建普通交易

// 创建outputs
// 如果有零钱需要找零钱

func NewTransaction(from, to string, amount float64, bc *BlockChain) *Transaction {
	// 找到合理的utxo集合 map[string][]int64
	utxos, resValue := bc.FindNeedUTXOs(from, amount)
	if resValue < amount {
		fmt.Println("余额不足,交易失败")
		return nil
	}
	var inputs []TXInput
	var outputs []TXOutput
	// 将这些utxo逐一转为inputs
	for id, indexArray := range utxos {
		for _, i := range indexArray {
			input := TXInput{[]byte(id), int64(i), from}
			inputs = append(inputs, input)
		}
	}
	output := TXOutput{amount, to}
	outputs = append(outputs, output)
	if resValue > amount {
		outputs = append(outputs, TXOutput{resValue - amount, from})
	}
	tx := Transaction{[]byte{}, inputs, outputs}
	tx.SetHash()
	return &tx
}
