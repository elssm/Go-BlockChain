package BLC

import "crypto/sha256"

//Merkle树实现管理
type MerkleTree struct {
	//根节点
	RootNode *MerkleNode
}

//merkle节点结构
type MerkleNode struct {
	//左子节点
	Left *MerkleNode
	//右子节点
	Right *MerkleNode
	//数据
	Data []byte
}

//创建Merkle树
//txHashes:区块中的交易哈希列表
//Merkle根节点之外的其他层次的节点数量必须是偶数个，如果是奇数个，则将最后一个节点复制一份
func NewMerkleTree(txHashes [][]byte) *MerkleTree {
	 //节点列表
	var nodes []MerkleNode
	//判断交易数据条数，如果是奇数，则拷贝最后一份
	if len(txHashes) % 2 != 0 {
		txHashes = append(txHashes,txHashes[len(txHashes)-1])
	}

	//遍历所有交易数据，通过哈希生成叶子节点
	for _,data := range txHashes {
		node := MakeMerkleNode(nil,nil,data)
		nodes = append(nodes,*node)
	}

	//通过叶子节点创建父节点
	for i := 0 ; i < len(txHashes) / 2; i++ {
		var parentNodes []MerkleNode //父节点列表
		for j := 0; j < len(nodes); j+=2 {
			node := MakeMerkleNode(&nodes[j],&nodes[j+1],nil)
			parentNodes = append(parentNodes,*node)
		}

		if len(parentNodes) % 2 != 0 {
			parentNodes = append(parentNodes,parentNodes[len(parentNodes) - 1])
		}
		nodes = parentNodes
	}
	//最终，nodes中只保存了根节点的哈希值
	mtree := MerkleTree{&nodes[0]}
	return &mtree
}

//创建Merkle节点
func MakeMerkleNode(left,right *MerkleNode,data []byte) *MerkleNode {
	node := &MerkleNode{}
	//判断叶子节点
	if left == nil && right == nil {
		hash := sha256.Sum256(data)
		node.Data = hash[:]
	} else {
		//非叶子节点
		prveHashes := append(left.Data,right.Data...)
		hash := sha256.Sum256(prveHashes)
		node.Data = hash[:]
	}

	node.Left = left
	node.Right = right
	return node
}