package domain

import (
	"encoding/hex"
	"reflect"
	"strconv"
	"testing"
)

const DIFICULTY = 16
const PREV_HASH = "0"

func TestNewPoW(t *testing.T) {
	t.Run("i should create PoW", func(t *testing.T) {
		fakeData := map[string]string{}

		block := NewBlock(fakeData, PREV_HASH)
		pow := NewPoW(block, DIFICULTY)

		if pow.Dificulty != DIFICULTY {
			t.Errorf("[Failed NewPoW()]: Expected: %d ---- Received: %d", DIFICULTY, pow.Dificulty)
		}

		if !reflect.DeepEqual(block, pow.Block) {
			t.Errorf("[Failed NewPoW()]: Expected: %x ---- Received: %x", block, pow.Block)
		}
	})
}

func TestPow(t *testing.T) {
	t.Run("Mine()", func(t *testing.T) {
		t.Run("i should return correct nonce and block hash", func(t *testing.T) {
			expectedHash := "0000b386f937d3a2120a4200cc316786df5031672d82fba496e95075f48a56b2"
			expectedNonce := 618

			fakeData := map[string]string{
				"BestTime": "Vasco",
			}

			block := NewBlock(fakeData, PREV_HASH)
			pow := NewPoW(block, DIFICULTY)

			hashByte, nonce, err := pow.Mine()

			if err != nil {
				t.Errorf("[Failed NewPoW()]: Error: %s", err.Error())
			}

			hash := hex.EncodeToString(hashByte)

			if hash != expectedHash {
				t.Errorf("[Failed pow.Mine()]: Expected Hash: %s --- Received %s", expectedHash, hash)
			}

			if nonce != expectedNonce {
				t.Errorf("[Failed pow.Mine()]: Expected Hash: %d --- Received %d", expectedNonce, nonce)
			}
		})
	})
}

func BenchmarkPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fakeData := map[string]string{
			"Count": strconv.Itoa(i),
		}

		block := NewBlock(fakeData, PREV_HASH)
		pow := NewPoW(block, DIFICULTY)

		pow.Mine()
	}

}
