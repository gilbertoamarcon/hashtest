package main

import (
	"github.com/cnf/structhash"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/sha3"

	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"hash"
	"hash/fnv"
	"math/rand"
	"time"
)

type state struct {
	batteryState   batteryState
	stepsCharging  int64
	lastDeepCharge int64
	step           int64
}

type batteryState struct {
	totalCapacity   float64
	currentCapacity float64
	temperature     float64
}

func genStates(size int) []*state {
	states := make([]*state, size)
	for i := 0; i < size; i++ {
		states[i] = genState()
	}
	return states
}

func genState() *state {
	batteryState := genBatteryState()
	state := &state{
		batteryState:   *batteryState,
		stepsCharging:  int64(rand.Intn(24 * 7 * 4)),
		lastDeepCharge: int64(rand.Intn(24 * 7 * 4)),
		step:           int64(rand.Intn(24 * 7 * 4)),
	}
	return state
}

func genBatteryState() *batteryState {
	state := &batteryState{
		totalCapacity:   50.0 * rand.Float64(),
		currentCapacity: 50.0 * rand.Float64(),
		temperature:     50.0 * rand.Float64(),
	}
	return state
}

type hasher interface {
	encode(state *state) string
}

type hasheTester struct {
	hashers map[string]hasher
	states  []*state
}

func (hasheTester *hasheTester) run() {
	for name, hasher := range hasheTester.hashers {
		success, duration, size := hasheTester.test(hasher)
		fmt.Printf("%s: %v %v %v\n", name, success, duration, size)
	}
}

func (hasheTester *hasheTester) test(hasher hasher) (bool, time.Duration, int) {
	start := time.Now()
	hashSet := map[string]bool{}
	size := 0
	for _, state := range hasheTester.states {
		hash := hasher.encode(state)
		hashSet[hash] = true
		size += len(hash)
	}
	success := bool(len(hashSet) == len(hasheTester.states))
	duration := time.Now().Sub(start)
	return success, duration, size
}

func encode(state *state) []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.LittleEndian, state)
	return buffer.Bytes()
}

type StructhashMd5 struct{}

func (encoder *StructhashMd5) encode(state *state) string {
	return string(structhash.Md5(state, 1))
}

type Fnv128 struct{}

func (encoder *Fnv128) encode(state *state) string {
	return CryptoHasher{hasher: fnv.New128()}.encode(state)
}

type Fnv128a struct{}

func (encoder *Fnv128a) encode(state *state) string {
	return CryptoHasher{hasher: fnv.New128a()}.encode(state)
}

type Fnv32 struct{}

func (encoder *Fnv32) encode(state *state) string {
	return CryptoHasher{hasher: fnv.New32()}.encode(state)
}

type Sha1 struct{}

func (encoder *Sha1) encode(state *state) string {
	return CryptoHasher{hasher: sha1.New()}.encode(state)
}

type Sha3 struct{}

func (encoder *Sha3) encode(state *state) string {
	return CryptoHasher{hasher: sha3.New256()}.encode(state)
}

type Sha256 struct{}

func (encoder *Sha256) encode(state *state) string {
	return CryptoHasher{hasher: sha256.New()}.encode(state)
}

type Md5 struct{}

func (encoder *Md5) encode(state *state) string {
	return CryptoHasher{hasher: md5.New()}.encode(state)
}

type Md4 struct{}

func (encoder *Md4) encode(state *state) string {
	return CryptoHasher{hasher: md4.New()}.encode(state)
}

type Blake2b struct{}

func (encoder *Blake2b) encode(state *state) string {
	hasher, _ := blake2b.New256(nil)
	return CryptoHasher{hasher: hasher}.encode(state)
}

type CryptoHasher struct {
	hasher hash.Hash
}

func (cryptoHasher CryptoHasher) encode(state *state) string {
	buffer := encode(state)
	cryptoHasher.hasher.Write(buffer)
	return string(cryptoHasher.hasher.Sum(nil))
}

func main() {
	hashers := map[string]hasher{
		"StructhashMd5": &StructhashMd5{},
		"Fnv32":         &Fnv32{},
		"Fnv128":        &Fnv128{},
		"Fnv128a":       &Fnv128a{},
		"Sha1":          &Sha1{},
		"Sha3":          &Sha3{},
		"Sha256":        &Sha256{},
		"Md5":           &Md5{},
		"Md4":           &Md4{},
		"Blake2b":       &Blake2b{},
	}
	tester := hasheTester{
		hashers: hashers,
		states:  genStates(100000),
	}
	tester.run()
}
