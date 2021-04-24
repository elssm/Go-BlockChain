package BLC
import (
	"bytes"
	"fmt"
	"math/big"
)

var b58Alphabet = []byte(""+"123456789"+"abcdefghijkmnopqrstuvwxyz"+"ABCDEFGHJKLMNPQRSTUVWXYZ")

//编码函数
func Base58Encode(input []byte) []byte {
	var result []byte
	//byte字节数组转换为big.int
	x := big.NewInt(0).SetBytes(input)
	//求余的基本长度
	base := big.NewInt(int64(len(b58Alphabet)))
	//求余数和商
	//判断条件，除掉的最终结果是否为0
	zero := big.NewInt(0)
	//设置余数，代表base58基数表的索引位置
	mod := &big.Int{}
	for x.Cmp(zero) != 0 {
		x.DivMod(x,base,mod)
		result = append(result,b58Alphabet[mod.Int64()])
	}
	Reverse(result)
	result = append([]byte{b58Alphabet[0]},result...)
	return result
}

func Reverse(data []byte)  {
	for i,j :=0,len(data)-1;i<j;i,j = i+1,j-1 {
		data[i],data[j] = data[j],data[i]
	}
}

//解码函数
func Base58Decode(input []byte) []byte {
	//fmt.Println(input)
	result := big.NewInt(0)
	zeroBytes := 1
	//去掉前缀
	data := input[zeroBytes:]
	//fmt.Println(data)
	for _,b := range data {
		//查找input中指定数字/字符在基数表中出现的索引(mod)
		charIndex := bytes.IndexByte(b58Alphabet,b)
		//余数*58
		result.Mul(result,big.NewInt(58))
		//乘积结果+mod(索引)
		result.Add(result,big.NewInt(int64(charIndex)))

	}
	//fmt.Println(result)
	//转换为byte字节数组
	decoded := result.Bytes()
	//fmt.Println(decoded)
	return decoded
}

func main()  {
	result := Base58Encode([]byte("elssm"))
	fmt.Printf("result : %s\n",result)
	decodeResult := Base58Decode([]byte("1crFth1D"))
	fmt.Printf("decode Result : %s\n",decodeResult)
}