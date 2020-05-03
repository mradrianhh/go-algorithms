package search

import "fmt"

// Linearsearch ...
func Linearsearch(datalist []int, key int) bool {
	for pos, item := range datalist {
		if item == key {
			fmt.Println(pos)
			return true
		}
	}
	return false
}
