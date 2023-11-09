package main

import "fmt"

func main() {
	//1.每21万个块奖励减半
	//2.最初奖励50BTC
	total := 0.0
	blockInterval := 21.0 //万
	currentReward := 50.0
	for currentReward > 0 {
		amount1 := blockInterval * currentReward
		currentReward *= 0.5
		total += amount1
	}
	fmt.Println("BTC TOTAL:", total, "万")
	//BTC TOTAL: 2100 万
}
