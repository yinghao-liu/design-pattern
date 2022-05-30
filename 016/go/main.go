package main

import "fmt"

func main() {
	var collect = CollectionSlice{1, 2, 3, 4, 5}
	for iter := collect.GetIterator(); iter.HasMore(); {
		i := iter.GetNext()
		fmt.Printf("i:%v\n", i)
	}
	return
}
