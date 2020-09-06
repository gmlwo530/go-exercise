package main

type user struct {
	name  string
	email string
}

func main() {
	/*

	 */

	stayOnStack()
	escapeToHeap()
}

/*
	stackOnStack은 어떻게 변수가 escape하지 않는 것을 보여준다. 컴파일 타임 때 user 변수의 사이즈를 알기 때문에,
	컴파일러는 user 변수를 stack frame에 다시 집어 넣는다.
*/
func stayOnStack() user {
	// Stack frame 내부에서 값을 만들고 초기화 한다.
	u := user{
		name:  "Hee Jae",
		email: "gmlwo530@gmail.com",
	}

	// 값을 리턴하고 main stack frame으로 다시 전달 합니다.
	return u
}

/*
	escapeToHeap은 어떻게 변수가 escape 하는지 보여준다. 이것은 stayOnStack 함수랑 거의 같아 보인다.
	stayOnStack() 함수와 앞 부분은 같지만 user 변수 자체의 값을 return하는게 아니고 user 변수의 주소를 리턴한다.
	이것은 값이 call stack으로 되돌아 가게 될것을 말한다.(?) 우리는 포인터 의미론을 사용하는 것이다.

	escapeToHeap을 호출하고 나서 어떤 것들을 가지고 있는지 생각 해볼 필요가 있다.
	main은 값을 가르키는 포인터를 stack frame 아래에 가지고 있다. 이 상황에서 문제가 생기게 된다. call stack으로 되돌아 올때, 메모리에 존재하고 재사용 가능하다(?).
	main가 언제든 함수를 호출 할 때, 우리는 frame을 할당하고 초기화 해야한다(?)

	여기서 잠시 zero value에 대해 생각해보자. 이것은 모든 stack frame을 초기화 할 수 있게 해준다. Stack은 자체적으로 지워진다.
	We clean our stack on the way down. Every time we make a function call, zero value, initialization, we are cleaning
	those stack frames. We leave that memory on the way up because we don't know if we
	need that again.

	다시 예제로 돌아오자. user의 address를 반환하고 user valu를 지워질 call stack으로 되돌려 보내는 것이 않 좋게 보인다.
	하지만 이런 일은 발생하지 않는다.

	실제로 발생하는 것은 escape analysis이다. "return &u" 라인 때문에, 힙 위에 놓게 되고 이 값은 '함수를 위한 stack frame'의 안에 넣어질 수 가 없다.

	Escape analysis는 어떤 것이 stack에 남고 남지 않을 것인지 결정한다. stayOnStack 함수에서, value의 복사본을 반환 했기 때문에 변수는 stack에 안전하게 보관된다.
	하지만 call stack 위에서 무언가를 공유 할 때, escape analysis는 우리가 main으로 돌아왔을 때 이 메모리가 더 이상 유효하지 않다고 판단하고 heap에 놓게 된다.
	main will end up having a pointer to the heap.
	사실상, 이 할당은 heap에서 이루어 진다. escapeToHeap은 힙에 대한 포인터를 가지게 된다. 하지만 u 변수는 value semantics에 기반 된다.
*/
func escapeToHeap() *user {
	u := user{
		name:  "Hee Jae",
		email: "gmlwo530@gmail.com",
	}

	return &u
}

/*
	What if we run out of stack space? (스택 공간이 부족하면 어떻게 될까?)

	함수 호출 후 일어나는 것은, 이 frame에 충분한 stack space가 있는가를 물어보는 것입니다.
	만약 있으면 컴파일 타임 때 모든 frame의 사이즈를 아니깐 아무런 문제가 없습니다.
	만약 아니라면 더욱더 큰 frame을 가져야 하고 이 값은 복사되어야 합니다. 메모리는 stack으로 이동합니다.
	이것은 trade off 이빈다. 이 현상은 많이 발생하지 않기에 우리는 이 cost를 가져야 합니다. The benefit of using less memory in any Goroutine outweighs the
	cost.

	stack을 증가 할 수 있기에, 어떤 Goroutine도 다른 Goroutine stack에 대한 포인터를 가지고 있을 수 없습니다.
	모든 포인터들에 대한 과정을 컴파일러가 유지한다면 미친 latency가 발생하게 될 것입니다.

	Goroutine을 위한 스택은 오직 그 Goroutine에 대한 스택입니다. 다른 Goroutine과 공유 될 수 없습니다.
*/

// ------------------------------------------------------------ //

/*
	Garbage collection

	무언가 heap으로 옮겨질 때, Garbage Collection은 그것을 들고 옵니다. Garbage Collector의 중요한 점은 pacing algorithm이라는 것입니다.
	이것은 GC가 실행되기 위한 가장 짧은 t를 유지하기 위한 frequency/pace를 결정합니다.

	프로그램이 4MB heap을 가졌다고 해봅시다. GC는 live heap을 2MB로 유지합니다.
	만약 live heap이 4MB를 넘으면, 더 큰 heap을 할당 합니다. GC의 pace는 heap size가 얼마나 빠르게 성장하는지에 의존합니다.
	The slower the pace, the less impact it is going to have. The goal is to get the live heap back
	down.

	GC가 실행 될 때, we have to take a performance cost so all Goroutines can keep running concurrently.
	GC는 또한 garbage collection을 하기 위한 Goroutine 그룹을 가지고 있습니다. 이것은 사용 할 수 있는 CPU 용량의 25%를 사용합니다.
*/
