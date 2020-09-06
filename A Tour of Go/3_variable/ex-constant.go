package main

import (
	"fmt"
)

const Pi = 3.14

// 숫자형 상수는 정밀한 값(values) 을 표현할 수 있습니다. 타입을 지정하지 않은 상수는 문맥(context)에 따라 타입을 가지게 됩니다.

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 { return x * 0.1 }

// 상수는 문자(character), 문자열(string), 부울(boolean), 숫자 타입 중의 하나가 될 수 있습니다.
func main() {
	fmt.Println(Pi)

	const hello = "안녕"

	fmt.Println(hello, ", world!")

	const foo = true

	fmt.Println(foo)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
