package BLC

import "bytes"

//交易输入管理
type TxInput struct {
	//交易哈希
	TxHash []byte
	//引用的上一笔交易的输出索引号
	Vout int
	//用户名
	//ScriptSig string
	//数字签名
	Signature []byte
	//公钥
	PublicKey []byte
}


//验证引用的地址是否匹配
//func (txInput *TxInput) CheckPubKeyWithAddress(address string) bool {
//	return address == txInput.ScriptSig
//}

//传递哈希160进行判断
func (in *TxInput) UnlockRipemd160Hash(ripemd160Hash []byte) bool {
	//获取input的ripemd160哈希值
	inputRipemd160Hash := Ripemd160Hash(in.PublicKey)
	return bytes.Compare(inputRipemd160Hash,ripemd160Hash) == 0
}