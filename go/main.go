package main

import "fmt"
import "basics/lib"

func main() {

	arr := lib.GenerateRandArr(10)
	fmt.Println(arr)
	fmt.Println(lib.Min(arr, len(arr)))
	fmt.Println(lib.Max(arr, len(arr)))
	fmt.Println(lib.Average(arr, len(arr)))

	lib.Sort(arr, len(arr))
	fmt.Println(arr)

}
