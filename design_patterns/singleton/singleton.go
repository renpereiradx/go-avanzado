package main

import (
	"fmt"
	"sync"
	"time"
)

type Database struct{}

func (Database) CreatingSingletonConection() {
	fmt.Println("Creating singleton for Database")
	time.Sleep(2 * time.Second)
	fmt.Println("Creation Done")
}

var db *Database
var mtx sync.Mutex

func GetDatabaseInstance() *Database {
	mtx.Lock()
	defer mtx.Unlock()
	if db == nil {
		fmt.Println("Creating DB Conection")
		db = &Database{}
		db.CreatingSingletonConection()
	} else {
		fmt.Println("DB Already Created")
	}
	return db
}

func main() {
	var wg sync.WaitGroup
	defer wg.Wait()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			GetDatabaseInstance()
		}()
	}
}
