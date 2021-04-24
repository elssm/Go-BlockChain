package BLC

//UTXO结构管理
type UTXO struct {
	//UTXO对应的交易哈希
	TxHash []byte
	//UTXO在其所属交易的输出列表中的索引
	Index int
	//Output本身
	Output *TxOutput
}
