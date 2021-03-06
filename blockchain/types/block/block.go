package block

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/xoreo/go-basics/blockchain/common"
	"github.com/xoreo/go-basics/blockchain/types/transaction"
)

// ErrInvalidBlock - Error for an attempt to create a new block with invalid parameters
var ErrInvalidBlock = errors.New("invalid parameters to construct block")

// Block - A block in the chain
type Block struct {
	Index        int                        `json:"index"`
	Difficulty   int                        `json:"difficulty"`
	Nonce        []byte                     `json:"nonce"`
	Transactions []*transaction.Transaction `json:"transactions"`
	PrevHash     []byte                     `json:"prevHash"`
	Timestamp    string                     `json:"timestamp"`
	Hash         []byte                     `json:"hash"`
}

// NewBlock - Create a new block
func NewBlock(index int, transactions []*transaction.Transaction, prevHash []byte) (*Block, error) {
	if transactions == nil {
		return nil, ErrInvalidBlock
	}
	block := &Block{
		Index:        index,
		Difficulty:   common.Difficulty,
		Nonce:        common.GetNonce(50),
		Transactions: transactions,
		PrevHash:     prevHash,
		Timestamp:    time.Now().UTC().String(),
		Hash:         nil,
	}
	(*block).Hash = common.Sha3(block.Bytes())
	return block, nil
}

// String - Encode a block into a string
func (block *Block) String() string {
	json, _ := json.MarshalIndent(*block, "", "  ")
	return string(json)
}

// Bytes - Encode a block into a []byte
func (block *Block) Bytes() []byte {
	json, _ := json.MarshalIndent(*block, "", "  ")
	return json
}
