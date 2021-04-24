package BLC

import (
	"flag"
	"fmt"
	"log"
	"os"
)

//对blockchain进行命令行管理

//client对象
type CLI struct {
	BC *BlockChain //区块链对象
}

//用法展示
func PrintUsage()  {
	fmt.Println("Usage:")
	fmt.Printf("\tcreatewallet --创建钱包 \n")
	fmt.Printf("\taccounts --获取钱包地址 \n")
	//初始化区块链
	fmt.Printf("\tcreateblockchain -address address -- 创建区块链\n")
	//添加区块
	//fmt.Printf("\taddblock -data DATA -- 添加区块\n")
	//打印完整的区块信息
	fmt.Printf("\tprintchain -- 输出区块信息\n")
	//通过命令后转账
	fmt.Printf("\tsend -from FROM -to TO -amount AMOUNT -- 发起转账\n")
	fmt.Printf("\t 转账参数说明\n")
	fmt.Printf("\t\t-from FROM -- 转账源地址\n")
	fmt.Printf("\t\t-to TO -- 转账目标地址\n")
	fmt.Printf("\t\t-amount AMOUNT -- 转账金额\n")
	//查询余额
	fmt.Printf("\tgetbalance -address FROM --查询指定地址的余额\n")
	fmt.Println("\t查询余额参数说明")
	fmt.Printf("\t\t -address --查询余额的地址\n")
	fmt.Printf("\tutxo -method METHOD -- 测试UTXO Table功能中指定的方法\n")
	fmt.Printf("\t\tMETHOD --方法名\n")
	fmt.Printf("\t\t\treset --重置UTXOtable\n")
	fmt.Printf("\t\t\tbalance -查找所有UTXO\n")
	fmt.Printf("\tset_id -port PORT -- 设置节点号\n")
	fmt.Printf("\t\tport -- 访问的节点号")
}

////查询余额
//func (cli *CLI) getBalance(from string) {
//	//查找该地址UTXO
//	//获取区块链对象
//	blockchain := BlockchainObject()
//	defer blockchain.DB.Close()
//	amount := blockchain.getBalance(from)
//	fmt.Printf("\t地址 [%s] 的余额：[%d]\n",from,amount)
//}


////发起交易
//func (cli *CLI) send(from,to,amount []string) {
//	if !dbExist() {
//		fmt.Println("数据库不存在...")
//		os.Exit(1)
//	}
//	//获取区块链对象
//	blockchain := BlockchainObject()
//	defer blockchain.DB.Close()
//	blockchain.MineNewBlock(from,to,amount)
//
//}

////初始化区块链
//func (cli *CLI) createBlockchain(address string) {
//	CreateBlockChainWithGenesisBlock(address)
//}
//添加区块
func (cli *CLI) addBlock(txs []*Transaction) {
	//判断数据库是否存在
	//if !dbExist() {
	//	fmt.Println("数据库不存在...")
	//	os.Exit(1)
	//}
	////获取到blockchain的对象实例
	//blockchain := BlockchainObject()
	//blockchain.AddBlock(txs)
	//cli.BC.AddBlock([]byte(data))
}
////打印完整区块链信息
//func (cli *CLI) printchain() {
//	//判断数据库是否存在
//	if !dbExist() {
//		fmt.Println("数据库不存在...")
//		os.Exit(1)
//	}
//	//cli.BC.PrintChain()
//	blockchain := BlockchainObject()
//	blockchain.PrintChain()
//}

////参数数量检测函数
//func IsValidArgs()  {
//	if len(os.Args) < 2 {
//		PrintUsage()
//		os.Exit(1)
//	}
//}


