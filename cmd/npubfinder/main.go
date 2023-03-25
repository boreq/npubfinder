package main

import (
	"encoding/hex"
	"fmt"
	"github.com/boreq/npubfinder"
	"os"
	"runtime"
	"sync"
)

func main() {
	numWorkers := runtime.NumCPU()

	wg := sync.WaitGroup{}
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			if err := generate(os.Args[1]); err != nil {
				panic(err)
			}
		}()
	}

	wg.Wait()
	return
}

func generate(phrase string) error {
	for {
		npub, sk, ok := npubfinder.Generate(phrase)
		if ok {
			fmt.Println(npub, hex.EncodeToString(sk))
		}
	}
}
