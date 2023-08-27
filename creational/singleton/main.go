package main

import (
	"fmt"
	"sync"
)

// make an empty struct, like a static instance
type DB struct {
	Port int
}

// have to aquire locks if we are working in an multithreaded environment with the same instance return value
var lock = &sync.Mutex{}

// initiate an instance
var db *DB

// the get instance checks if db is nil
// if the db is nil then it aquires a lock
// then if the db is still nil then it just instanciates the instance with the DB object
// else it just returns a message that the db is already instanciated
// by locking we ensure that the db does not get instanciated at the same time when two threads are doing the same thing at the same time
func getInstance() *DB {
	if db == nil {
		lock.Lock()
		defer lock.Unlock()

		if db == nil {
			fmt.Println("creating a db instance")
			db = &DB{
				Port: 8080,
			}
		} else {
			fmt.Println("internal else: db instance already created", db)
		}
	} else {
		fmt.Println("external else: db instance already created", db)
	}
	return db
}

func main() {
	for i := 0; i < 10; i = i + 1 {
		go getInstance()
	}
	fmt.Scanln()
}
