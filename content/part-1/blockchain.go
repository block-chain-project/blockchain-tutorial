package main

// 是一个 Block 指针数组
type BlockChain struct {
	blocks []*Block
}

// 创建一个有创世块的链
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

// 向链中加入一个新块
//
// data 在实际中就是交易
func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func (bc *BlockChain) List() []*Block {
	return bc.blocks
}
