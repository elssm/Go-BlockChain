package BLC

import (
	"fmt"
	"os"
)

//发起交易
func (cli *CLI) send(from,to,amount []string,nodeID string) {
	if !dbExist(nodeID) {
		fmt.Println("数据库不存在...")
		os.Exit(1)
	}
	//获取区块链对象
	blockchain := BlockchainObject(nodeID)
	defer blockchain.DB.Close()
	if len(from) != len(to) || len(from) != len(amount){
		fmt.Println("交易参数输入有误，请检查一致性....")
		os.Exit(1)
	}
	//发起交易，生成新的区块
	blockchain.MineNewBlock(from,to,amount)
	//调用utxo table的函数更新utxo table
	utxoSet := &UTXOSet{Blockchain: blockchain}
	utxoSet.update()

}