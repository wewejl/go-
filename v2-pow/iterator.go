package main

import "github.com/boltdb/bolt"

type Iterator struct {
	db *bolt.DB
	currhash []byte
}
//创建这个迭代器
func NewIterator(bc *Blockchain) *Iterator {
	return &Iterator{
		db:bc.db,
		currhash:bc.listHash,
	}
}

func (i *Iterator)next() *Block {
	var block *Block
	i.db.View(func(tx *bolt.Tx) error {
		b:=tx.Bucket([]byte(blockBucket))
		blockdata:=b.Get(i.currhash)
		block=DeSerialize(blockdata)
		i.currhash=block.PrevHash
		return nil
	})
	return block
}

