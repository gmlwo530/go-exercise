package main

import (
	"fmt"
)

// 매개변수의 타입은 변수명 *뒤에* 명시합니다.
// 타입을 뒤에 명시하는 이유를 간단하게 설명하면, 코드를 왼쪽에서 오른쪽으로 읽을 때 자연스럽게 읽기 위해서 입니다. 자세한 설명은 아래 사이트를 참조하면 됩니다.
// https://blog.golang.org/gos-declaration-syntax

func add(x int, y int) int {
	return x + y
}

/*
	이런 식으로 타입 명시를 할 수도 있습니다.
	func add(x, y int) ...
*/

// 여러 개의 결과를 반환 할 수 있습니다.
func swap(x, y string) (string, string) {
	return y, x
}

/*
	반환 값에 이름을 부여하면 변수처럼 사용할 수도 있습니다.
	결과에 이름을 붙히면, 반환 값을 지정하지 않은 return 문장으로 결과의 현재 값을 알아서 반환합니다.
*/
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("function example")
	fmt.Println("Add Function :", add(1, 2))

	a, b := swap("Hello", "World")
	fmt.Println("Swap Function :", a, b)

	fmt.Print("Split Function : ")
	fmt.Println(split(17))
}
