package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"time"
)

//定义区块结构

//区块头
type Block struct {
	//版本号
	version string
	//前区块哈希
	prevHash []byte

	//当前区块的哈希
	//在比特币定义的区块中，是没有这个当前的区块哈希这个字段的
	//为了方便
	hash []byte

	//merkle根，根据当前区块的交易数据计算出来的
	merkleRoot []byte

	//时间戳
	timeStamp int64

	//难度值,系统提供的
	bits int64

	//随机数
	nonce int64

	//区块体，交易数据
	data []byte
}

func NewBlock(data string, preHash []byte) *Block {

	newblock := &Block{
		version:  "0",
		prevHash: preHash,

		merkleRoot: nil,
		timeStamp:  time.Now().Unix(),
		bits:       0,
		nonce:      0,
		data:       []byte(data),
	}
	
	newblock.setHash()

	return newblock
}
func (b *Block) setHash() {

	//data1 := append([]byte(b.version), b.prevHash...)
	//data1 = append(data1, b.data...)

	data1 := [][]byte{
		[]byte(b.version),
		b.prevHash,
		b.merkleRoot,
		dizi2byte(b.timeStamp),
		dizi2byte(b.bits),
		dizi2byte(b.nonce),
		b.data,
	}
	data := bytes.Join(data1, []byte(""))
	datahash := sha256.Sum256(data)
	b.hash = datahash[:]
}

//把int转化成[]byte
func dizi2byte(num int64) []byte {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.LittleEndian, num)
	if err != nil {
		fmt.Println("dizi2byte err:", err)
		return nil
	}
	return buf.Bytes()
}
