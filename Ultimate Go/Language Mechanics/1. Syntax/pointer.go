package main

func increment1(inc int) {
	// Increment the "value of" inc.
	inc++
	println("inc1:\tValue Of[", inc, "]\tAddr Of[", &inc, "]")
}

/*
	increment2은 count를 값이 address고 정수의 타입을 가르키는 것을 나타내는 pointer 변수로 선언한다.
	여기서 * 는 operator가 아니다. 타입 이름의 한 부분이다.
	모든 타입은 선언이 되었든, 선언을 하든, 선언 전이든 포인터에 대한 자유를 가진다.(?)
*/
func increment2(inc *int) {
	// Increment the "value of" count that the "pointer points to"
	// 여기서 * 는 operator이다. 이것은 포인터가 가르키는 부분의 값을 말해준다.
	*inc++
	println("inc2:\tValue Of[", inc, "]\tAddr Of[", &inc, "]")
}

func main() {
	/*
		### Everything is about pass by value

		Pointers의 목적은 단지 하나다: sharing. Pointers는 program boundaries들을 걸치며 값을 공유한다.
		Program boundaries에는 몇몇 종류가 있다. 대표적으로는 function call 있고, Gorutines 같은 것들이 있다.

		프로그램이 시작 할 때, 런타임은 Goroutine을 생성한다. 모든 Goroutine은 기계로부터 실행 되어야 할 지시들을 가지고 있는 분리된 path이다.
		또한 Goroutines은 가벼운 쓰레드라고 생각 할 수도 있다.

		모든 Goroutines은 stack이라고 불리는 메모리 블럭을 제공 받는다. Go에서 stack memory는 매우 작은 2k부터 시작한다.
		이 수치는 시간에 따라 바뀔 수 있다. 함수가 호출되는 매순간, 스택의 한 부분은 함수가 실행되게끔 사용 된다. The growing direction of the stack is downward.

		모든 functions에는 stack frame이 주어지고, 함수의 메모리 실행에 사용된다. 모든 stack frame의 사이즈는 compile 시간 때 정해진다.
		컴파일러가 미리 크기를 알지 않는 이상, 스택에 값을 넣을 수 없습니다. 만약 우리가 compile time에 어떤 것에 대한 사이즈를 알지 못한다면, 그 어떤 것은 heap에 올라가게 됩니다.

		Zero value는 우리가 모든 stack frame을 초기화 할 수 있게 해줍니다. Stacks are
		self cleaning. We clean our stack on the way down. Every time we make a function,
		zero value initialization cleans the stack frame. We leave that memory on the way
		up because we don't know if we would need that again.
	*/

	// ### Pass by value
	count := 10

	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	// Pass the "value of" count
	increment1(count)
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	// Pass the "address of" count. This is still considered pass by value,
	// not by reference because the address itself is a value.
	increment2(&count)
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")
}
