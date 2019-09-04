// 모든 Go 프로그램은 패키지로 구성되어 있습니다.
// main package에서 부터 실행을 시작합니다.
package main

/*
패키지 이름은 디렉토리 경로의 마지막 이름을 사용하는 것이 규칙입니다.
예를 들어 "path/filepath" 를 사용한다면 패키지명은 filepath 입니다.
*/

import (
	"fmt"
	"math"
) // 여러개의 패키지를 소괄호로 감싸서 표현합니다. 아래와 같은 의미입니다.
/*
	import "fmt"
	import "math"
*/

func main() {
	/*
		패키지를 import하면 패키지에서 export한 것들(메서드, 상수, 변수 등)에 접근 할 수 있습니다.
		Go에서는 첫 문자가 대문자로 시작하면 그 패키지를 사용하는 곳에서 접근할 수 있는 exported name이 됩니다.
		Go는 변수명에서 제일 앞 글자가 대문자인지 소문자인지에 따라 public과 private이 결정 됩니다.
		Foo는 public, foo는 private입니다.
	*/
	fmt.Println("Happy", math.Pi, "Day")
}
