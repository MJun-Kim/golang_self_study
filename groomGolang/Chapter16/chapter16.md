# 에러처리

- 에러 처리를 하는 이유는 컴파일러가 알아차리지 못하는 프로그램상의 오류를 예방하기 위함

*fmt 함수들은 에러를 같이 반환함!
ex ) 
- func Scan(a ...interface{}) (n int, err error)
- func Scanln(a ...interface{}) (n int, err error)
...

자세한것은 https://golang.org/pkg/ 에서 확인하자!

위의 반환된 error 값으로 에러 처리!


## error 타입

- 인터페이스형
- string을 반환하는 Error() 메소드 한개만 가지고 있음
~~~go
type error interface {
	Error() string
}
~~~

Error 메서드 구조체
~~~go
func (e *errorString) Error() string {
	return e.s
}
~~~

- errorString의 포인터 초기화 -> errors.New()함수 사용

ex)

~~~go
package main

import (
	"fmt"
	"errors"
)

func divide(a float32, b float32) (result float32, err error) {
	if b == 0 {
		return 0, errors.New("0으로 나누지마") 
	}
	result = a / b
	return 
}

func main() {
	var num1, num2 float32
	fmt.Scanln(&num1, &num2)
	
	result, err := divide(num1, num2)
	
	if err != nil {
		panic(err)
	}
	
	fmt.Println(result)
}
~~~

---

## 에러 출력 및 처리 

panic()을 사용하지 않는 다른 방법
-> log 패키지를 사용!

log 패키지의 함수들

- func Fatal(v ...interface{}) : 에러 로그 출력 및 프로그램 종료
- func Panic(v ...interface{}) : 시간, 에러 메시지 출력 및 패닉 발생, defer 구문이 없을 시 런타임 패닉을 발생시키고 콜스택 출력
- func Print(v ...interface{}) : 시간, 에러 메시지 출력 하지만 프로그램 종료하지 않음

ex ) 
~~~go
package main

import (
	"fmt"
	"log"
)

func errorChek(n int) (string, error) {
	if n == 1 {
		s := "Goorm"
		return s, nil // 정상 동작이므로 에러 값은 nil
	}

	return "", fmt.Errorf("%d는 1이 아닙니다.", n) // 1이 아닐 때는 에러 리턴
}

func main() {
	s, err := errorChek(1) // 매개변수에 1을 넣었으므로 정상 동작
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
	fmt.Println("정상 작동")

	s, err = errorChek(2) 
	if err != nil {
		log.Print(err)
	}
	fmt.Println("에러가 발생하지만 코드 속행")

	defer func() {
		s, err = errorChek(4)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("에러를 발생하고 프로그램을 완전 종료 -> 이 부분 출력 X")
	}()

	s, err = errorChek(3) 
	if err != nil {
		log.Panic(err) // defer 함수로 이동
	}
	// 에러가 발생하여 프로그램이 종료되었으므로 이 아래부터는 실행되지 않음
	fmt.Println("이부분도 panic함수로 인해 실행되지 않음 -> defer 부분 실행")

	fmt.Println(s)

	fmt.Println("Hello, world!")
}
~~~
출력
~~~go
정상 작동
2021/05/17 10:33:03 2는 1이 아닙니다.

에러가 발생하지만 코드 속행
2021/05/17 10:33:03 3는 1이 아닙니다.
2021/05/17 10:33:03 4는 1이 아닙니다.
~~~

--- 

## 예제 문제 

- 평균점수 3

~~~go
package main

import (
	"fmt"
)

func inputSubNum() (num int, err error) {
	fmt.Scanln(&num)
	if num <= 0 {
		return 0, fmt.Errorf("잘못된 과목 수입니다.")
	}
	return
}

func average(num int) (avg float64, err error) {
	var score, total int
	for i := 0; i < num; i++ {
		fmt.Scanln(&score)
		if score < 0 {
			return 0, fmt.Errorf("잘못된 점수입니다.")
		}
		total += score
	}
	avg = float64(total) / float64(num)
	return

}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	subjectCount, err := inputSubNum()
	if err != nil {
		panic(err.Error())
	}

	result, err := average(subjectCount)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(result)
}
~~~