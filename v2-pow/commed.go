package main

import "fmt"

func (c *Cli)addBlockchain(data string)  {
	c.Bc.addBlockchain(data)
}
func (c *Cli)Print()  {
		//从bold数据读到最后一个哈希
		//db:=blockchain.db
		//defer db.Close()
		//创建一个迭代器
		iterator:=NewIterator(c.Bc)
		for  {
			b:=iterator.next()
						//fmt.Println("***********************b",b.PrevHash)
						fmt.Printf("++++++++++++++++++++\n")
						fmt.Printf("version:%s\n",b.Version)
						fmt.Printf("prevHash:%x\n",b.PrevHash)
						fmt.Printf("hash:%x\n",b.Hash)
						//fmt.Printf("merkleRoot%x\n",b.merkleRoot)
						//fmt.Printf("timeStamp%d\n",b.timeStamp)
						//fmt.Printf("bits%d\n",b.bits)
						//fmt.Printf("nonce%d\n",b.nonce)
						fmt.Printf("data:%s\n",b.Data)
						if b.PrevHash==nil{
							//fmt.Println("跳出来了")
							break
						}
			//iterator.currhash=b.PrevHash
		}
}