package main

import "fmt"

type CType interface {
	int | int32 | uint | uint32
}

func sum[T CType](a, b T) T {
	return a + b
}

func sumMany[T CType](a ...T) T {
	var s T
	for _, v := range a {
		s += v
	}
	return s
}

func main() {
	var a int = 1
	var b int = 2
	fmt.Println(sum(a, b))

	c := []int{1, 2, 3}
	d := []int32{4, 5, 6}
	fmt.Println(sumMany(c...))
	fmt.Println(sumMany(d...))
}
