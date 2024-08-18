package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/H1rono/goraii"
)

func main() {
	mu := &sync.Mutex{}
	c := 0
	wg := &sync.WaitGroup{}
	// increment
	wg.Add(1)
	go func() {
		for i := range 5 {
			for range goraii.MutexLockGuard(mu) {
				c += i
			}
			time.Sleep(10 * time.Millisecond)
		}
		wg.Done()
	}()
	// print
	wg.Add(1)
	go func() {
		for range 5 {
			for range goraii.MutexLockGuard(mu) {
				fmt.Println(c)
			}
			time.Sleep(10 * time.Millisecond)
		}
		wg.Done()
	}()
	wg.Wait()
}
