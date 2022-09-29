package hasher

import (
	"github.com/gilbertoamarcon/hashtest/state"

	"github.com/cnf/structhash"

	"bytes"
	"encoding/binary"
	"hash"
)

// Hasher is a generic hasher interface that gets a state and hashes it into a string.
type Hasher interface {
	Encode(state *state.State) string
}

// StructhashMD5 is a wrapper to the structhash hasher.
type StructhashMD5 struct{}

func (encoder *StructhashMD5) Encode(state *state.State) string {
	return string(structhash.Md5(state, 1))
}

// BuiltIn is wrapper to a golang built-in type of hasher.
type BuiltIn struct {
	Hasher hash.Hash
}

func (cryptoHasher BuiltIn) Encode(state *state.State) string {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, state)
	cryptoHasher.Hasher.Write(buffer.Bytes())
	return string(cryptoHasher.Hasher.Sum(nil))
}
