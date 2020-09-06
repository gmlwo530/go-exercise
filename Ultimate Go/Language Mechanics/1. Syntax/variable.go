package main

import "fmt"

func main() {
	/*
		### Built-in types

		타입은 아래의 질문을 통해 가독성과 무결성을 제공한다.
		- 할당하는 메모리의 양은 얼마인가?(ex. 32-bit, 64-bit)
		- memory represent란 무엇인가?(ex. int, uint, bool)

		타입은 int32 또는 int64와 같이 특정 될 수 있다.
		- uint8은 1 바이트의 메모리를 사용해 10진수를 가진다.
		- int32은 4 바이트의 메모리를 사용해 10진수를 가진다.

		만약 uint 또는 int와 같이 특정하지 않으면 코드가 빌드 되는 아키텍처에 따라 결정 된다.
		64-bit OS에서는 int는 int64로 매핑 된다.


		### Word size

		word size는 word의 바이트 수다. 예를 들어 64-bit architecture,
		일 때 word size는 64 bit (8bytes)이다.
		주소 사이즈가 64 bit이므로 우리의 정수는 64 bit이다.


		### Zero value concept

		모든 변수는 초기화가 되어야 한다. 만약 특정 값을 주지 않으면, zero value로 할당된다.
		전체 메모리 할당(타입에 의해 우선적으로 메모리 할당이 되는 것으로 추측)에서 해당 비트를 0으로 재설정합니다.

		- Boolean : false
		- Integer : 0
		- Floating Point : 0
		- Complex : 0i
		- String : ""
		- Pointer : nil


		### Declare and initialize

		var는 타입의 zero value로 초기화를 보장해주는 유일한 선언 방식이다.

		String(문자열)은 일련의 uint8 타입입니다.

		string은 두 단어로 된 데이터 구조입니다: 첫 번째 단어는 backing array를 향하는 포인터를 표현하고,
		두 번째 단어는 길이를 표현합니다. 만약 zero value라면, 첫 번째는 nil 두번째는 0을 나타냅니다.
	*/

	var a int
	var b string
	var c float64
	var d bool

	// %T는 값의 타입을 %v는 point 구조체의 인스턴스를 표시합니다.
	fmt.Printf("var a int \t %T [%v]\n", a, a)     // var a int int [0]
	fmt.Printf("var b string \t %T [%v]\n", b, b)  // var b string string []
	fmt.Printf("var c float64 \t %T [%v]\n", c, c) // var c float64 float64 [0]
	fmt.Printf("var d bool \t %T [%v]\n", d, d)    // var d bool bool [false]

	// short variable 선언식을 사용하면 정의와 초기화를 동시에 할 수 있습니다.
	aa := 10
	bb := "hello"
	cc := 3.14159
	dd := true

	fmt.Printf("aa := 10 \t %T [%v]\n", aa, aa)
	fmt.Printf("bb := \"hello\" \t %T [%v]\n", bb, bb)
	fmt.Printf("cc := 3.14159 \t %T [%v]\n", cc, cc)
	fmt.Printf("dd := true \t %T [%v]\n", dd, dd)

	// ### Conversion vs casting
	// Go에는 casting이 없지만, conversion(변환)이 있습니다.
	// 컴파일러에게 메모리를 더 많이 가진척 하라고 하는 것 대신에, 우리가 더 많은 메모리를 할당 해줄 수 있습니다
	aaa := int32(10)
	fmt.Printf("aaa := int32(10) %T [%v]\n", aaa, aaa) // aaa := int32(10) int32 [10]
}
