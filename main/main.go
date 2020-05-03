package main

import (
	"fmt"

	"github.com/mradrianhh/go-algorithms/search"
)

func main() {
	items := []int{12, 23, 25, 12, 15, 35, 135, 235, 2263, 14}
	fmt.Println(search.Linearsearch(items, 23))
}
