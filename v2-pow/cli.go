package main

import "C"
import "fmt"

type Cli struct {
	Bc *Blockchain
}
//这个我们定的的规则
const agstestgz =`
./agstest addBlockchain <data> //添加数据
./agstest println 	//打印数据
`
func NewCli(bc *Blockchain) *Cli {
	return &Cli{Bc:bc}
}

func (c *Cli)Run(data []string)  {
	if len(data) < 2 {//小于2就一定就错误的
		fmt.Println(agstestgz)
		return
	}
	//拿到数据2 命令
	switch data[1] {
	case "println":
		if len(data)!=2{
			fmt.Println("参数错误")
			fmt.Println(agstestgz)
			return
		}
		fmt.Println("我正在打印 ")
		c.Print()
	case "addBlockchain":
		if len(data)!=3 {
			fmt.Println("参数错误")
			fmt.Println(agstestgz)
			return
		}
		fmt.Println("我正在挖矿 数据是",data[2])
		c.addBlockchain(data[2])
	default:
		fmt.Println(agstestgz)
		return
	}
}


//	blockchain.addBlockchain("helloword")
//	blockchain.addBlockchain("我想放屁了")
