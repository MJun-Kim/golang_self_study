# Chapter 10 컬랙션

- 컬랙션 : 2개 이상의 변수를 모아놓은것
- Go에서는 3가지
  - Array 배열
  - Slice 슬라이스
  - Map 맵

---

## 배열 Array

- 정적이다! -> 배열의 크기가 고정적
- 부분 배열 발췌, 크기 동적 증가 x
- 배열의 크기는 자료형의 한 요소!
- 다차원 배열 가능

배열 선언 방법들~

```Go
package main

import (
    "fmt"
	"reflect"
)

func main() {
	var arr1 [5]int                     //길이가 5인 int형 배열 arr1을 선언
	arr1 = [5]int{1, 2, 3, 4, 5}        //배열 초기화
	arr2 := [4]int{4, 5, 6, 7}          //:= 를 이용해 선언
	var arr3 = [...]int{9, 8, 7, 6}     //[...]을 이용한 배열 크기 자동 설정

    arr4 := [2][3]int{ { 1, 2 ,3} , { 2,6,9}}   //다차원 배열

	fmt.Printf("arr1:%T\narr2:%T\narr3:%T\n%T\n", arr1, arr2, arr3,arr4)
	fmt.Printf("arr1 타입 == arr2 타입 : %t\n", reflect.TypeOf(arr1) == reflect.TypeOf(arr2))
	fmt.Printf("arr1 타입 == testArr1 타입 : %t\n", reflect.TypeOf(arr2) == reflect.TypeOf(arr3))
}
```

결과출력

```
arr1:[5]int
arr2:[4]int
arr3:[4]int
[2][3]int
arr1 타입 == arr2 타입 : false
arr2 타입 == arr3 타입 : true
```

---

## Slice 슬라이스

- 배열과 비슷하지만 다른
- 동적으로 크기 조절가능
- 부분 발췌 가능
- 선언만 했을시 포인터까지만생성

```Go
package main

import "fmt"

func main() {
	var a []int        //슬라이스 변수 선언 아무것도 초기화 되지 않은 상태
	a = []int{1, 2, 3} //슬라이스에 리터럴값 지정

	a[1] = 10           //값이 할당되어 메모리가 생겼기 때문에 이렇게 접근 가능

	fmt.Println(a)

	var b []int         //nil slice 선언

	if b == nil {
		fmt.Println("용량이", cap(b), "길이가", len(b), " Nil Slice입니다.")
	}
}
```

### make 함수 사용

- Slice를 선언만 하면서 초기화 하는 방법! -> make 함수를 사용하자!

```
make(슬라이스 타입, 슬라이스 길이, 슬라이스의 용량) //용량 생략시 용량 == 길이
```

```Go
package main

import "fmt"

func main() {
	s := make([]int, 0, 3) // len=0, cap=3 인 슬라이스 선언
	fmt.Println(s)
}
```

---

### 슬라이스 추가, 병합, 복사

- 추가 또는 병합 : append()

```GO
package main

import "fmt"

func main() {
    sliceA := []int{1, 2, 3}
    sliceB := []int{4, 5, 6}

    sliceA = append(sliceA, sliceB...)
    sliceA = append( sliceA, 23)

    fmt.Println(sliceA) // [1 2 3 4 5 6] 출력
}
```

- 복사 : copy or [:]
- 슬라이스[처음인덱스 : 끝인덱스(끝은 포함하지 않음)]

```Go
package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}               //
	sliceCopy := slice[3:8]                              //얕은 복사?
	sliceDeepCopy := make([]int, len(slice), cap(slice)) //깊은복사

	array := [4]int{1, 2, 3, 4}
	arrayCopy := array
	// var sliceDeepCopy []int로 포인터만 생성시에는 메모리를 할당 받지 않아서? 복사가 되지 않음. (에러는 나지 않는다..)
	// 복사하려는 슬라이스의 용량이 더 크면 에러 발생!

	copy(sliceDeepCopy, slice)

	fmt.Println("변경전")
	fmt.Println("-슬라이스-")
	fmt.Println(slice)
	fmt.Println(sliceCopy)
	fmt.Println(sliceDeepCopy)
	fmt.Println("-배열-")
	fmt.Println(array)
	fmt.Println(arrayCopy)
	fmt.Println()

	slice[2] += 3
	slice[3] += 9

	array[2] += 3
	array[3] += 9

	fmt.Println("변경후")
	fmt.Println("-슬라이스-")
	fmt.Println(slice)
	fmt.Println(sliceCopy)
	fmt.Println(sliceDeepCopy)
	fmt.Println("-배열-")
	fmt.Println(array)
	fmt.Println(arrayCopy)
}


```

