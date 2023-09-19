package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// 由区块头和交易两部分构成
//
// Timestamp, PrevBlockHash, Hash 属于区块头（block header）
type Block struct {
	Head        // 区块头
	Data []byte // 区块实际存储的信息，比特币中也就是交易
}

type Head struct {
	Timestamp     int64  // 当前时间戳，也就是区块创建的时间
	PrevBlockHash []byte // 前一个块的哈希
	Hash          []byte // 当前块的哈希
}

// 用于生成新块，参数需要 Data 与 PrevBlockHash
//
// 当前块的哈希会基于 Data 和 PrevBlockHash 计算得到
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Head: Head{
			Timestamp:     time.Now().Unix(),
			PrevBlockHash: prevBlockHash,
			Hash:          []byte{},
		},
		Data: []byte(data),
	}

	block.SetHash()

	return block
}

// 设置当前块哈希
//
// Hash = sha256(PrevBlockHash + Data + Timestamp)
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	// 当前区块hash = sha256(上一个数据hash + 数据 + 时间截)
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// 生成创世块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
