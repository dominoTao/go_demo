package share_mem

import (
	"sync"
	"testing"
	"time"
)

func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i:= 0; i < 5000; i++ {  // 定义5000 个协程
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
				wg.Done()
			}()
			mut.Lock()
			counter++
			//wg.Done()   // 也可以
		}()
	}
	//time.Sleep( 1 * time.Second) // 用栅格替换   waitgroup
	wg.Wait()
	t.Logf("counter = %d", counter)
}

func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i:= 0; i < 5000; i++ {  // 定义5000 个协程
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep( 1 * time.Second) // 用栅格替换   waitgroup

	t.Logf("counter = %d", counter)
}

func TestCounter(t *testing.T) {
	counter := 0
	for i:= 0; i < 5000; i++ {  // 定义5000 个协程
		go func() {
			counter++
		}()
	}
	time.Sleep( 1 * time.Second)
	t.Logf("counter = %d", counter)
}
