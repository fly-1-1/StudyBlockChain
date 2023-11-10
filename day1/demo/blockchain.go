package main

import (
	"BlockChain/bolt"
	"log"
)

// BlockChain 引入区块链
type BlockChain struct {
	//定义一个区块链数组
	//blocks []*Block
	db   *bolt.DB
	tail []byte //存储最后一个区块的哈希
}

const blockChainDB = "blockChain.db"
const blockBucket = "blockBucket"

// NewBlockChain 定义一个区块链
func NewBlockChain() *BlockChain {

	//return &BlockChain{
	//	blocks: []*Block{genesisBlock},
	//}

	//最后一个区块的哈希 从数据库读出
	var lastHash []byte
	//打开数据库
	db, err := bolt.Open(blockChainDB, 0600, nil)
	//defer db.Close()
	if err != nil {
		log.Panic("打开数据库失败!")
	}
	//找到抽屉
	db.Update(func(tx *bolt.Tx) error {
		//找到抽屉bucket 如果没有就创建
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			//没有抽屉 需要创建抽屉
			bucket, err = tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				log.Panic("bucket创建失败!")
			}
			//创建一个创世块,并作为第一个区块添加到区块链中
			//hash作为key block字节流数据作为value
			genesisBlock := GenesisBlock()
			bucket.Put(genesisBlock.Hash, genesisBlock.Serialize())
			bucket.Put([]byte("LastHashKey"), genesisBlock.Hash)
			lastHash = genesisBlock.Hash

			//Test
			//blockBytes := bucket.Get(genesisBlock.Hash)
			//block := DeSerialize(blockBytes)
			//fmt.Printf("block info :%s\n", block)

		} else {
			lastHash = bucket.Get([]byte("LastHashKey"))
		}
		return nil
	})
	return &BlockChain{db, lastHash}
}

// GenesisBlock 创世块
func GenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// AddBlock 添加区块
func (bc *BlockChain) AddBlock(data string) {
	//获取前区块的哈希
	db := bc.db         //区块链数据库
	lastHash := bc.tail //最后一个区块的哈希

	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("bucket 不应该为空,请检查!")
		}
		//创建新的区块
		block := NewBlock(data, lastHash)
		//完成数据的添加
		//添加到区块db中
		bucket.Put(block.Hash, block.Serialize())
		bucket.Put([]byte("LastHashKey"), block.Hash)
		//更新内存的区块链(最后的尾巴需要更新)
		bc.tail = block.Hash
		return nil
	})

}
