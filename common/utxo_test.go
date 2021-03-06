package common

import (
	"testing"

	"github.com/MixinNetwork/mixin/crypto"
	"github.com/stretchr/testify/assert"
)

func TestUTXO(t *testing.T) {
	assert := assert.New(t)

	s := &Snapshot{}
	utxos := s.UnspentOutputs()
	assert.Len(utxos, 0)

	genesisHash := crypto.Hash{}
	script := Script{OperatorCmp, OperatorSum, 2}
	accounts := make([]Address, 0)
	for i := 0; i < 3; i++ {
		accounts = append(accounts, randomAccount())
	}

	tx := NewTransaction(XINAssetId)
	tx.AddInput(genesisHash, 0)
	tx.AddInput(genesisHash, 1)
	tx.AddScriptOutput(accounts, script, NewInteger(20000))
	s.Transaction = &SignedTransaction{
		Transaction: *tx,
	}

	utxos = s.UnspentOutputs()
	assert.Len(utxos, 1)
	utxo := utxos[0]
	assert.Equal(tx.Hash(), utxo.Input.Hash)
	assert.Equal(0, utxo.Input.Index)
	assert.Equal(uint8(OutputTypeScript), utxo.Output.Type)
	assert.Equal("20000.00000000", utxo.Output.Amount.String())
	assert.Equal("fffe02", utxo.Output.Script.String())
	assert.Len(utxo.Output.Keys, 3)
	assert.Equal(XINAssetId, utxo.Asset)
}
