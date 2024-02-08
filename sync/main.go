package main

import (
	"fmt"
	"sync"
)

var balance int = 100

func Depositar(amount int, wg *sync.WaitGroup /* mu *sync.Mutex */) {
	defer wg.Done()
	// mu.Lock()
	b := balance
	balance = b + amount
	// mu.Unlock()
}

func Balance() int {
	return balance
}

func main() {
	var wg sync.WaitGroup
	// var mt sync.Mutex
	for i := 1; i < 5; i++ {
		wg.Add(1)
		go Depositar(i*100, &wg /* &mt */)
	}
	wg.Wait()
	fmt.Println(Balance())
}
