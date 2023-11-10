package main

import (
	"fmt"
	"os"
)

//接收命令行参数&控制区块链操作的文件

type CLI struct {
	bc *BlockChain
}

const Usage = `
	addBlock --data DATA  "add data to blockchain"
	printChain			  "print all blockchain data"
`
const addCil = "addBlock"
const printCil = "printChain"

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
		fmt.Print("打印区块\n")
		cil.PrintBlockChain()
	default:
		fmt.Print("无效的命令\n")
		fmt.Printf(Usage)
	}
	//执行相应动作
}
