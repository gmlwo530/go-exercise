package main

import (
	"fmt"
)

var x, y, z int              // 초기값 0
var clang, python, java bool // 초기값 false

// 초기화를 하는 경우 타입(type)을 생략할 수 있습니다. 변수는 초기화 하고자 하는 값에 따라 타입이 결정됩니다.
var a, b, c int = 1, 2, 3
var ruby, dart, elixr = true, false, "no"

func main() {
	short_val := "Short Value!" // 함수 내에서 :=을 사용하면 var과 명시적 타입 선언을 생략 할 수 있다. 대신 함수 밖에서는 사용 할 수 없다.

	fmt.Println(x, y, z, clang, python, java)
	fmt.Println()
	fmt.Println(a, b, c, ruby, dart, elixr)
	fmt.Println(short_val)
}
