package domain

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"log"
	"time"

	"github.com/antunesgabriel/golang-blockchain/helpers"
)

type Block struct {
	Timestamp []byte
	Hash      []byte
	Nonce     []byte
	Data      []byte
	PrevHash  []byte
}

func NewBlock(data map[string]string, prevHash string) *Block {
	dataByte, err := json.Marshal(data)

	if err != nil {
		log.Fatalln("[ERROR NewBlock()]:", err.Error())
	}

	timestampByte, _ := helpers.IntToBytes(int(time.Now().Unix()))

	block := &Block{
		Data:      dataByte,
		PrevHash:  []byte(prevHash),
		Timestamp: timestampByte,
	}

	return block
}

func (b *Block) HashBlock(nonce []byte) ([]byte, error) {
	timestampByte, _ := helpers.IntToBytes(int(time.Now().Unix()))
	b.Timestamp = timestampByte

	content := bytes.Join([][]byte{
		b.PrevHash,
		b.Data,
		nonce,
		b.Timestamp,
	}, []byte{})

	hashBytes := sha256.Sum256(content)

	return hashBytes[:], nil
}
