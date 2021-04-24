package BLC

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"os"
)


//参数数量检测函数
func IsValidArgs()  {
	if len(os.Args) < 2 {
		PrintUsage()
		os.Exit(1)
	}
}


//实现int64转[]byte
func IntToHex(data int64) []byte {
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer,binary.BigEndian,data)
	if err != nil {
		log.Panicf("int transact to []byte failed! %v\n",err)
	}
	return buffer.Bytes()
}


//标准json转切片
//格式为 send -from "[\"1j4DnszjJbaDxjjpn3kiDZKPthzzWbBwPq\"]" -to "[\"16hQSbGy4SyGkwD96XaYZs6wVzH5nXt9zU\"]" -amount "[\"2\"]"
//格式为 send -from "[\"elssm\",\"warry\"]" -to "[\"warry\",\"elssm\"]" -amount "[\"5\",\"2\"]"
func JSONToSlice(jsonString string) []string {
	var strSlice []string
	//json
	if err := json.Unmarshal([]byte(jsonString),&strSlice); nil!=err {
		log.Panicf("json to []string failed %v\n",err)
	}
	return strSlice
}

//string转hash160
func StringToHash160(address string) []byte {
	pubKeyHash := Base58Decode([]byte(address))
	hash160 := pubKeyHash[:len(pubKeyHash)-addressCheckSumLen]
	return hash160
}

//获取节点ID
func GetEnvNodeId() string {
	nodeID := os.Getenv("NODE_ID")
	if nodeID == "" {
		fmt.Println("NODE_ID is not set...")
		os.Exit(1)
	}
	return nodeID
}