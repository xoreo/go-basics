package blockchain

import (
	"fmt"
	"testing"
)

func TestWriteBlockchainToMemory(t *testing.T) {
	chain, err := PopulateChain(10)
	if err != nil {
		t.Fatal(err)
	}

	err = chain.WriteBlockchainToMemory()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("success")
}

func TestReadBlockchainFromMemory(t *testing.T) {
	chain, err := PopulateChain(10)
	if err != nil {
		t.Fatal(err)
	}
	err = chain.WriteBlockchainToMemory()
	if err != nil {
		t.Fatal(err)
	}

	genesis := chain.Blocks[0]
	hexHash := fmt.Sprintf("%x", genesis.Hash)[:8]
	chain, err = ReadBlockchainFromMemory(hexHash)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(chain)
}
