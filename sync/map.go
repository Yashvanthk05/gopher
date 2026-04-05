package main

import (
	"fmt"
	"sync"
)

func main(){
	// sync map - concurrent map structure without requiring for mutex and map separately
	// used for key based read heavy and write once purpose
	var mp sync.Map

	// store
	mp.Store(1,"Yashvanth")

	// load
	val, ok := mp.Load(1)
	if ok {
		fmt.Println(val)
	}

	// load-store - checks whether key exist if not stores
	// LoadOrStore returns the existing value for the key if present. Otherwise, it stores and returns the given value. The loaded result is true if the value was loaded, false if stored.
	actual, loaded := mp.LoadOrStore(2, "Kavipriya")
	// loaded - false: value is stored
	// loaded - true: value is loaded
	if loaded {
		fmt.Println("Loaded:",actual)
	} else {
		fmt.Println("Stored:",actual)
	}

	// delete
	mp.Delete(1)

	// range
	mp.Store(1, "Kavin")
	mp.Store(2, "Deepak")
	mp.Store(1, "Yashvanth")

	mp.Range(func(key any, val any) bool {
		fmt.Println(key,val)
		return true // return false to stop iteration
	})
}