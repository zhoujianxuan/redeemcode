package redeemcode

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"testing"
)

var incr = atomic.Int64{}

func TestGen(t *testing.T) {
	incr.Store(895151)
	for i := 0; i < 100; i++ {
		fmt.Println(Gen(int(incr.Add(99))))
	}
}

func TestSyncGen(t *testing.T) {
	incr.Store(688917)
	wg := sync.WaitGroup{}
	wg.Add(10)
	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 10000*100; i++ {
				num := incr.Add(1)
				Gen(int(num))
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestLargeGen(t *testing.T) {
	incr.Store(99)
	codeMap := map[string]int64{}
	for i := 0; i < 10000*100; i++ {
		num := incr.Add(1)
		code := Gen(int(num))
		if _, ok := codeMap[code]; ok {
			//repeat
			log.Println("repeat:", code)
		}
		codeMap[code] = num
	}
}
