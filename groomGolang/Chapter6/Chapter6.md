# Chapter6

## 반복문!!

- Go에서는 while이 없다!
- For문만 있다!

기본적인 사용

```Go
/*
for 초기식; 조건식; 조건 변화식 {
	반복 수행할 구문
}
*/

package main

import "fmt"

func main() {
	sum := 0

	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("1부터 10까지 정수 합계:", sum)
}
```

---

## For문 종류!

### 조건문만!

- 말그대로 조건문만 존재 ( while 문과 비슷한듯?)

example

```Go
package main

func main() {
	n := 2

	for n < 100 {
		n *= 2
	}
}
```

---

### 무한루프

- 아무것도 안쓰고 for만 사용

```Go
package main

import "fmt"

func main() {
	for {
		fmt.Printf("무한루프입니다.\n")
	}
}
```

---

### for range 문

- 다른 언어의 foreach문과 비슷
- " for 인덱스, 요소값 := range 컬렉션이름 " ( '컬랙션' -> 지금은 배열, 순환이 가능한? 정도의 객체정도로 생각하자 )

배열

```Go
package main

import "fmt"

func main() {
	var arr [6]int = [6]int{1, 2, 3, 4, 5, 6}

	//일반적인 사용
	for index, num := range arr {
		fmt.Printf("arr[%d]의 값은 %d입니다.\n", index, num)
	}

	//요소값만 사용
	for _, num := range arr {
		fmt.Printf("이번 num의 값은 %d입니다.\n", num)
	}

	//인덱스만 사용
	for index := range arr {
		fmt.Printf("arr[%d]의 값은 %d입니다.\n", index, arr[index])
	}
}

```

- 각 컬랙션의 경우 조금씩 다른형태

Map의 경우

- key, value 값으로 받아옵니다.

```Go
package main

import "fmt"

func main() {
	var fruits map[string]string = map[string]string{
		"apple":  "red",
		"banana": "yellow",
		"grape":  "purple",
	}

	for fruit, color := range fruits {
		fmt.Printf("%s의 색깔은 %s입니다.\n", fruit, color)
	}
}
```

---

## 연습문제!

1. 구구단

```Go
package main

import "fmt"

func main() {
	var dan int
	fmt.Scanln(&dan)

	for n := 1; n < 10; n++ {
		fmt.Printf("%d X %d = %d\n", dan, n, dan*n)
	}
}
```

2. 이등변 삼각형?

```Go
package main

import "fmt"

func main() {
	var i, j, k int
	//i,j는 두 개의 반복문에 쓰일 변수

	fmt.Scanln(&k)

	for i = 0; i < k; i++ {
		for j = 0; j < i; j++ {
			fmt.Print("o ")
		}
		fmt.Println("* ")
	}
}
```
