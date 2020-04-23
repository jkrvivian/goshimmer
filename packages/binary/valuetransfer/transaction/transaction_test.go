package transaction

import (
	"bytes"
	"strings"
	"testing"

	"strings"
	"testing"

	"github.com/iotaledger/goshimmer/packages/binary/valuetransfer/address"
	"github.com/iotaledger/goshimmer/packages/binary/valuetransfer/address/signaturescheme"
	"github.com/iotaledger/goshimmer/packages/binary/valuetransfer/balance"
	"github.com/iotaledger/hive.go/crypto/ed25519"
	"github.com/stretchr/testify/assert"
)

func TestEmptyDataPayload(t *testing.T) {
	sigScheme := signaturescheme.ED25519(ed25519.GenerateKeyPair())
	addr := sigScheme.Address()
	o1 := NewOutputId(addr, RandomId())
	inputs := NewInputs(o1)
	bal := balance.New(balance.COLOR_IOTA, 1)
	outputs := NewOutputs(map[address.Address][]*balance.Balance{addr: {bal}})
	tx := New(inputs, outputs)
	tx.Sign(sigScheme)
	check := tx.SignaturesValid()

	assert.Equal(t, true, check)
}

func TestShortDataPayload(t *testing.T) {
	sigScheme := signaturescheme.ED25519(ed25519.GenerateKeyPair())
	addr := sigScheme.Address()
	o1 := NewOutputId(addr, RandomId())
	inputs := NewInputs(o1)
	bal := balance.New(balance.COLOR_IOTA, 1)
	outputs := NewOutputs(map[address.Address][]*balance.Balance{addr: {bal}})
	tx := New(inputs, outputs)

	dataPayload := []byte("data payload test")
	err := tx.SetDataPayload(dataPayload)
	assert.Equal(t, nil, err)

	dpBack := tx.GetDataPayload()
	assert.Equal(t, true, bytes.Equal(dpBack, dataPayload))

	tx.Sign(sigScheme)
	check := tx.SignaturesValid()
	assert.Equal(t, true, check)

	// corrupt data payload bytes
	// reset essence to force recalculation
	tx.essenceBytes = nil
	dataPayload[2] = '?'
	err = tx.SetDataPayload(dataPayload)
	assert.Equal(t, nil, err)

	// expect signature is not valid
	check = tx.SignaturesValid()
	assert.Equal(t, false, check)
}

func TestTooLongDataPayload(t *testing.T) {
	sigScheme := signaturescheme.ED25519(ed25519.GenerateKeyPair())
	addr := sigScheme.Address()
	o1 := NewOutputId(addr, RandomId())
	inputs := NewInputs(o1)
	bal := balance.New(balance.COLOR_IOTA, 1)
	outputs := NewOutputs(map[address.Address][]*balance.Balance{addr: {bal}})
	tx := New(inputs, outputs)

	dataPayload := []byte(strings.Repeat("1", MAX_DATA_PAYLOAD_SIZE+1))
	err := tx.SetDataPayload(dataPayload)
	assert.Equal(t, true, err != nil)
}

func TestMarshalingEmptyDataPayload(t *testing.T) {
	sigScheme := signaturescheme.RandBLS()
	addr := sigScheme.Address()
	o1 := NewOutputId(addr, RandomId())
	inputs := NewInputs(o1)
	bal := balance.New(balance.COLOR_IOTA, 1)
	outputs := NewOutputs(map[address.Address][]*balance.Balance{addr: {bal}})
	tx := New(inputs, outputs)

	tx.Sign(sigScheme)
	check := tx.SignaturesValid()
	assert.Equal(t, true, check)

	v := tx.ObjectStorageValue()

	tx1 := Transaction{}
	err, _ := tx1.UnmarshalObjectStorageValue(v)
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, true, tx1.SignaturesValid())
	assert.Equal(t, true, bytes.Equal(tx1.Id().Bytes(), tx.Id().Bytes()))
}

func TestMarshalingDataPayload(t *testing.T) {
	sigScheme := signaturescheme.RandBLS()
	addr := sigScheme.Address()
	o1 := NewOutputId(addr, RandomId())
	inputs := NewInputs(o1)
	bal := balance.New(balance.COLOR_IOTA, 1)
	outputs := NewOutputs(map[address.Address][]*balance.Balance{addr: {bal}})
	tx := New(inputs, outputs)

	dataPayload := []byte("data payload test")
	err := tx.SetDataPayload(dataPayload)
	assert.Equal(t, nil, err)

	tx.Sign(sigScheme)
	check := tx.SignaturesValid()
	assert.Equal(t, true, check)

	v := tx.ObjectStorageValue()

	tx1 := Transaction{}
	err, _ = tx1.UnmarshalObjectStorageValue(v)

	assert.Equal(t, nil, err)
	assert.Equal(t, true, tx1.SignaturesValid())

	assert.Equal(t, true, bytes.Equal(tx1.Id().Bytes(), tx.Id().Bytes()))
}
