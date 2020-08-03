package main

import "fmt"

func main() {
	var (
		name string
	)
	name = "Andi"

	fmt.Println("nama saya", name)

	var n1 float32 = 4.567
	var n2 int8 = 50
	var s1 string = "Hi"
	var b1 bool = true

	fmt.Println(n1, n2, s1, b1)

	var a int8 = 1

	if a != 5 {
		fmt.Println("tidak sama dengan 5 cuy")
	}

	switch a {
	case 1:
		fmt.Println("satu")
	case 2:
		fmt.Println("dua")
	case 3:
		fmt.Println("tiga")
	default:
		fmt.Println("ya default cuy")
	}

	var arrayTest [3]int
	arrayTest[1] = 1
	fmt.Println(arrayTest)

	var arrayMulti [2][3]int
	arrayMulti[0][0] = 122
	arrayMulti[1][2] = 322
	arrayMulti[0][2] = 11

	fmt.Println(arrayMulti)

	lanjut := 1
	for lanjut <= 10 {
		fmt.Println(lanjut)
		lanjut++
	}

	slice1 := []int{2, 3, 4}
	slice2 := []byte("Hello")
	fmt.Println(slice1, slice2)

	var slice3 = make([]int, 3, 5)
	fmt.Println(slice3)
	fmt.Println("length", len(slice3))
	fmt.Println("cap", cap(slice3))

	slice3 = append(slice3, 1, 2, 3, 4)
	fmt.Println(slice3)
	fmt.Println("length", len(slice3))
	fmt.Println("cap", cap(slice3))
}
