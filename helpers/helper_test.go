package helpers

import (
	"math/big"
	"testing"
)

func TestIntToBytes(t *testing.T) {
	t.Run("i should convert int to bytes", func(t *testing.T) {
		expected := 10

		result, err := IntToBytes(expected)

		if err != nil {
			t.Errorf("[Failed IntToBytes()]: Received Error: %s", err.Error())
		}

		num := big.NewInt(0).SetBytes(result).Int64()

		if num != int64(expected) {
			t.Errorf("[Failed IntToBytes()]: Expected: %d --- Received: %d", expected, num)
		}

	})
}

func BenchmarkIntToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntToBytes(i)
	}
}
