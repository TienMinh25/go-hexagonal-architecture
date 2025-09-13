package hashing

import (
	"golang.org/x/crypto/blake2b"
)

// GenericHash The same as crypto_generichash as referred to in the Paseto spec
func GenericHash(in, out, key []byte) {
	blake, err := blake2b.New(len(out), key)
	if err != nil {
		panic(err)
	}

	if _, err := blake.Write(in); err != nil {
		panic(err)
	}

	copy(out, blake.Sum(nil))
}
