package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {

	file, err := os.Create("acendente.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	file2, err := os.Create("decendente.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	var s []string
	var cosa string
	var canti int

	fmt.Println("Cantidad")
	fmt.Scanln(&canti)

	for i := 0; i < canti; i++ {
		fmt.Scanln(&cosa)
		s = (append(s, cosa))
		sort.Strings(s)

	}
	for v := range s {
		file.WriteString(s[v])

	}

	fmt.Println("Sorted :", s)
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println("Reverse sorted :", s)

	for v := range s {
		file2.WriteString(s[v])

	}

	//stat, err := file.Stat()

	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//total := stat.Size()
	//bs := make([]byte, total)

	//count, err := file.Read(bs)

	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//str := string(bs)
	//fmt.Println(str, "Bytes:", count)

	//file, err := os.Create("test.txt")

	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer file.Close()

	//file.WriteString("Hola mundo")

	//fmt.Println(strings.Contains("distribuidos", "bui"))
	//fmt.Println(strings.HasPrefix("distribuidos", "di"))
	//fmt.Println(strings.HasSuffix("distribuidos", "os"))
	//fmt.Println(strings.Join([]string{"Sistemas", "Distribuidos", "CUCEI"}, "-"))
	//fmt.Println(strings.Repeat("Distribuidos\n", 2))
	//fmt.Println(strings.Replace("aaaaaaaaabbbbb", "a", "b", 2))
	//fmt.Println(strings.Split("mi mama me mima", " "))
	//fmt.Println(strings.ToLower("MI MAMA ME MIMA"))
	//fmt.Println(strings.ToUpper("mi mama me mima"))

	//b := []byte("test")
	//fmt.Println(b)

	//s := string([]byte{'t', 'e', 's', 't'})
	//fmt.Println(s)
}