//命令行运行函数
func (cli *CLI) Run() {
	nodeId := GetEnvNodeId()
	IsValidArgs()
	//创建钱包
	createWalletCmd := flag.NewFlagSet("createwallet",flag.ExitOnError)
	//获取地址列表
	getAccountsCmd := flag.NewFlagSet("accounts",flag.ExitOnError)
	//utxo测试命令
	UTXOTestCmd := flag.NewFlagSet("utxo",flag.ExitOnError)
	//节点号设置命令
	setNodeIdCmd := flag.NewFlagSet("set_id",flag.ExitOnError)
	//添加区块
	addBlockCmd := flag.NewFlagSet("addblock",flag.ExitOnError)
	//输出区块链完整信息
	printChainCmd := flag.NewFlagSet("printchain",flag.ExitOnError)
	//创建区块链
	createBLCWithGenesisBlockCmd := flag.NewFlagSet("createblockchain",flag.ExitOnError)

	//发起交易
	sendCmd := flag.NewFlagSet("send",flag.ExitOnError)

	//查询余额命令
	getBalanceCmd := flag.NewFlagSet("getbalance",flag.ExitOnError)

	//数据参数
	flagAddBlockArg := addBlockCmd.String("data","sent 100 btc to player","添加区块数据")

	//创建区块链时指定的矿工地址（接受奖励）
	flagCreateBlockchainArg := createBLCWithGenesisBlockCmd.String("address","elssm","指定接受系统奖励的矿工地址")

	//发起交易参数
	flagSendFromArg := sendCmd.String("from","","转账源地址")
	flagSendToArg := sendCmd.String("to","","转账目标地址")
	flagSendAmountArg := sendCmd.String("amount","","转账金额")

	//查询余额命令行参数
	flagGetBalanceArg := getBalanceCmd.String("address","","要查询的地址")
	//UTXO测试命令行参数
	flagUTXOArg := UTXOTestCmd.String("method","","UTXO Table相关操作\n")
	//端口号参数
	flagPortArg := setNodeIdCmd.String("port","","设置节点 ID")



	//判断命令
	switch os.Args[1] {
	case "set_id":
		if err := setNodeIdCmd.Parse(os.Args[2:]); nil!=err {
			log.Panicf("parse cmd  set node id failed! %v\n",err)
		}

	case "utxo":
		if err := UTXOTestCmd.Parse(os.Args[2:]); nil!=err {
			log.Panicf("parse cmd operate utxo table failed! %v\n",err)
		}
	case "accounts":
		if err := getAccountsCmd.Parse(os.Args[2:]); nil!=err {
			log.Panicf("parse cmd getaccounts failed! %v\n",err)
		}
	case "createwallet":
		if err := createWalletCmd.Parse(os.Args[2:]); nil!=err {
			log.Panicf("parse cmd create wallet failed! %v\n",err)
		}
	case "getbalance":
		if err := getBalanceCmd.Parse(os.Args[2:]); nil!=err {
			log.Panicf("parse cmd get balance failed! %v\n",err)
		}

	case "send":
		if err := sendCmd.Parse(os.Args[2:]); nil!=err {
			log.Panicf("parse send failed! %v\n",err)
		}

	case "addblock":
		if err := addBlockCmd.Parse(os.Args[2:]); nil!=err {
			log.Panicf("parse addBlockCmd failed! %v\n",err)
		}

	case "printchain":
		if err := printChainCmd.Parse(os.Args[2:]); nil!=err {
			log.Panicf("parse printChainCmd failed! %v\n",err)
		}

	case "createblockchain":
		if err := createBLCWithGenesisBlockCmd.Parse(os.Args[2:]); nil!=err {
			log.Panicf("parse createblockchain failed! %v\n",err)
		}

	default:
		PrintUsage()
		os.Exit(1)
	}

	//节点ID设置
	if setNodeIdCmd.Parsed() {
		if *flagPortArg == "" {
			fmt.Println("请输入端口号...")
			os.Exit(1)
		}
		cli.SetNodeId(*flagPortArg)
	}


	//utxo table操作
	if UTXOTestCmd.Parsed() {
		switch *flagUTXOArg {
		case "balance":
			cli.TestFindUTXOMap()
		case "reset":
			cli.TestResetUTXO(nodeId)
		default:



		}
	}

	//获取地址列表
	if getAccountsCmd.Parsed() {
		cli.GetAccounts()
	}

	//创建钱包
	if createWalletCmd.Parsed() {
		cli.CreateWallets()
	}

	//查询余额
	if getBalanceCmd.Parsed() {
		if *flagGetBalanceArg == "" {
			fmt.Println("请输入查询地址...")
			os.Exit(1)
		}
		cli.getBalance(*flagGetBalanceArg,nodeId)
	}

	//发起转账
	if sendCmd.Parsed() {
		if *flagSendFromArg == "" {
			fmt.Println("源地址不能为空")
			PrintUsage()
			os.Exit(1)
		}
		if *flagSendToArg == "" {
			fmt.Println("目标地址不能为空")
			PrintUsage()
			os.Exit(1)
		}
		if *flagSendAmountArg == "" {
			fmt.Println("转账金额不能为空")
			PrintUsage()
			os.Exit(1)
		}
		fmt.Printf("\tFROM:[%s]\n",JSONToSlice(*flagSendFromArg))
		fmt.Printf("\tTO:[%s]\n",JSONToSlice(*flagSendToArg))
		fmt.Printf("\tAMOUNT:[%s]\n",JSONToSlice(*flagSendAmountArg))

		cli.send(JSONToSlice(*flagSendFromArg),JSONToSlice(*flagSendToArg),JSONToSlice(*flagSendAmountArg),nodeId)
	}

	//添加区块命令
	if addBlockCmd.Parsed() {
		if *flagAddBlockArg == "" {
			PrintUsage()
			os.Exit(1)
		}
		cli.addBlock([]*Transaction{})
	}
	//输出区块链信息
	fmt.Println("nodeID : ",nodeId)
	if printChainCmd.Parsed() {
		cli.printchain(nodeId)
	}
	//创建区块命令
	if createBLCWithGenesisBlockCmd.Parsed() {
		if *flagCreateBlockchainArg == "" {
			PrintUsage()
			os.Exit(1)
		}
		cli.createBlockchain(*flagCreateBlockchainArg,nodeId)
	}
}
