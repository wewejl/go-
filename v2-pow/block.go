package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"time"
)

//定义区块结构

//区块头
type Block struct {
	//版本号
	Version string
	//前区块哈希
	PrevHash []byte

	//当前区块的哈希
	//在比特币定义的区块中，是没有这个当前的区块哈希这个字段的
	//为了方便
	Hash []byte

	//merkle根，根据当前区块的交易数据计算出来的
	MerkleRoot []byte

	//时间戳
	TimeStamp int64

	//难度值,系统提供的
	Bits int64

	//随机数
	Nonce int64

	//区块体，交易数据
	Data []byte
}

func NewBlock(data string, preHash []byte) *Block {

	newblock := &Block{
		Version:  "0",
		PrevHash: preHash,

		MerkleRoot: nil,
		TimeStamp:  time.Now().Unix(),
		Bits:       0,
		Nonce:      0,
		Data:       []byte(data),
	}
	//setHash是没有挖矿  临时的
	//newblock.setHash()
	pfd := NewProofOfWord(newblock)
	hash, nonce := pfd.Run()
	newblock.Nonce = nonce
	newblock.Hash = hash
	fmt.Println("创建的时候newblock",newblock.Hash)
	return newblock
}

//区块字节流序列化
func (b *Block) Serialize() []byte {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer) //创建编码器
	err := encoder.Encode(b)
	if err != nil {
		fmt.Println("区块字节流 序列化 err", err)
		return nil
	}
	return buffer.Bytes()
}

//区块字节流反序列化
func DeSerialize(buf []byte) *Block {
	var block Block
	Decoder := gob.NewDecoder(bytes.NewReader(buf)) //创建解码器
	err := Decoder.Decode(&block)
	if err != nil {
		fmt.Println("区块字节流 反序列化 err", err)
		return nil
	}
	return &block
}
func (b *Block) setHash() {

	//data1 := append([]byte(b.version), b.prevHash...)
	//data1 = append(data1, b.data...)

	data1 := [][]byte{
		[]byte(b.Version),
		b.PrevHash,
		b.MerkleRoot,
		dizi2byte(b.TimeStamp),
		dizi2byte(b.Bits),
		dizi2byte(b.Nonce),
		b.Data,
	}
	data := bytes.Join(data1, []byte(""))
	datahash := sha256.Sum256(data)
	b.Hash = datahash[:]
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
