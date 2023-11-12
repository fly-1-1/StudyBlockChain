package main

import (
	"fmt"
	"os"
	"strconv"
)

//接收命令行参数&控制区块链操作的文件

type CLI struct {
	bc *BlockChain
}

const Usage = `
	addBlock --data DATA  "add data to blockchain"
	printChain			  "print all blockchain data"
	printChainR			  "print all blockchain data revers"
	getBalance --address ADDRESS      "get balance of address"
	send FROM TO AMOUNT MINER DATA    "由FROM转AMOUNT给TO,由MINER挖矿,同时写入DATA"
`

const addCil = "addBlock"
const printCil = "printChain"
const rprintCil = "printChainR"
const getbalaceCil = "getBalance"
const sendCil = "send"

// Run 接收参数的动作
func (cil *CLI) Run() {
	//得到所有的参数
	args := os.Args
	if len(args) < 2 {
		fmt.Println("cmd:--------------", args[1])
		fmt.Printf("参数过少:%s", Usage)
		return
	}
	//分析命令
	cmd := args[1]
	switch cmd {
	case addCil:
		fmt.Print("添加区块\n")
		if len(args) == 4 && args[2] == "--data" {
			//a 获取数据
			data := args[3]
			//b 使用bc添加区块
			cil.AddBlock(data)
		} else {
			fmt.Println("添加区块参数有误")
			fmt.Printf(Usage)
		}
	case printCil:
		fmt.Printf("正向打印区块\n")
		cil.PrintChain()
	case rprintCil:
		fmt.Printf("反向打印区块\n")
		cil.PrintBlockChainRevers()
	case getbalaceCil:
		fmt.Printf("获取余额\n")
		if len(args) == 4 && args[2] == "--address" {
			address := args[3]
			cil.GetBalance(address)
		}
	case sendCil:
		fmt.Printf("转账开始...\n")
		if len(args) != 7 {
			fmt.Println("参数个数错误")
			fmt.Printf(Usage)
			return
		}
		from := args[2]
		to := args[3]
		amount, _ := strconv.ParseFloat(args[4], 64)
		miner := args[5]
		data := args[6]
		cil.Send(from, to, amount, miner, data)
	default:
		fmt.Print("无效的命令\n")
		fmt.Printf(Usage)
	}
	//执行相应动作
}
