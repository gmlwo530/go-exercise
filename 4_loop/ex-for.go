package main

import (
	"fmt"
)

// Go Lang에서는 반복문이 for 밖에 없습니다. 특이한 점은 ()가 필요하지 않습니다.
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 다음과 같이 전.후 처리를 제외하고 조건문만 표현 할 수 있습니다. 이 표현은 while문을 표현 하고 있다고 할 수 있습니다.
	sum2 := 1

	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	/*
		다음과 같이 무한 루프를 만들 수 있습니다
		for {
		}
	*/
}
