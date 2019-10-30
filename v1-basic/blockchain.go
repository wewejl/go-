package main

//定义区块链
type Blockchain struct {
	blocks []*Block
}

const genesisInfo = "天气转凉，注意御寒！"

//创建blockchain，同时添加一个创世快
func NewBlockchain() *Blockchain {
	genesisBlock := NewBlock(genesisInfo, nil)
	return &Blockchain{
		blocks: []*Block{
			genesisBlock,
		},
	}
}

//添加区块的方法
func (bc *Blockchain) addBlockchain(data string) {
	prevHash := bc.blocks[len(bc.blocks)-1].hash

	//创建新的区块
	newblock := NewBlock(data, prevHash)

	bc.blocks = append(bc.blocks, newblock)
}
