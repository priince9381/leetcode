package gobasics

import (
	"fmt"
	"sync"
)

func reader(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		val, ok := <-ch
		if !ok {
			fmt.Println("Channel is Close")
			return
		}
		fmt.Printf("Reader id %d and Value Readed %d\n", id, val)
	}
}

func TestChannel() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)
	go reader(1, ch, wg)
	go reader(2, ch, wg)
	wg.Add(2)
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
	wg.Wait()
}
