package main

import (
	"fmt"
	"sort"
)

func main() {

	s := []string{"Hola", "Como", "Estas", "Alberto"}
	fmt.Println("Unsorted:", s)
	sort.Strings(s)
	fmt.Println("Sorted : ", s)

	s1 := []int{1, 2, 3, 4}
	fmt.Println("Unsorted:", s1)
	sort.Sort(sort.Reverse(sort.IntSlice(s1)))
	fmt.Println("Sorted : ", s1)

}
