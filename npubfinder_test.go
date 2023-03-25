package npubfinder_test

import (
	"encoding/hex"
	"github.com/boreq/npubfinder"
	"testing"
)

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		npub, sk, ok := npubfinder.Generate("test")
		if ok {
			b.Log(npub, hex.EncodeToString(sk))
		}
	}
}
