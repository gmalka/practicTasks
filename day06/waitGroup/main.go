package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type wg struct {
	park *sync.Mutex
	m *sync.Mutex
	count int32
}

func (w *wg) Wait() {
	w.park.Lock()
}

func (w *wg) Done() {
	w.m.Lock()
	w.count--
	if w.count <= 0 {
		w.park.Unlock()
		w.count = 0
	}
	w.m.Unlock()
}

func (w *wg) Add(i int32) {
	if i <= 0 { 
		i = 1
	}
	atomic.AddInt32(&w.count, i)
	w.park.TryLock()
}

func NewWaitGroup() wg {
	return wg{park: &sync.Mutex{}, m: &sync.Mutex{}, count: 0}
}

func main() {
	wg := NewWaitGroup()
	ch := make(chan int)
	count := 1000

	for i:= 1; i < count; i++ {
		wg.Add(1)
		go func(i int)  {
			for t := 0; t < count; t++ {
				ch <- i
			}
			wg.Done()
		}(i)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	m := make(map[int]int, 10)
	for k := range ch {
		m[k]++
	}

	for k, v := range m {
		if v != count {
			fmt.Println("ERROR, ", k, v)
			return
		}
	}

	fmt.Println("All ok")
}