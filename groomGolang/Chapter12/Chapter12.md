# Chapter12 Closer 클로저

## 클로저란?

- 클로저는 함수 안에서 익명 함수를 정의해서 바깥쪽 함수에 선언한 변수에도 접근할 수 있는 함수

EX

```go
package main

import "fmt"

func main() {
	a, b := 10, 20
	str := "Hello goorm!"

	result := func () int{ // 익명함수 변수에 초기화
		return a + b // main 함수 변수 바로 접근
	}()

	func() {
		fmt.Println(str) // main 함수 변수 바로 접근
	}()

	fmt.Println(result)
}
```

- 함수 안에서 함수를 반환할때 특징!
- 익명함수 밖에 변수를 참조시에 외부 변수 상태를 유지함
- 새로운 변수에 함수를 초기화시에 변수도 초기화

```go
package main

import "fmt"

func next() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextInt := next()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt := next()
	fmt.Println(newInt())
}
```

출력

```go
1
2
3
1
```

---

## 연습문제

```go
package main

import "fmt"

/*타입문 작성은 선택입니다*/

func calCoin( coinValue int ) func( n int) int {
	return func(coinCount int ) int { //클로저
		return  coinValue * coinCount
	}
}

func main() {
	var coin10, coin50, coin100, coin500 int
	fmt.Scan(&coin10, &coin50, &coin100, &coin500)

	if coin10 < 0 || coin50 < 0 || coin100 < 0 || coin500 < 0 {
		fmt.Println("잘못된입력입니다.")
		return
	}

	add10 := calCoin(10)
	add50 := calCoin(50)
	add100 :=	calCoin(100)
	add500 := calCoin(500)

	//이렇게 만들라는 얘기였군요...
	totalmoney := add10(coin10) + add50( coin50 ) + add100( coin100 ) + add500( coin500 )

	fmt.Println(totalmoney)
}
```
