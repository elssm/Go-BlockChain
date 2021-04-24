package BLC

import (
	"github.com/boltdb/bolt"
	"log"
)

//区块链迭代器管理文件

//迭代器基本结构
type BlockChainIterator struct {
	DB *bolt.DB //迭代目标
	CurrentHash []byte //当前迭代目标的哈希
}

//创建迭代器对象
func (blc *BlockChain) Iterator() *BlockChainIterator {
	return &BlockChainIterator{blc.DB,blc.Tip}
}

//实现迭代函数next，获取到每一个区块
func (bcit *BlockChainIterator) Next() *Block {
	var block *Block

	err := bcit.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {
			currentBlockBytes := b.Get(bcit.CurrentHash)
			block = DeserializeBlock(currentBlockBytes)
			bcit.CurrentHash = block.PrevBlochHash
		}
		return nil
	})
	if err != nil {
		log.Panicf("iterator the db failed! %v\n",err)
	}
	return block
}