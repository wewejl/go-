package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

//这个是一个pow工作证明的

type ProofOfWork struct {
	//区块
	block *Block
	//系统难度值
	target *big.Int
}


//提供一个创建的方法
func NewProofOfWord(block *Block) *ProofOfWork  {
	NewProofOfWord:= &ProofOfWork{
		block:block,
	}
	targetStr:="0001000000000000000000000000000000000000000000000000000000000000"
	tarint:=big.Int{}
	tarint.SetString(targetStr,16)
	NewProofOfWord.target= &tarint
	return NewProofOfWord
}

//先要进行变量拼接
func (pf *ProofOfWork)Prepardata(nonce int64) []byte {
	b:=pf.block
	data1 := [][]byte{
		[]byte(b.Version),
		b.PrevHash,
		b.MerkleRoot,
		dizi2byte(b.TimeStamp),
		dizi2byte(b.Bits),
		dizi2byte(nonce),
		b.Data,
	}
	data := bytes.Join(data1, []byte(""))
	return data
}


//核心函数 run  挖矿的函数
func (pf *ProofOfWork)Run() ([]byte,int64) {
	//定义一个全局的nonce的变量
	var nonce int64
	var datahash [32]byte
	for  {
		
		blockdata:=pf.Prepardata(nonce)
		//数据进行计算
		datahash=sha256.Sum256(blockdata)
		//拿到系统难度
		tpman:=big.Int{}
		tpman.SetBytes(datahash[:])
		//   -1 if x <  y
		//    0 if x == y
		//   +1 if x >  y
		if tpman.Cmp(pf.target)==-1 {
			fmt.Printf("挖矿成功--完成的hash：%x\n,随机数是:%d\n",datahash,nonce)
			break
		}else {
			//fmt.Printf("挖矿失败：hash：%x\n,随机数是:%d\n",datahash,nonce)
			nonce++
		}
	}
	return datahash[:],nonce
}

//添加一个验证的挖矿的验证
func (pf *ProofOfWork)ivVMord(nonce int64) bool {
		proodata:=pf.Prepardata(nonce)
		hash:=sha256.Sum256(proodata)
		tpman:=big.Int{}
		tpman.SetBytes(hash[:])
		if tpman.Cmp(pf.target)==-1{
			return true
		}
		return false
}