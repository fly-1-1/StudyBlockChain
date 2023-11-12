package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
	"log"
)

//钱包结构 每一个钱包保存了公钥私钥对

type Wallet struct {
	Private *ecdsa.PrivateKey
	//PubKey不存储原始公钥 而是存储原始的x和y拼接的字符串 在校验端重新拆分
	PubKey []byte
}

// NewWallet 创建钱包
func NewWallet() *Wallet {
	Curve := elliptic.P224()
	privateKey, err := ecdsa.GenerateKey(Curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	//生成公钥
	pubKeyOrig := privateKey.PublicKey
	pubKey := append(pubKeyOrig.X.Bytes(), pubKeyOrig.Y.Bytes()...)
	return &Wallet{Private: privateKey, PubKey: pubKey}
}

// NewAddress 生成地址
func (w *Wallet) NewAddress() string {
	pubKey := w.PubKey
	hash := sha256.Sum256(pubKey)
	//编码器
	rip160hasher := ripemd160.New()
	_, err := rip160hasher.Write(hash[:])
	if err != nil {
		log.Panic(err)
	}
	rip160hashValue := rip160hasher.Sum(nil)
	version := byte(00)
	payload := append([]byte{version}, rip160hashValue...)
	//checksum
	hash1 := sha256.Sum256(payload)
	hash2 := sha256.Sum256(hash1[:])
	//字节校验码
	checkCode := hash2[:4]
	//25字节数据
	payload = append(payload, checkCode...)
	//btcd BTC全节点源码
	address := base58.Encode(payload)
	return address
}
