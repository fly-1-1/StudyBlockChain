package main

import (
	"BlockChain/bolt"
	"log"
)

type BlockChainIterator struct {
	db *bolt.DB
	// 游标用于索引
	currentHashPointer []byte
}

func (bc *BlockChain) NewIterator() *BlockChainIterator {

	return &BlockChainIterator{
		bc.db,
		//最初指向区块链的最后一个区块 随着Next的调用不断优化
		bc.tail,
	}
}

func (it *BlockChainIterator) Next() *Block {
	var block *Block
	it.db.View(func(tx *bolt.Tx) error {

		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("迭代器遍历bucket不应该为空,请检查")
		}
		blockTmp := bucket.Get(it.currentHashPointer)
		//解码
		block = DeSerialize(blockTmp)
		//游标左移
		it.currentHashPointer = block.PrevHash
		return nil
	})
	return block
}
