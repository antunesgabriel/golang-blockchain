package domain

import (
	"encoding/hex"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/antunesgabriel/golang-blockchain/helpers"
)

func TestNewBlock(t *testing.T) {
	t.Run("i should return new block reference with correct params", func(t *testing.T) {
		data := map[string]string{
			"To":    "Marina",
			"From":  "Gabriel",
			"Value": "34.5",
		}

		prevHash := "0"

		block := NewBlock(data, prevHash)

		var myData map[string]string

		err := json.Unmarshal(block.Data, &myData)

		if err != nil {
			t.Errorf("[Failed]: %s", err.Error())
		}

		if !reflect.DeepEqual(data, myData) {
			t.Errorf("[Failed]: Expected: %s --- Received %s", data, myData)
		}

		if string(block.PrevHash) != prevHash {
			t.Errorf("[Failed]: Expected: %s --- Received %s", prevHash, string(block.PrevHash))
		}
	})
}

func TestHashBlock(t *testing.T) {
	t.Run("i should returns block content hash", func(t *testing.T) {
		data := map[string]string{
			"To":    "Mariana",
			"From":  "Gabriel",
			"Value": "11.0",
		}

		block := NewBlock(data, "0")
		nonce, _ := helpers.IntToBytes(1)
		hashBytes, err := block.HashBlock(nonce)

		if err != nil {
			t.Errorf("[Failed HashBlock()]: Received error: %s", err.Error())
		}

		result := hex.EncodeToString(hashBytes)

		if result == "" {
			t.Fatalf("[Failed block.HashBlock()]: Expected: hash --- Received: %s", result)
		}
	})
}

func BenchmarkNewBlock(b *testing.B) {
	data := map[string]string{
		"foo": "bzz",
	}
	prevHash := "dsfdsf"

	for i := 0; i < b.N; i++ {
		NewBlock(data, prevHash)
	}
}

func BenchmarkHashBlock(b *testing.B) {
	data := map[string]string{
		"foo": "bzz",
	}
	prevHash := "dsfdsf"

	block := NewBlock(data, prevHash)

	for i := 0; i < b.N; i++ {
		nonce, _ := helpers.IntToBytes(i)
		block.HashBlock(nonce)
	}
}
