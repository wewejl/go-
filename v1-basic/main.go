package main

import (
	"fmt"
)



func main()  {
	//block:=NewBlock("hellowordd",nil)
	blockchain:=NewBlockchain()
	blockchain.addBlockchain("helloword")
	blockchain.addBlockchain("我想放屁了")
	for _,b:=range blockchain.blocks{
		fmt.Printf("++++++++++++++++++++\n")
		fmt.Printf("version:%s\n",b.version)
		fmt.Printf("prevHash:%x\n",b.prevHash)
		fmt.Printf("hash:%x\n",b.hash)
		//fmt.Printf("merkleRoot%x\n",b.merkleRoot)
		//fmt.Printf("timeStamp%d\n",b.timeStamp)
		//fmt.Printf("bits%d\n",b.bits)
		//fmt.Printf("nonce%d\n",b.nonce)
		fmt.Printf("data:%s\n",b.data)
	}
}
