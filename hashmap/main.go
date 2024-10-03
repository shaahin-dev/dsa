package main

import (
	"fmt"
	"hashmap/hashmap"
)

func main() {
	hashMap := hashmap.NewHashMap(10)

	hashMap.Insert("Shahin", 30)
	hashMap.Insert("Reza", 25)
	hashMap.Insert("Mohsen", 40)
	hashMap.Insert("Shahin", 35)
	hashMap.Insert("Hasan", 20)
	hashMap.Insert("Hosein", 20)
	hashMap.Insert("Hosein1", 20)

	hashMap.Display()

	value, found := hashMap.Get("Shahin")
	if found {
		fmt.Printf("Value for 'Shahin': %d\n", value)
	} else {
		fmt.Println("'Shahin' not found")
	}

	hashMap.Delete("Hasan")
	hashMap.Display()
}
