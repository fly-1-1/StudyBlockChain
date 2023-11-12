package main

import (
	"BlockChain/bolt"
	"bytes"
	"fmt"
)

func (cli *CLI) PrintBlockChainRevers() {

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
		fmt.Printf("区块数据:%s\n", block.Transactions[0].TXInputs[0].Sig)
		if len(block.PrevHash) == 0 {
			break
		}
	}
}

func (cli *CLI) PrintChain() {

	bc := cli.bc
	blockHeight := 0
	bc.db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("blockBucket"))

		//从第一个key-> value 进行遍历，到最后一个固定的key时直接返回
		b.ForEach(func(k, v []byte) error {
			if bytes.Equal(k, []byte("LastHashKey")) {
				return nil
			}

			block := DeSerialize(v)
			//fmt.Printf("key=%x, value=%s\n", k, v)
			fmt.Printf("=============== 区块高度: %d ==============\n", blockHeight)
			blockHeight++
			fmt.Printf("版本号: %d\n", block.Version)
			fmt.Printf("前区块哈希值: %x\n", block.PrevHash)
			fmt.Printf("梅克尔根: %x\n", block.MerkelRoot)
			fmt.Printf("时间戳: %d\n", block.TimeStamp)
			fmt.Printf("难度值(随便写的）: %d\n", block.Difficulty)
			fmt.Printf("随机数 : %d\n", block.Nonce)
			fmt.Printf("当前区块哈希值: %x\n", block.Hash)
			fmt.Printf("区块数据 :%s\n", block.Transactions[0].TXInputs[0].Sig)
			return nil
		})
		return nil
	})
}

func (cli *CLI) AddBlock(data string) {
	//cli.bc.AddBlock(data) TODO
	fmt.Println("添加区块成功")
}

func (cli *CLI) GetBalance(address string) {
	utxos := cli.bc.FindUTXOs(address)

	total := 0.0
	for _, utxo := range utxos {
		total += utxo.Value
	}
	fmt.Printf("\"%s\"的余额为%f\n", address, total)
}

func (cli *CLI) Send(from, to string, amount float64, miner, data string) {
	fmt.Printf("from: %s\n", from)
	fmt.Printf("to: %s\n", to)
	fmt.Printf("amount: %f\n", amount)
	fmt.Printf("miner: %s\n", miner)
	fmt.Printf("data: %s\n", data)
	//具体的逻辑 TODO
	//创建挖矿交易
	coinbase := NewCoinbaseTX(miner, data)
	//创建一个普通交易
	tx := NewTransaction(from, to, amount, cli.bc)
	if tx == nil {
		return
	}
	cli.bc.AddBlock([]*Transaction{coinbase, tx})
	fmt.Printf("转账成功!")
}
