package main

import (
	"BlockChain/bolt"
	"fmt"
	"log"
)

func main() {

	//打开数据库
	db, err := bolt.Open("test.db", 0600, nil)
	defer db.Close()
	if err != nil {
		log.Panic("打开数据库失败!")
		return
	}
	//找到抽屉
	db.Update(func(tx *bolt.Tx) error {
		//找到抽屉bucket 如果没有就创建
		bucket := tx.Bucket([]byte("b1"))
		if bucket == nil {
			//没有抽屉 需要创建抽屉
			bucket, err = tx.CreateBucket([]byte("b1"))
			if err != nil {
				log.Panic("bucket创建失败!")
			}
		}
		//写数据
		//bucket.Put([]byte("11111"), []byte("hello"))
		//bucket.Put([]byte("22222"), []byte("world"))

		return nil
	})

	//读数据
	db.View(func(tx *bolt.Tx) error {
		// 找到抽屉,没有的话直接退出
		bucket := tx.Bucket([]byte("b1"))
		if bucket == nil {
			log.Panic("bucket b1 不应为空,请检查!")
			return nil
		}
		//读取数据
		v1 := bucket.Get([]byte("11111"))
		v2 := bucket.Get([]byte("22222"))
		fmt.Printf("v1: %s\n", v1)
		fmt.Printf("v1: %s\n", v2)

		return nil
	})

}
