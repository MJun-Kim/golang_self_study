# 제어문

---

## `break`, `continue`, `goto` 문

- 제어문을 반복문에서 주로 사용하는데 반복문을 사용하다보면 같은 구문을 반복해서 실행하기 때문에 이를 제어할 상황이 옴
- 예를 들어, `1 ~ 10까지 계속 순차적으로 출력하다가 3의 배수는 건너뛰어`, `1 ~ 10까지 출력하다가 4가 되면 반복문을 빠져나와`, `1 ~ 10까지 출력하다가 4가 되면 50 ~ 80을 출력하는 부분으로 가`
- 반복문을 탈출하는 `break`, 반복문의 첫 부분으로 돌아가는 `continue`, 특정 부분으로 갈 수 있는 `goto`

<br />

## 이 곳을 빠져나와 ! - 똑똑한 `break`

- 쉽게 말해 `해당 부분`을 빠져나오는 용법
- `해당 부분`이라 표현한 이유는 `break`가 꼭 `for`문 뿐만이 아니라, `switch`, `select`문에서도 사용 가능하기 떄문. `for`문과 동일하게 모든 용법에서 그 쓰임이 동일
- 하지만 제어문 자체가 반복문에서 많이 쓰이기 때문에 이번 챕터에선 반복문에서의 쓰임을 중심으로 설명함

1. `break` 문은 '직속 for'를 빠져나옴. 여러 for문이 중첩되어 사용될 때 `break`문을 쓰면 `break`문이 존재한 for문만 빠져나오게 됨. 빠져나온 바로 다음 문장을 실행
2. `break`문은 보통 단독으로 실행되지만, 경우에 따라 `break 레이블이름`과 같이 사용되어 지정된 레이블로 이동 가능. 쉽게 말해 해당 for문을 빠져나옴과 동시에 지정한 레이블로 감. 소제목에서 `똑똑한 break`라고 소개를 했는데 `break 레이블이름`으로 되어있을 때 직속 for문을 빠져나와 해당 레이블로 이동하고 `break문이 빠져나왔던 for문은 건너뛰고 다음 문장을 실행

### 설명 2번에 대한 예제

```js
package main

import "fmt"

func main() {
	i := 0

TEST1:
	for {
		if i == 0 {
			break TEST1
		}
	}

	fmt.Println("End")
}
```

실행 결과

```
End
```

<br />

### 또 다른 예제

```js
package main

import "fmt"

func main() {
	var sum = 0
	var i int

	i = 1

	for {
		sum += i
		if sum > 100 {
			break
		}
		i++
	}

	fmt.Println("1에서 ", i, " 까지 더하면 처음으로 100이 넘어요.")
	fmt.Println("합계:", sum)
}
```

실행결과

```
1에서  14  까지 더하면 처음으로 100이 넘어요.
합계: 105
```

<br />

## 원하는 조건을 걸러주는 `continue`

- 소제목에서 말한 그대로 명시한 조건을 걸러주는 기능
- `break`와는 다르게 `for`문과 연관되어 사용해야만 함. 왜냐면 `continue`문을 만나면 해당 반복문의 조건검사(반복문의 처음) 부분으로 다시 이동하기 때문

### 예제

```js
package main

import "fmt"

func main() {

	for i := 1; i < 16; i++ {
		if i%2 == 0 {
			continue //반복문 처음 부분으로 간다
		}

		fmt.Printf("%d ", i)
	}
}
```

실행결과

```
1 3 5 7 9 11 13 15
```

<br />

## 그 곳으로 바로 가줘! - 하지만 잘 안 쓰이는 goto

- `goto` 그 이름이 의미하듯 프로그램의 흐름을 원하는 위치로 이동시킬 때 사용함
- 위치는 레이블로 표시. `goto 레이블명` 입력 시 해당 레이블로 흐름이 이동함
- 아주 오래전에는 `goto`의 필요성에 대한 논쟁이 있었지만, 요즘은 `goto`의 사용에 대해 부정적. 실제 근래에 출간되는 프로그래밍 언어 서적중에는 `goto`를 아예 언급조차 안함
- `goto`를 부정적으로 보는 이유는, `프로그램의 자연스러운 흐름을 방해한다`는 것. 프로그램에 있어 흐름은 굉장히 중요한 부분인데 그렇다고 해서 꼭 `goto`를 사용해야만 해결할 수 있는 문제 상황도 딱히 존재하지 않음
- 사용하지 않을 거라면 그냥 넘어가도 좋지만, 알고 사용안하는 것과 모르고 사용하지 않는 것에는 차이가 있기 때문에 간단히 알고 넘어가는게 좋음

### 예제

```js
package main

import "fmt"

func main() {
	var num int

	fmt.Print("자연수를 입력하세요 : ")
	fmt.Scanln(&num)

	if num == 1 {
		goto ONE
	} else if num == 2 {
		goto TWO
	} else {
		goto OTHER
	}

ONE:
	fmt.Print("1을 입력했습니다.\n")
	goto END
TWO:
	fmt.Print("2를 입력했습니다.\n")
	goto END
OTHER:
	fmt.Print("1이나 2를 입력하지 않으셨습니다!\n")
END:
}
```

실행결과

```
자연수를 입력하세요 : 1
1을 입력했습니다.
```

<br />

## 실습예제 - 구구단 2

```js
package main

import "fmt"

func main() {
	for  i := 2; i <= 9; i++ {
		if i % 2 == 0 {
			continue;
		}
		for j := 1; j <= i; j++ {
			fmt.Printf("%d x %d = %d\n", i, j, i * j);
		}
		fmt.Printf("\n");
	}
}
```

실행결과

```
3 x 1 = 3
3 x 2 = 6
3 x 3 = 9

5 x 1 = 5
5 x 2 = 10
5 x 3 = 15
5 x 4 = 20
5 x 5 = 25

7 x 1 = 7
7 x 2 = 14
7 x 3 = 21
7 x 4 = 28
7 x 5 = 35
7 x 6 = 42
7 x 7 = 49

9 x 1 = 9
9 x 2 = 18
9 x 3 = 27
9 x 4 = 36
9 x 5 = 45
9 x 6 = 54
9 x 7 = 63
9 x 8 = 72
9 x 9 = 81
```

<br />

## 실습예제 - 두 수를 더하면 99

```js
package main

import "fmt"

func main() {
	var result int

	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			result = (i * 10 + j) + (j * 10 + i)
			if result == 99 {
				fmt.Printf("%02d + %02d = 99\n", i * 10 + j, j * 10 + i);
			}
			continue;
		}
	}
}
```

실행결과

```
09 + 90 = 99
18 + 81 = 99
27 + 72 = 99
36 + 63 = 99
45 + 54 = 99
54 + 45 = 99
63 + 36 = 99
72 + 27 = 99
81 + 18 = 99
90 + 09 = 99
```
