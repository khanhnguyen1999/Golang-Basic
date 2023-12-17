package main

import "fmt"
import "strconv"

// global scope , chỉ được khai báo theo kiểu này
var k string = "Hello"
var (
	n int = 10
	m int = 20
	h string = "ABC"
)

func main(){
	// 1. Định nghĩa biến trong ngôn ngữ lập trình
	// khai báo biến
	// 2. Cú pháp khai báo biến
	var i int = 10
	j := 20.5
	fmt.Println(i)
	fmt.Println(j)
	// %v print giá trị, %T print kiểu dữ liệu
	fmt.Printf("%v, %T",j, j)
	fmt.Println()

	// 3. Global and block scope
	fmt.Println(k)
	fmt.Println(n)
	fmt.Println(m)
	fmt.Println(h)


	// 4. Shadow
	var n int = 10000
	fmt.Println(n)

	// 5. when var declared, it must be use
	// 6. Export global scope -> Viết hoa chữ cái đầu ở global scope
	// 7. Naming convention -> camel
	var helloString string = "Hello"
	fmt.Println(helloString)
	// 8. Convert type

	var i1 int = 10
	fmt.Printf("%v, %T",i1, i1)
	fmt.Println()
	var j1 float32 = float32(i1)
	fmt.Printf("%v, %T",j1, j1)

	// parse float to string use strconv
	fmt.Println()
	var k1 string = strconv.Itoa(10)
	fmt.Printf("%v, %T",k1, k1)

}


/*
1. Định nghĩa biến trong ngôn ngữ lập trình
2. Cú pháp khai báo biến
3. Global and block scope
4. Shadow
5. Declared and block scope
6. Export global scope
7. Naming convention -> camel
8. Convert type
*/