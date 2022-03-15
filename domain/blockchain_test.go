package domain

import (
	"encoding/binary"
	"testing"
)

func TestNewBlockChain(t *testing.T) {
	t.Run("i should return blockchain with genesis block", func(t *testing.T) {
		fakeData := map[string]string{
			"To":    "Ukraine",
			"Value": "1000000",
		}
		prevHash := "0"

		genesisBlock := NewBlock(fakeData, prevHash)

		blockchain := NewBlockChain(genesisBlock)

		if length := len(blockchain.Chain); length != 1 {
			t.Errorf("[Failed NewBlockChain()]: Expected: %d --- Received: %d", 1, length)
		}

		firstBlock := blockchain.Chain[0]

		result := string(firstBlock.PrevHash)

		if result != prevHash {
			t.Errorf("[Failed NewBlockChain()]: Expected: %s --- Received: %s", prevHash, result)
		}
	})
}

func TestBlockChain(t *testing.T) {
	t.Run("blockchain.GetLastBlock()", func(t *testing.T) {
		t.Run("i should return last block of chain", func(t *testing.T) {
			fakeData := map[string]string{}
			prevHash := "0000"

			genesisBlock := NewBlock(fakeData, prevHash)

			blockchain := NewBlockChain(genesisBlock)

			lastBlock := blockchain.GetLastBlock()

			result := string(lastBlock.PrevHash)

			if result != prevHash {
				t.Errorf("[Failed NewBlockChain()]: Expected: %s --- Received: %s", prevHash, result)
			}
		})
	})

	t.Run("blockcahin.AddBlock()", func(t *testing.T) {
		t.Run("i should add new block into chain with nonce and hash", func(t *testing.T) {
			genesisData := map[string]string{
				"To":    "Ukraine",
				"Value": "1000000",
			}

			prevHash := "0"

			blockchain := NewBlockChain(NewBlock(genesisData, prevHash))

			data := map[string]string{
				"From":    "Ukraine",
				"To":      "World",
				"Message": "s2",
			}

			var lastBlock = blockchain.GetLastBlock()

			block := NewBlock(data, string(lastBlock.Hash))

			blockchain.AddBlock(block)

			lastBlock = blockchain.GetLastBlock()

			if binary.BigEndian.Uint64(lastBlock.Nonce) == 0 {
				t.Error("[Failed AddBlock()]: Expected: > 0 --- Received: 0")
			}

			if string(lastBlock.Hash) == "" {
				t.Error("[Failed AddBlock()] Expected != '' --- Received: ''")
			}
		})
	})
}
