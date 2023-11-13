package main

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"io/ioutil"
	"log"
)

// Wallets 保存所有的wallet 以及它的地址
type Wallets struct {
	//map[address]wallet
	WalletsMap map[string]*Wallet
}

// NewWallets 创建方法
func NewWallets() *Wallets {
	var ws Wallets
	//ws := loadFile()
	ws.WalletsMap = make(map[string]*Wallet)
	return &ws
}

func (ws *Wallets) CreateWallet() string {
	wallet := NewWallet()
	address := wallet.NewAddress()
	ws.WalletsMap[address] = wallet
	ws.saveToFile()
	return address
}

func init() {
	gob.RegisterName("elliptic.P256", elliptic.P256())
}

// 保存方法,把新建的wallet添加进去
func (ws *Wallets) saveToFile() {

	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(ws)
	if err != nil {
		log.Panic(err)
	}
	ioutil.WriteFile("wallet.data", buffer.Bytes(), 0600)
}

//把所有的钱包读出来
