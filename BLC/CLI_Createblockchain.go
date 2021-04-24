package BLC

//初始化区块链
func (cli *CLI) createBlockchain(address string,nodeID string) {
	bc := CreateBlockChainWithGenesisBlock(address,nodeID)
	defer bc.DB.Close()

	//设置utxo重置操作
	utxoSet := &UTXOSet{bc}
	utxoSet.ResetUTXOSet()
}
