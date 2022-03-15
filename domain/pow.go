package domain

import (
	"math"
	"math/big"

	"github.com/antunesgabriel/golang-blockchain/helpers"
)

type PoW struct {
	Block     *Block
	Target    *big.Int // base para dizer se um hash é considerado valido ou não
	Dificulty int
}

func NewPoW(b *Block, dificulty int) *PoW {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-dificulty))

	pow := &PoW{
		Block:     b,
		Target:    target,
		Dificulty: dificulty,
	}

	return pow
}

func (p *PoW) Mine() ([]byte, []byte, error) {
	var hash []byte
	var err error
	var nonceByte []byte

	dificultyByte, _ := helpers.IntToBytes(int(p.Dificulty))

	for nonce := 0; nonce < math.MaxInt64; nonce++ {
		nonceByte, _ = helpers.IntToBytes(nonce)

		if hash, err = p.Block.HashBlock(nonceByte, dificultyByte); err != nil {
			break
		}

		intHashToCompare := big.NewInt(0).SetBytes(hash)

		compare := intHashToCompare.Cmp(p.Target)

		if compare == -1 || compare == 0 { // menor ou igual ao limite estabelecido
			break
		}
	}

	return hash, nonceByte, err
}
