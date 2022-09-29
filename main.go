package main

import (
	"github.com/gilbertoamarcon/hashtest/hasher"
	"github.com/gilbertoamarcon/hashtest/state"

	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/sha3"
	"hash/fnv"

	"fmt"
	"time"
)

type testPair struct {
	name   string
	hasher hasher.Hasher
}

type tester struct {
	testPairs []testPair
	testSet   []*state.State
}

func (tester *tester) run() {
	for _, testPair := range tester.testPairs {
		success, duration, size := tester.test(testPair.hasher)
		fmt.Printf("%s: %v %v %v\n", testPair.name, success, duration, size)
	}
}

func (tester *tester) test(hasher hasher.Hasher) (bool, time.Duration, int) {
	start := time.Now()
	hashSet := map[string]bool{}
	size := 0
	for _, state := range tester.testSet {
		hash := hasher.Encode(state)
		hashSet[hash] = true
		size += len(hash)
	}
	success := bool(len(hashSet) == len(tester.testSet))
	duration := time.Now().Sub(start)
	return success, duration, size
}

func main() {
	numSamples := 1000000
	blake2bHasher, _ := blake2b.New256(nil)
	testPairs := []testPair{
		testPair{name: "StructhashMD5", hasher: &hasher.StructhashMD5{}},
		testPair{name: "Fnv32", hasher: &hasher.BuiltIn{Hasher: fnv.New32()}},
		testPair{name: "Fnv32a", hasher: &hasher.BuiltIn{Hasher: fnv.New32a()}},
		testPair{name: "Fnv64", hasher: &hasher.BuiltIn{Hasher: fnv.New64()}},
		testPair{name: "Fnv64a", hasher: &hasher.BuiltIn{Hasher: fnv.New64a()}},
		testPair{name: "Fnv128", hasher: &hasher.BuiltIn{Hasher: fnv.New128()}},
		testPair{name: "Fnv128a", hasher: &hasher.BuiltIn{Hasher: fnv.New128a()}},
		testPair{name: "Sha1", hasher: &hasher.BuiltIn{Hasher: sha1.New()}},
		testPair{name: "Sha3", hasher: &hasher.BuiltIn{Hasher: sha3.New256()}},
		testPair{name: "Sha256", hasher: &hasher.BuiltIn{Hasher: sha256.New()}},
		testPair{name: "Md4", hasher: &hasher.BuiltIn{Hasher: md4.New()}},
		testPair{name: "Md5", hasher: &hasher.BuiltIn{Hasher: md5.New()}},
		testPair{name: "Blake2b", hasher: &hasher.BuiltIn{Hasher: blake2bHasher}},
	}
	tester := tester{
		testPairs: testPairs,
		testSet:   state.GenStates(numSamples),
	}
	tester.run()
}
