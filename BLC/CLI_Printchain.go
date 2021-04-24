package BLC

import (
	"fmt"
	"os"
)

//打印完整区块链信息
func (cli *CLI) printchain(nodeID string) {
	//判断数据库是否存在
	if !dbExist(nodeID) {
		fmt.Println("数据库不存在...")
		os.Exit(1)
	}
	//cli.BC.PrintChain()
	blockchain := BlockchainObject(nodeID)
	blockchain.PrintChain()
}
