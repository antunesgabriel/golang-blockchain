package domain

type BlockChain struct {
	Chain []*Block
}

const POW_DIFICULTY = 17

func NewBlockChain(genesisBlock *Block) *BlockChain {

	blockchain := &BlockChain{
		Chain: []*Block{
			genesisBlock,
		},
	}

	return blockchain
}

func (b *BlockChain) GetLastBlock() *Block {
	lastBlock := b.Chain[len(b.Chain)-1]

	return lastBlock
}

func (b *BlockChain) AddBlock(block *Block) error {
	pow := NewPoW(block, POW_DIFICULTY)

	hash, nonce, err := pow.Mine()

	if err != nil {
		return err
	}

	block.Hash = hash
	block.Nonce = nonce

	b.Chain = append(b.Chain, block)

	return nil
}
