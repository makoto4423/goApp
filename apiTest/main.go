package main

import "fmt"

func main() {
	n := 3
	arr := make([][]int, n, 2)
	fmt.Print(arr)
}

//func main() {
//	sample1 := []string{"a", "b"}
//	test(&sample1)
//	fmt.Printf("Sample1 After Test Function Call: %v\n", len(sample1))
//	fmt.Printf("Sample1 After Test Function Call: %v\n", sample1)
//}
//func test(sample *[]string) {
//	fmt.Printf("ptr: %p\n", &sample)
//	*sample = (*sample)[:1]
//	(*sample)[0] = "c"
//	fmt.Printf("ptr: %p\n", &sample)
//	fmt.Printf("Sample in Test function: %v\n", sample)
//}
