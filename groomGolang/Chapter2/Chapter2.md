# Chapter2

## 콘솔 출력 함수

### 기본 출력 함수
 - print    :개행 없이 출력
 - println  :출력후 개행문자 
 - 여러 데이터 출력시 ,(콤마) 사용
 - 연산식을 결과 값으로 출력

 ~~~go
package main

func main() {
	var num1 int = 1
	var num2 int = 2
	
	print("Hello goorm!")
	print(num1)
	print(num2)
	print(num1 + num2)
	print("Hello goorm!", num1 + num2,"\n")
	
	println("Hello goorm!")
	println(num1)
	println(num2)
	println(num1 + num2)
	println("Hello goorm!", num1 + num2)	
}
 ~~~
출력 결과
 ~~~
Hello goorm!123Hello goorm!3
Hello goorm!
1
2
3
Hello goorm! 3
 ~~~


 ### fmt
package main 밑 줄에 import "fmt" 패키지 사용 선언
- fmt.Print()   
- fmt.Println()
- fmt.Printf()  : 함수 서식 문자를 채워 원하는  포맷으로 출력

~~~go
package main

import "fmt"

func main() {
    var num1 int = 1
    var num2 int = 2
    
    fmt.Print("Hello goorm!", num1, num2, "\n")
    
    fmt.Println("Hello goorm!", num1, num2)
	
    fmt.Printf("num1의 값은:%d num2의 값은:%d\n", num1, num2)
}
~~~
결과
~~~
Hello goorm!1 2
Hello goorm! 1 2
num1의 값은:1 num2의 값은:2
~~~


## 	변수
- 어떤 형의 값을 저장할 공간을 선언하는 것

선언 방식
- var <변수이름> <변수형>
- Short Assignment Statement  :=
	- 형 선언 없이 타입 추론이 가능
	- 함수 내에서만 사용 가능
- 전역 변수는 var 키워드 반드시 사용


 변수를 선언하고 초기값을 설정하지 않으면 Zero value로 설정 

 |type| zero value|
 |----|---------|
 |bool|false|
 |숫자|0|
 |string|""(빈문자열)|



`변수를 선언하고 사용하지 않으면 오류 발생`

~~~go
package main

import "fmt"

var globalA = 5 //함수 밖에서는 'var' 키워드를 입력해야함.
				// 꼭 형을 명시하지 않아도 됨
func main() {
    var a string = "goorm"
    fmt.Println(a)

    var b int = 10
    fmt.Println(b)

    var d = true
    fmt.Println(d)

    var e int
    fmt.Println(e)

    f := "short"
    fmt.Println(f)
	
	fmt.Println(globalA)
}
~~~

출력 결과
~~~
goorm
10
true
0
short
5
~~~

동시에 여러개 초기화 가능.
 - 초기화 하는 값의 개수가 동일해야함
 - 초기화 하지 않는다면 모두 안해야함
 ~~~go
 package main

import "fmt"

func main() {
    var a, b int = 10, 20
    fmt.Println(a, b)

	i, j, k := 1, 2, 3
    fmt.Println(i, j, k)

    var str1, str2 string = "Hello", "goorm"
    fmt.Println(str1, str2)
}
~~~


## 상수
선언 초기화 후 변경 불가<br>
- 선언 형태 - const 상수이름 상수형 ( 상수형은 생략가능 )
- 선언후 초기화 반드시 필요 - 선언만 할시 오류
- 초기화 후 사용하지 않아도 오류 X - 명시자체로 의미( 변수와 다른 점)
- const 키워드 생략 불가 따라서 := 사용 불가

~~~go
package main

import "fmt"

const username = "kim"

func main() {
	const a int = 1    
    const b, d= 10, 20 //여러 값을 초기화할 수 있음
	const c = "goorm"
	
	
	fmt.Println(username)
}
~~~

상수만의 초기화 방법

- 괄호로 같이 묶여있는 상수들은 다른 형으로 초기화 가능
- 괄호 시작(과 괄호 마지막)의 위치는 상관 없지만 각 상수들은 개행하여 초기화 -> 개행하지 않고 초기화하면 에러가 발생합니다.
- 각 상수들 사이에 콤마(,)를 입력하면 안됨 -> 입력하면 에러가 발생
- 묶어서 선언된 상수들 중에서 첫번째 값은 꼭 선언되어야 함 선언되지 않은 값은 바로 전 상수의 값을 가짐
- iota라는 식별자를 값으로 초기화하면 그 후에 초기화하지 않고 이어지는 상수들은 순서(index)가 값으로 저장

~~~go
package main

import "fmt"

const ( 
	c1 = 10 //첫 번째 값은 무조건 초기화해야 함
	c2
	c3
	c4 = "goorm" //다른 형 선언 가능
	c5
	c6 = iota //c8까지 index값 저장
	c7
	c8
	c9 = "earth"
	c10
	c11 = "End"
)

func main() {
	fmt.Println(c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11)
}
~~~
출력 결과
~~~
10 10 10 goorm goorm 5 6 7 earth earth End
~~~

## 실습 코드
~~~go
package main

import "fmt"

func main() {
	
	var num1 int = 3;
	num2 := 7;
	
	result := num1 + num2;
	
	fmt.Printf("%d과 %d의 합은 %d입니다.", num1 , num2 , result)
}
~~~


~~~go
package main

import "fmt"

const (
	name = "kim"
	RRN = "800101-1000000"
	job
)

func main() {
	fmt.Println( name , RRN, job)
}
~~~