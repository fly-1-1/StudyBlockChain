package main

import (
	"BlockChain/bolt"
	"fmt"
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
func NewBlockChain(address string) *BlockChain {

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
			genesisBlock := GenesisBlock(address)
<<<<<<< Updated upstream
			fmt.Printf("genesisBlock: %s\n", genesisBlock)
=======
			//fmt.Printf("genesisBlock: %s\n", genesisBlock)
>>>>>>> Stashed changes
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
func GenesisBlock(address string) *Block {
	coinbase := NewCoinbaseTX(address, "Genesis Block")
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

// AddBlock 添加区块
func (bc *BlockChain) AddBlock(txs []*Transaction) {
	//获取前区块的哈希
	db := bc.db         //区块链数据库
	lastHash := bc.tail //最后一个区块的哈希

	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("bucket 不应该为空,请检查!")
		}
		//创建新的区块
		block := NewBlock(txs, lastHash)
		//完成数据的添加
		//添加到区块db中
		bucket.Put(block.Hash, block.Serialize())
		bucket.Put([]byte("LastHashKey"), block.Hash)
		//更新内存的区块链(最后的尾巴需要更新)
		bc.tail = block.Hash
		return nil
	})

}

// FindUTXOs 找到指定地址的utxo
func (bc *BlockChain) FindUTXOs(address string) []TXOutput {
	var UTXO []TXOutput
	//map[交易id][]int64
	//map保存消费过的output,key是output所在id value是这个交易中索引的数组
	spentOutputs := make(map[string][]int64)
	//TODO
	//遍历区块
	//遍历交易
	//迭代器
	it := bc.NewIterator()
	for {
		block := it.Next()
		for _, tx := range block.Transactions {
			fmt.Printf("当前的交易Id: %x\n", tx.TXID)
		OUTPUT:
			//遍历output 找到和自己相关的utxo(在添加output之前是否已经消耗过)
			for i, output := range tx.TXOutputs {
				fmt.Printf("当前的Index: %d\n", i)
				//过滤将所有消耗过得outputs和当前的所即将添加的outputs对比一下 如果相同则跳过 否则添加
				//如果当前的交易ID存在于已经标识的map 说明这个交易中 有消耗过的output
				if spentOutputs[string(tx.TXID)] != nil {
					for _, j := range spentOutputs[string(tx.TXID)] {
						if int64(i) == j {
							//说明当前准备的output已经用过了,不用再添加
							continue OUTPUT
						}
					}
				}
				//这个output和目标地址相同 满足条件 加到返回utxo数组中
				if output.PukKeyHash == address {
					UTXO = append(UTXO, output)
				}
			}
			//如果当前是coinbase交易直接跳过
			if !tx.IsCoinBase() {
				//遍历input  找到自己花费过的utxo集合(把自己消费过得标识出来)
				for _, input := range tx.TXInputs {
					//判断当下input和目标是否一致 如果相同说明是目标消耗过得output,加进去
					if input.Sig == address {
						//spentOutputs := make(map[string][]int64)
						indexArray := spentOutputs[string(input.TXid)]
						indexArray = append(indexArray, input.Index)
						//map[2222]=[]int64{0}
						//map[3333]=[]int64{0,1}
					}
				}
			} else {
				fmt.Println("这是coinbase不做input遍历")
			}
		}
		if len(block.PrevHash) == 0 {
<<<<<<< Updated upstream
			break
			fmt.Printf("区块打印完成!")
=======
			fmt.Printf("区块打印完成!\n")
			break
>>>>>>> Stashed changes
		}
	}
	return UTXO
}

func (bc *BlockChain) FindNeedUTXOs(from string, amount float64) (map[string][]uint64, float64) {
	//找到合理的utxos集合
<<<<<<< Updated upstream
	var utxos map[string][]uint64
	//找到的utxos里面包含的钱的总数
	var calc float64
=======
	utxos := make(map[string][]uint64)
	//标识已经消耗过的utxo
	spentOutputs := make(map[string][]int64)
	//找到的utxos里面包含的钱的总数
	var calc float64

	//----------------------------------
	it := bc.NewIterator()
	for {
		block := it.Next()
		for _, tx := range block.Transactions {
			fmt.Printf("当前的交易Id: %x\n", tx.TXID)
		OUTPUT:
			//遍历output 找到和自己相关的utxo(在添加output之前是否已经消耗过)
			for i, output := range tx.TXOutputs {
				fmt.Printf("当前的Index: %d\n", i)
				//过滤将所有消耗过得outputs和当前的所即将添加的outputs对比一下 如果相同则跳过 否则添加
				//如果当前的交易ID存在于已经标识的map 说明这个交易中 有消耗过的output
				if spentOutputs[string(tx.TXID)] != nil {
					for _, j := range spentOutputs[string(tx.TXID)] {
						if int64(i) == j {
							//说明当前准备的output已经用过了,不用再添加
							continue OUTPUT
						}
					}
				}
				//这个output和目标地址相同 满足条件 加到返回utxo数组中
				if output.PukKeyHash == from {
					//UTXO = append(UTXO, output)
					//找到需要的最少的UTXO //TODO
					//把UTXO加进来
					//统计一下当前UTXO的总额
					//比较一下是否满足转账需求 满足直接返回utxos calc 不满足继续统计
					if calc < amount {
						//array := utxos[string(tx.TXID)]
						//array = append(array, uint64(i))
						utxos[string(tx.TXID)] = append(utxos[string(tx.TXID)], uint64(i))
						calc += output.Value

						//满足条件
						if calc >= amount {
							fmt.Printf("找到了满足的金额:%f\n", calc)
							return utxos, calc
						}
					} else {
						fmt.Printf("不满足转账金额,当前总额%f,目标金额%f", calc, amount)
					}

				}
			}
			//如果当前是coinbase交易直接跳过
			if !tx.IsCoinBase() {
				//遍历input  找到自己花费过的utxo集合(把自己消费过得标识出来)
				for _, input := range tx.TXInputs {
					//判断当下input和目标是否一致 如果相同说明是目标消耗过得output,加进去
					if input.Sig == from {
						//spentOutputs := make(map[string][]int64)
						indexArray := spentOutputs[string(input.TXid)]
						indexArray = append(indexArray, input.Index)
						//map[2222]=[]int64{0}
						//map[3333]=[]int64{0,1}
					}
				}
			} else {
				fmt.Println("这是coinbase不做input遍历")
			}
		}
		if len(block.PrevHash) == 0 {
			break
			fmt.Printf("区块打印完成!")
		}
	}
	//--------------------------------

>>>>>>> Stashed changes
	return utxos, calc
}
