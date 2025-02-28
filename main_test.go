package main_test

import (
	"sync"
	"testing"
	"time"
)

func TestSimple(t *testing.T) {
	ctx := t.Context()

	var wg sync.WaitGroup
	t.Cleanup(wg.Wait)

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()

			println("start")
			<-ctx.Done()
			println("shutdown")
		}()
	}

	time.Sleep(3 * time.Second)
}
