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
		expected := "a859f59f412eb42b495b945f8a5d0a158b3b7f314e9ba598cd41166a72d15154"

		data := map[string]string{
			"To":    "Mariana",
			"From":  "Gabriel",
			"Value": "11.0",
		}

		block := NewBlock(data, "0")
		dificulty, _ := helpers.IntToBytes(14)
		nonce, _ := helpers.IntToBytes(1)
		hashBytes, err := block.HashBlock(nonce, dificulty)

		if err != nil {
			t.Errorf("[Failed HashBlock()]: Received error: %s", err.Error())
		}

		result := hex.EncodeToString(hashBytes)

		if result != expected {
			t.Fatalf("[Failed block.HashBlock()]: Expected: %s --- Received: %s", expected, result)
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

	dificulty, _ := helpers.IntToBytes(14)

	block := NewBlock(data, prevHash)

	for i := 0; i < b.N; i++ {
		nonce, _ := helpers.IntToBytes(i)
		block.HashBlock(nonce, dificulty)
	}
}
