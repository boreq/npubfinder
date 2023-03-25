package npubfinder

import (
	"crypto/rand"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"io"
	"math/big"
	"strings"
)

type PrivateKey []byte
type PublicKey []byte
type NPub string

func Check(npub NPub, phrase string) bool {
	npubS := normalize(string(npub))
	phrase = normalize(phrase)

	if strings.Contains(npubS, phrase) {
		return true
	}

	return false
}

func normalize(s string) string {
	s = strings.ToLower(s)

	s = strings.ReplaceAll(s, "o", "0")
	s = strings.ReplaceAll(s, "i", "1")
	//s = strings.ReplaceAll(s, "i", "2")
	s = strings.ReplaceAll(s, "e", "3")
	s = strings.ReplaceAll(s, "a", "4")
	s = strings.ReplaceAll(s, "s", "5")
	//s = strings.ReplaceAll(s, "i", "6")
	//s = strings.ReplaceAll(s, "a", "7")
	//s = strings.ReplaceAll(s, "a", "8")
	//s = strings.ReplaceAll(s, "a", "9")
	return s
}

func GeneratePrivateKey() (PrivateKey, error) {
	params := btcec.S256().Params()
	one := new(big.Int).SetInt64(1)

	b := make([]byte, params.BitSize/8+8)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		return nil, err
	}

	k := new(big.Int).SetBytes(b)
	n := new(big.Int).Sub(params.N, one)
	k.Mod(k, n)
	k.Add(k, one)

	return k.Bytes(), nil
}

func GetPublicKey(b PrivateKey) PublicKey {
	_, pk := btcec.PrivKeyFromBytes(b)
	return schnorr.SerializePubKey(pk)
}

func EncodePublicKey(b PublicKey) (NPub, error) {
	bits5, err := convertBits(b, 8, 5, true)
	if err != nil {
		return "", err
	}

	return encode("npub", bits5)
}

func Generate(phrase string) (NPub, PrivateKey, bool) {
	sk, err := GeneratePrivateKey()
	if err != nil {
		panic(err)
	}

	pk := GetPublicKey(sk)

	npub, err := EncodePublicKey(pk)
	if err != nil {
		panic(err)
	}

	if Check(npub, phrase) {
		return npub, sk, true
	}

	return "", nil, false
}