출력

```Go
변경전

-슬라이스-
[1 2 3 4 5 6 7 8]
[4 5 6 7 8]
[1 2 3 4 5 6 7 8]

-배열-
[1 2 3 4]
[1 2 3 4]

변경후

-슬라이스-
[1 2 6 13 5 6 7 8]
[13 5 6 7 8]
[1 2 3 4 5 6 7 8]

-배열-
[1 2 6 13]
[1 2 3 4]
```

- 배열은 값을 복사해오고 슬라이스는 같은 주소값은 참조한다!

---

## Map 맵

- key , value 값으로 이루어진 값을 저장하는 컬렉션
- 슬라이스와 같은 값을 참조해서 쓰는 참조타입!
- 'var 맵이름 map[key자료형]value자료형' 로 선언해서 사용.
- 초기화 후에 ( make() 함수 or {} ) 값을 추가, 갱신, 삭제 가능

```go
package main

import "fmt"

func main() {
	var a map[int]string

	if a == nil {
		fmt.Println("nil map")
	}

	var m = map[string]string{ //key:value, 형식으로 초기화한다
		"apple":  "red",
		"grape":  "purple",
		"banana": "yellow",
	}

	fmt.Println(m, "\nm의 길이는", len(m))
}
```

## 값 읽기

- 콘솔 출력 함수에 "맵이름[key]"을 바로 입력할 때는 key 값에 해당되는 value 값만 출력
- value 값과 true/false 값을 반환받기 위해서는 변수 두 개를 선언.<br>
  ex) val, exist := 맵이름[key] -> val에는 value 값이, exist에는 true/false 값이 초기화
- value 값만 반환받고 싶다면 변수 한 개만 선언<br>
  ex) val := 맵이름[key] -> val에 value 값이 초기화
- true/false 값만 반환받고 싶다면 "\_, bool변수"형식으로 선언

```go
package main

import "fmt"

func main() {
	//지역번호와 지역 저장
	var m = make(map[string]string)

	m["02"] = "서울특별시"
	m["031"] = "경기도"
	m["032"] = "인천"
	m["053"] = "대구광역시"

	fmt.Println(m["032"])
	fmt.Println(m["042"], "빈 칸입니다.") //string형태로 존재하지 않는 key값은 ""가 출력된다

	val, exist := m["02"] //존재하는 key
	fmt.Println(val, exist)

	val, exist = m["042"] //존재하지 않는 key
	fmt.Println(val, exist)

	val = m["053"] //value 값만 반환
	fmt.Println(val)

	_, exist = m["053"] //true/false 값만 반환
	fmt.Println(exist)

	//맵도 똑같이 len() 함수를 사용할 수 있다 하지만 cap() 함수는 사용할 수 없다
	fmt.Println(len(m))
}
```

---

# 연습문제

1. 역행렬

```go
package main

import "fmt"

func main() {

	var metrix [2][2]int = [2][2]int{{7 ,3 } , { 5 , 2}}

	var isReverse bool

	d := metrix[0][0] * metrix[1][1] - metrix[0][1] * metrix[1][0]

	isReverse = ( d != 0 )

	fmt.Println( isReverse )

	if isReverse {
		fmt.Println( metrix[1][1] / d , (-1) * metrix[0][1] / d)
		fmt.Println( (-1) * metrix[1][0] /d , metrix[0][0] /d )
	}
}
```

2. 가장 긴 이름

```go
package main

import "fmt"

func main() {
	names := make([]string, 1, 2)

	var name string

	for name != "1"{
		fmt.Scanln(&name)
		names = append(names, name)

		//가장 긴 이름을 인덱스 0 에 저장
		if len(names[0]) < len(name) {
			names[0] = name
		}
	}

	var result string = names[0]

	fmt.Println(result, len(result))
}
```

3. 중간고사 평균점수

```go
package main

import "fmt"

func main() {
	var scores = make(map[string]int)

	var subject string
	var score int

	for {
		fmt.Scanln( &subject, &score)
		if subject == "0" {
			break
		}

		scores[subject] = score
	}

	var avg float32 = 0.0;

	for key, value := range scores {
		fmt.Printf( "%s %d\n" , key, value )
		avg += float32( value );
	}


	fmt.Printf( "%.2f\n" , avg / float32(len(scores) ) )
}
```
