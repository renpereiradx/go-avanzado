package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Memory struct {
	f     Function
	cache map[int]FunctionResult
	mtx   sync.RWMutex
}

type Function func(key int) (interface{}, error)

type FunctionResult struct {
	value interface{}
	err   error
}

func NewCache(f Function) *Memory {
	memory := new(Memory)
	memory.f = f
	memory.cache = make(map[int]FunctionResult)
	return memory
}

func (m *Memory) Get(key int) (interface{}, error) {
	m.mtx.RLock()
	result, exist := m.cache[key]
	m.mtx.RUnlock()
	if !exist {
		m.mtx.Lock()
		result.value, result.err = m.f(key)
		m.cache[key] = result
		m.mtx.Unlock()
	}
	return result.value, result.err
}

func GetFibonacci(n int) (interface{}, error) {
	return Fibonacci(n), nil
}

func main() {
	cache := NewCache(GetFibonacci)
	fibo := []int{42, 40, 41, 42, 46, 46}
	var wg sync.WaitGroup
	for _, n := range fibo {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			start := time.Now()
			value, err := cache.Get(index)
			if err != nil {
				log.Println(err)
			}
			fmt.Printf("Fibonacci de %d es %d. Tiempo en ser procesado %s\n", index, value, time.Since(start))
		}(n)
	}
	wg.Wait()
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
