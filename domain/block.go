package domain

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"log"
)

type Block struct {
	Hash     []byte
	Nonce    []byte
	Data     []byte
	PrevHash []byte
}

func NewBlock(data map[string]string, prevHash string) *Block {
	dataByte, err := json.Marshal(data)

	if err != nil {
		log.Fatalln("[ERROR NewBlock()]:", err.Error())
	}

	block := &Block{
		Data:     dataByte,
		PrevHash: []byte(prevHash),
	}

	return block
}

func (b *Block) HashBlock(nonce, dificulty []byte) ([]byte, error) {
	content := bytes.Join([][]byte{
		b.PrevHash,
		b.Data,
		nonce,
		dificulty,
	}, []byte{})

	hashBytes := sha256.Sum256(content)

	return hashBytes[:], nil
}
