package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

//定义区块链
type Blockchain struct {
	//blocks []*Block
	db *bolt.DB
	listHash []byte
}

const genesisInfo = "天气转凉，注意御寒！"
const blockChainFilename ="blockchain.db"
const blockBucket="blockBucket"
const lastBlockHashKey ="lastBlockHashKey"

//创建blockchain，同时添加一个创世快
func NewBlockchain() *Blockchain {

	//连接数据库
	db,err:=bolt.Open(blockChainFilename,0600,nil)
	if err != nil {
		fmt.Println("NewBlockchain bolt.Open err:",err)
		return nil
	}
	var listHash []byte
	db.Update(func(tx *bolt.Tx) error {
		//打开桶
		b:=tx.Bucket([]byte(blockBucket))
		//如果桶没有就创建一个桶
		if b==nil {
			b,err=tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				fmt.Println("创建桶错误 err：",err)
				return err
			}
			//这里一定是第一次创建桶就要放进一个创世区块
			genesisBlock :=NewBlock(genesisInfo,nil)
			_=b.Put(genesisBlock.Hash,genesisBlock.Serialize()/*这个是把区块字节流*/)
			_=b.Put([]byte(lastBlockHashKey),genesisBlock.Hash)
			listHash=genesisBlock.Hash
		}else {
			lastBlockHashKey:=b.Get([]byte(lastBlockHashKey))
			listHash=lastBlockHashKey
		}
		return nil
	})

	//genesisBlock := NewBlock(genesisInfo, nil)

	return &Blockchain{
		db:       db,
		listHash: listHash,
	}
}

//添加区块的方法
func (bc *Blockchain) addBlockchain(data string) {
	//prevHash := bc.blocks[len(bc.blocks)-1].hash
	//获取区块连的最后一个区块的哈希
	prevHash:=bc.listHash
	//创建新的区块
	newblock := NewBlock(data, prevHash)

	bc.db.Update(func(tx *bolt.Tx) error {
		b:=tx.Bucket([]byte(blockBucket))
		if b==nil {
			log.Fatal("addBlockchain是 bucket不应该为空！！")
		}
		//fmt.Println("addBlockchain newblock.Hash:",newblock.Hash)
		err:=b.Put(newblock.Hash,newblock.Serialize())
		//数据库的listhHash也要更新
		_=b.Put([]byte(lastBlockHashKey),newblock.Hash)
		if err!=nil {
			fmt.Println("addBlockchain p.Put err:",err)
		}
		//要把listHash进行更新
		bc.listHash=newblock.Hash

		return nil
	})
}
