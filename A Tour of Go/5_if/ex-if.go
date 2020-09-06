package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x)) // Sprint는 값을 그대로 string으로 만듭니다. 개행을 하려면 Sprintln을 사용하면 됩니다.
}

func pow(x, n, lim float64) float64 { // 조건문을 다음과 같이 짧게 쓸 수도 있습니다.
	if v := math.Pow(x, n); v < lim { // 이때, 짧은 실행문을 통해 선언된 변수는 if 안쪽 범위(scope) 에서 만 사용할 수 있습니다.
		return v
	}
	return lim
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
