package main

import (
	"fmt"
)

func main() {
	// Các loại toán tử
	var a bool
	a = 1 == 1
	a = 1 != 2
	a = 1 <= 2
	a = 1 < 2
	a = 1 >= 2
	a = 1 > 2
	fmt.Printf("%v, %T", a,a)

	// Integer
	
	var b int32
	fmt.Printf("%v, %T",b ,b)

	var c int8 = 10  // 1010
  	var d int8 = 3   // 0011
	fmt.Printf("%v, %T",c & d ,c & d) // -> 0010
	fmt.Printf("%v, %T",c | d ,c | d) // -> 1011
	fmt.Printf("%v, %T",c ^ d ,c ^ d) // -> 2 bit giống nhau là 0, khác nhau là 1  -> 1001
	fmt.Printf("%v, %T",c &^ d ,c &^ d) // -> 2 bit đều là 0 sẽ thành 1, trường hợp còn lại là 0  -> 1001

	// dịch bit
	var e int8 = 8  // 2^3 -> 0000 0100
	fmt.Printf("%v, %T", e << 3, e << 3)  // 2^3 * 2^3 = 2^6 -> 0010 0000 -> 64
	fmt.Printf("%v, %T", e >> 3, e >> 3)  // 2^3 / 2^3 = 2^0 = 1 -> 0000 0001 -> 1 
	

	// Float
	var a1 float32 = 3.14
	// a1 = 13.7e72
	// a1 = 2.1E14
	var b1 float32 = 7.21

	fmt.Printf("%v, %T", a1+b1, a1+b1)

	// Complex
	var a2 complex64 = 1 + 2i
	fmt.Printf("%v, %T", a2, a2)
	fmt.Printf("%v, %T", real(a2), real(a2))
	fmt.Printf("%v, %T", imag(a2), imag(a2))

	// String
	var a3 string = "Hello"
	fmt.Println(a3)
	var b3 []byte = []byte(a3)
	fmt.Println(b3)
	var c3 byte  = 'H'
	fmt.Printf("%v, %T", c3, c3)
	var d3 rune  = 'H'
	fmt.Printf("%v, %T", d3, d3)
}



/*
Primitive Data type // các kiểu dữ liệu nguyên thủy được ngôn ngữ lập trình cung cấp thẳng cho chúng ta sử dụng. là nguyên tố để xây dựng các kiểu dữ liệu khác
1. Boolean
2. Numerics
	Integer
		signed  integer         int8   int16    int32    int64
		unsigned integer		uint8  uint16   uint32   uint64
	Float
		float32
		float64
	Complex
		complex64
		complex128
3. Text
	String
	Byte -> UTF-8
	Rune -< UTF-32
*/