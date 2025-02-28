package main_test

import (
	//"sync"
	"os"
	"testing"
	"time"
)

func TestSimple(t *testing.T) {
	ctx := t.Context()

	os.WriteFile("testing", []byte{}, 0644)
	go func() {
		<-ctx.Done()
		os.Remove("testing")
	}()

	time.Sleep(3 * time.Second)
}
