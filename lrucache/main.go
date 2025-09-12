package main

import "fmt"

func main() {
	fmt.Println("LRUCache")

	lrucache := NewLRUCache[int, string](3)
	lrucache.Put(1, "Value 1")
	lrucache.Put(2, "Value 2")
	lrucache.Put(3, "Value 3")

	if val, exist := lrucache.Get(2); exist {
		fmt.Println(val)
	}

	if val, exist := lrucache.Get(3); exist {
		fmt.Println(val)
	}

	lrucache.Put(4, "Value 4")
	if val, exists := lrucache.Get(3); exists {
		fmt.Println(val)
	} else {
		fmt.Println("Value 3 was evicted") // Output: Value 3 was evicted
	}

	if val, exists := lrucache.Get(4); exists {
		fmt.Println(val) // Output: Value 4
	}

	lrucache.Put(2, "Updated Value 2")

	// Get the values again
	if val, exists := lrucache.Get(1); exists {
		fmt.Println(val) // Output: Value 1
	}
	if val, exists := lrucache.Get(2); exists {
		fmt.Println(val) // Output: Updated Value 2
	}
}
