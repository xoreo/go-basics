package transaction

import (
	"encoding/json"
	"errors"

	"github.com/xoreo/go-basics/blockchain/common"
)

// ErrInvalidTransaction - Error for an attempt to create a new transaction with invalid parameters
var ErrInvalidTransaction = errors.New("invalid parameters to construct transaction")

// Transaction - A transaction in a block's list of transactions
type Transaction struct {
	Sender    string  `json:"sender"`
	Recipient string  `json:"recipient"`
	Amount    float64 `json:"amount"`
	Hash      []byte  `json:"hash"`
	Nonce     []byte  `json:"nonce"`
}

// NewTransaction - Create a new transaction
func NewTransaction(sender string, recipient string, amount float64) (*Transaction, error) {
	if sender == "" || recipient == "" || amount == 0 {
		return nil, ErrInvalidTransaction
	}

	transaction := &Transaction{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
		Hash:      nil,
		Nonce:     common.GetNonce(50),
	}
	(*transaction).Hash = common.Sha3(transaction.Bytes())
	return transaction, nil
}

// String - Encode a transaction into a string
func (transaction *Transaction) String() string {
	json, _ := json.MarshalIndent(*transaction, "", "  ")
	return string(json)
}

// Bytes - Encode a transaction into a []byte
func (transaction *Transaction) Bytes() []byte {
	json, _ := json.MarshalIndent(*transaction, "", "  ")
	return json
}

// FromBytes - Create a transaction from a []byte
func FromBytes(b []byte) (*Transaction, error) {
	buffer := &Transaction{}
	err := json.Unmarshal(b, buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}
