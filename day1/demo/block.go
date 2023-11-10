package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/gob"
	"log"
	"time"
)

// Block 定义结构
type Block struct {
	//版本号
	Version uint64
	//前区块哈希
	PrevHash []byte
	//Merkel根 Hash v4版本 //TODO
	MerkelRoot []byte
	//时间戳
	TimeStamp uint64
	//难度值
	Difficulty uint64
	//随机数(挖矿需要找的数据)
	Nonce uint64
	//当前区块哈希(BTC区块中无当前区块哈希)
	Hash []byte
	//数据
	Data []byte
}

// Uint64ToByte 将uint转为[]byte
func Uint64ToByte(num uint64) []byte {
	//TODO
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}

// NewBlock 创建区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		Version:    00,
		PrevHash:   prevBlockHash,
		MerkelRoot: []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Difficulty: 00, //无效值
		Nonce:      00, //无效值
		Hash:       []byte{},
		Data:       []byte(data),
	}
	//block.SetHash()
	//创建pow对象
	pow := NewProofOfWork(&block)
	hash, nonce := pow.Run()

	//根据挖矿结果对区块数据进行更新
	block.Hash = hash
	block.Nonce = nonce

	return &block
}

// Serialize 序列化
func (block *Block) Serialize() []byte {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(&block)
	if err != nil {
		log.Panic("编码出错")
	}
	return buffer.Bytes()
}

// DeSerialize 反序列化
func DeSerialize(data []byte) *Block {
	//TODO
	decoder := gob.NewDecoder(bytes.NewReader(data))

	var block Block
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic("解码出错")
	}

	return &block
}

// SetHash 生成哈希
func (block *Block) SetHash() {

	tmp := [][]byte{
		Uint64ToByte(block.Version),
		block.PrevHash,
		block.MerkelRoot,
		Uint64ToByte(block.TimeStamp),
		Uint64ToByte(block.Difficulty),
		Uint64ToByte(block.Nonce),
		block.Data,
	}
	//二维切片数组连接 返回一位数组切片
	blockInfo := bytes.Join(tmp, []byte{})
	//2 sha256
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}
