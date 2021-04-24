package BLC

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

//区块基本结构与功能管理文件


//实现一个最基本的区块结构
type Block struct {
	TimeStamp int64 //区块时间戳
	Hash []byte //当前区块哈希
	PrevBlochHash []byte //前区块哈希
	Heigth int64 //区块高度
	//Data []byte //交易数据
	Txs []*Transaction //交易数据
	Nonce int64 //在运行pow时生成的哈希变化值，也代表pow运行时动态修改的数据
}

//新建区块
func NewBlock(height int64,prevBlockHash []byte,txs []*Transaction) *Block {
	var block Block

	block = Block{
		TimeStamp: time.Now().Unix(),
		Hash: nil,
		PrevBlochHash: prevBlockHash,
		Heigth: height,
		Txs: txs,
	}
	//生成哈希
	//block.SetHash()
	//替换setHash通过POW生成新的哈希
	pow := NewProofOfWork(&block)
	hash,nonce := pow.Run()
	block.Hash = hash
	block.Nonce = int64(nonce)
	return &block
}

//技术区块哈希
//func (b *Block) SetHash()  {
//	//调用sha256实现哈希生成
//	//实现int->hash
//	timeStampBytes := IntToHex(b.TimeStamp)
//	heigthBytes := IntToHex(b.Heigth)
//	blockBytes := bytes.Join([][]byte{
//	heigthBytes,
//	timeStampBytes,
//	b.PrevBlochHash,
//	b.Data,
//	},[]byte{})
//	hash := sha256.Sum256(blockBytes)
//	b.Hash = hash[:]
//}

//生成创世区块
func CreateGenesisBlock(txs []*Transaction) *Block {
	return NewBlock(1,nil,txs)
}

//区块结构序列化
func (block *Block)Serialize() []byte {
	var buffer bytes.Buffer
	//新建编码对象
	encoder := gob.NewEncoder(&buffer)
	if err := encoder.Encode(block);nil != err {
		log.Panicf("serialize the block to []byte failed %v\n",err)
	}
	return buffer.Bytes()
}


//区块结构反序列化
func DeserializeBlock(blockBytes []byte) *Block {
	var block Block
	//新建decode对象
	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	if err := decoder.Decode(&block);nil != err {
		log.Panicf("deserialize the block to []byte failed %v\n",err)
	}
	return &block
}

//把指定区块中的所有交易结构都序列化
func (block *Block) HashTransaction() []byte {
	var txHashes [][]byte
	for _,tx := range block.Txs {
		txHashes = append(txHashes,tx.Txhash)
	}
	//将交易数据存入Merkle树中，然后生成Merkle根节点
	mtree := NewMerkleTree(txHashes)
	//txHash := sha256.Sum256(bytes.Join(txHashes,[]byte{}))
	//return txHash[:]
	return mtree.RootNode.Data
}