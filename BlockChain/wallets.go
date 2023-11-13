package main

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"io/ioutil"
	"log"
	"os"
)

const walletFile = "wallet.dat"

// Wallets 保存所有的wallet 以及它的地址
type Wallets struct {
	//map[address]wallet
	WalletsMap map[string]*Wallet
}

// NewWallets 创建方法
func NewWallets() *Wallets {
	var ws Wallets
	ws.WalletsMap = make(map[string]*Wallet)
	ws.loadFile()
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
	ioutil.WriteFile(walletFile, buffer.Bytes(), 0600)
}

//把所有的钱包读出来
func (ws *Wallets) loadFile() {
	//在读取之前要确认文件是否存在
	_, err := os.Stat(walletFile)
	if os.IsNotExist(err) {
		//ws.WalletsMap = make(map[string]*Wallet)
		return
	}

	content, err := ioutil.ReadFile(walletFile)
	if err != nil {
		log.Panic(err)
	}

	decoder := gob.NewDecoder(bytes.NewReader(content))

	var ws1 Wallets

	err = decoder.Decode(&ws1)
	if err != nil {
		log.Panic(err)
	}
	ws.WalletsMap = ws1.WalletsMap
}

func (ws *Wallets) ListAllAddresses() []string {

	var addresses []string
	for address := range ws.WalletsMap {
		addresses = append(addresses, address)
	}
	return addresses
}
