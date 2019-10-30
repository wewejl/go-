package main

import "os"

func main()  {
	//block:=NewBlock("hellowordd",nil)
	//获取命令行
	data:=os.Args
	blockchain:=NewBlockchain()
	cli:=NewCli(blockchain)
	cli.Run(data)
}

