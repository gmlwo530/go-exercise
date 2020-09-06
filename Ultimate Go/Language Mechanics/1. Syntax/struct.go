package main

import "fmt"

type example struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {
	var el example

	// %+v 변수는 값이 구조체인 경우 필드명도 표시합니다.
	// 구조체의 zero values는 아래와 같습니다
	fmt.Printf("%+v\n", el) // {flag:false counter:0 pi:0}

	/*
		example 구조체에는 얼마나 많은 메모리를 할당 할까요?

		bool은 1바이트, int16은 2바이트, float32는 4바이트이므로 합하면 7바이트이다.
		그러나 실제로는 8바이트이다. 이 사실은 우리가 padding과 alignment(조정) 개념을 알아야 함을 말해준다.
		padding byte는 bool과 int16 사이에 놓이게 된다. 이유는 alignment 때문이다.

		alignment의 개념 : 하드웨어가 조정 바운더리 내에서 메모리를 읽게 하는 것이 더 효율적이다.

		Rule 1:
		정확한 value의 사이즈에 의존합니다. Go가 우리가 필요한 만큼의 alingment를 결정합니다.
		모든 2 bytes value는 무조건 2 bytes boundary를 따라야 합니다.
		bool value는 1bytes이고 address 0부터 시작하게 되면, 그 다음인 int16은 무조건 address 2에서 부터 시작해야 합니다.
		그 만큼 건너 뛴 주소의 바이트는 1바이트 패딩이 됩니다. 비슷하게, 만약 4 bytes value는 3 bytes padding value를 가져야 합니다.

		Rule 2:
		제일 큰 필드가 전체 구조체의 padding을 표현합니다. 우리는 가능한 padding의 양을 최소화 해야 합니다.
		그러기 위해 항상 필드를 높은 필드부터 낮은 필드로 나열 해야 한다. 이렇게 하면 padding은 아래로 내려가게 된다(?)
		위의 example 구조체는 int64가 8 bytes이기에 전체 구조체의 크기는 8bytes이다.
	*/

	// example 타입을 선언하고 struct literal을 사용하여 변수를 초기화 할 수 있습니다
	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)

	// 익명의 타입을 선언하고 struct literal을 사용하여 변수를 초기화 할 수 있습니다
	e3 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	fmt.Println("Flag", e3.flag)
	fmt.Println("Counter", e3.counter)
	fmt.Println("Pi", e3.pi)

	/*
		Name type vs anonymous type

		만약 두 가지 이름의 타입 구조가 있다고 하면, 하나의 타입을 하나의 타입에 할당 할 수 없다.
		예를 들어, example1과 example2 타입이 있다고 하고 아래와 같이 초기화 한다고 하자.

		var ex1 example1
		var ex2 example2

		이때 ex1 = ex2는 허용 되지 않는다. ex1 = example1(ex2)와 같이 변환을 명시해야한다.

		그러나 ex라는 변수가 위의 e3처럼 익명의 구조 타입이면 ex1 = ex가 가능하다.
	*/
	var e4 example
	e4 = e3
	fmt.Printf("%+v\n", e4)
}
