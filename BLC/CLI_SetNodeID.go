package BLC

import (
	"fmt"
	"log"
	"os"
)

//设置端口号（环境变量）
func (cli *CLI) SetNodeId(nodeId string) {
	fmt.Println("SetNodeId:",nodeId)
	if nodeId == "" {
		fmt.Println("please set the node id...")
		os.Exit(1)
	}
	err := os.Setenv("NODE_ID",nodeId)
	if err != nil {
		log.Fatalf("set env failed! %v\n",err)
	}
}
