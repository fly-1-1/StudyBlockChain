package main

// BlockChain 引入区块链
type BlockChain struct {
	//定义一个区块链数组
	blocks []*Block
}

// NewBlockChain 定义一个区块链
func NewBlockChain() *BlockChain {
	//创建一个创世块,并作为第一个区块添加到区块链中
	genesisBlock := GenesisBlock()
	return &BlockChain{
		blocks: []*Block{genesisBlock},
	}
}

// GenesisBlock 创世块
func GenesisBlock() *Block {
	return NewBlock("First Block", []byte{})
}

// AddBlock 添加区块
func (bc *BlockChain) AddBlock(data string) {
	//获取最后一个区块
	lastBlock := bc.blocks[len(bc.blocks)-1]
	prevHash := lastBlock.Hash
	block := NewBlock(data, prevHash)
	bc.blocks = append(bc.blocks, block)
}
