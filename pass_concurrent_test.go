package pass

import (
	"sync"
	"testing"
)

func TestNowConcurrent(t *testing.T) {
	var counter int

	exitOld := exit
	exit = func() {
		counter++
	}
	t.Cleanup(func() {
		exit = exitOld
	})

	const maxConcurrency = 100
	const loops = 10

	wg := sync.WaitGroup{}
	wg.Add(maxConcurrency)
	for i := 0; i < maxConcurrency; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < loops; j++ {
				Now(t)
			}
		}()
	}

	wg.Wait()

	if counter != maxConcurrency*loops {
		t.Fatalf("invalid value for counter (data race condition), expected %d, got: %d", maxConcurrency*loops, counter)
	}
}
