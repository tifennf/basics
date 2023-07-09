package main

import "fmt"
import "github.com/basics/go/lib"

func main() {

	arr := Generate_rand_arr(10)
	fmt.Println(arr)
	fmt.Println(Min(arr, len(arr)))
	fmt.Println(Max(arr, len(arr)))
	fmt.Println(Average(arr, len(arr)))

}
