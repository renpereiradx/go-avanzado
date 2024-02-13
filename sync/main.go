package main

import (
	"fmt"
	"sync"
)

var balance int = 100

func Depositar(amount int, wg *sync.WaitGroup, mu *sync.RWMutex) {
	defer wg.Done()
	mu.Lock()
	b := balance
	balance = b + amount
	mu.Unlock()
	fmt.Println("Deposited: ", amount)
}

func Balance(mu *sync.RWMutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.RLock()
	b := balance
	mu.RUnlock()
	fmt.Println("Current Balance is: ", b)
}

func main() {
	var wg sync.WaitGroup
	var mt sync.RWMutex
	for i := 1; i < 5; i++ {
		wg.Add(1)
		go Depositar(i*200, &wg, &mt)
		wg.Add(1)
		go Balance(&mt, &wg)
	}
	wg.Wait()
	fmt.Println("Finish Balance is: ", balance)
}
