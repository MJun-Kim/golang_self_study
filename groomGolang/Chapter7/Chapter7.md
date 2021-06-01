## 조건문 - if / else

---

### 조건에 따른 실행과 흐름 분기

- Go언어의 if문에는 몇 가지 엄격하다고 느낄 수 있는 특징이 존재
  - 길게 봤을 때 Go언어를 사용하는 개발자들에게 좋은 코딩 습관을 심어주고 코드 자체도 누구나 쉽게 알아볼 수 있게끔 해주는 장점이 `Go언어 개발자들이 지향했던 방향과 일치`

### 조건문

- #### 간단히 음료수 자판기를 생각하면

  - 돈이 들어왔을 경우 음료 선택을 위해 빨간불을 띄움
  - 음료 선택 시 넣은 돈보다 비쌀 경우 음료를 내보내면 안됨
  - 넣은 돈과 같거나 싼 음료를 선택하면 음료를 내보냄
  - 잔돈이 남았다면 잔돈을 내보냄

  이런 자판기와같은 간단하게 실생활에도 이렇게 다양한 조건문이 존재

<br />

- #### True / false
  - Go언어의 조건문은 반드시 Boolean 형으로 표현되어야 함
  - C, C++의 조건식에 0, 1을 사용할 수 있는 것과는 대조적으로 Go에서는 true, false만 사용 가능

<br />

- #### 조건식의 괄호는 생략 가능
  - 대부분의 언어들은 조건문에서 조건식을 쓸 때 if옆에 괄호를 입력 ex> "if(n == 0)"
  - 하지만 Go에선 "if n == 0"과 같이 괄호 생략 가능
  - Go에선 괄호를 생략해서 실행하는 것을 권장 Atom과 같은 에디터에서 패키지 설치하고 실행하면 자동으로 생략해줌
  - 하지만 괄호를 쓴다고해서 실행이 안되는 것은 아님

<br />

- #### 괄호의 시작과 else문은 같은 줄에
  - if ~else if문의 키워드가 반복해서 이어지는 독특한 구문이기 때문에 개행하는 방법이 개발자마다 다름
  - 하지만 어느정도 개발자들 사이에서 통용되는 코딩 방법이 존재하는데 그 방법을 Go 에서는 사용

#### 잘 된 예제

```js
package main

import "fmt"

func main() {
	var num int

	fmt.Print("정수입력 :")
	fmt.Scan(&num)

	if num == 1 {
		fmt.Print("hello\n")
	} else if num == 2 {
		fmt.Print("world\n")
	} else {
		fmt.Print("worng number..\n")
	}
}
```

<br />

#### 잘못된 예제

```js
	fmt.Print("정수입력 :")
	fmt.Scan(&num)

	if num == 1
	{
		fmt.Print("hello\n")
	}
	else if num == 2
	{
		fmt.Print("world\n")
	} else
	{
		fmt.Print("worng number..\n")
	}
```

<br />

- #### 조건식에 간단한 문장(Optional Statement) 실행 가능
  - Go에서는 조건식을 실행하기 전에 간단한 문장을 함께 실행 가능
  - "if val := num \* 2; val == 2" 과 같이 조건식 앞에 변수를 선언하고 식을 입력 가능
  - 여기서 주의할 점은 `조건식 전에 정의된 변수는 해당 조건문 블록에서만 사용할 수 있음`
  - 이러한 표현은 `switch`, `for` 등 여러 문법에서 사용 가능

#### 위 예제

```js
package main

import "fmt"

func main() {
	var num int

	fmt.Print("정수입력 : ")
	fmt.Scan(&num)

	if val := num * 2; val == 2 {
		fmt.Print("hello\n")
	} else if val := num * 3; val == 6 {
		fmt.Print("world\n")
	} else {
		fmt.Print("worng number..\n")
	}
}
```

출력 값

```
정수입력 : 2
world
```

<br />

#### 조건문 총 예제

```js
package main

import "fmt"

func main() {
	var opt int
	var num1, num2, result float32

	fmt.Print("1.덧셈 2.뺄셈 3.곱셈 4.나눗셈 선택 : ")
	fmt.Scan(&opt)
	fmt.Print("두 개의 실수 입력 : ")
	fmt.Scan(&num1, &num2)

	if opt == 1 {
		result = num1 + num2
	} else if opt == 2 {
		result = num1 - num2
	} else if opt == 3 {
		result = num1 * num2
	} else if opt == 4 {
		result = num1 / num2
	}

	fmt.Printf("결과: %f\n", result)
}
```

출력 값

```
1.덧셈 2.뺄셈 3.곱셈 4.나눗셈 선택 : 2
두 개의 실수 입력 : 5 10
결과: -5.000000
```

<br />

### 실습 1. 7과 9의 배수

```js
package main

import "fmt"

func main() {

	for i := 1; i < 100; i++ {
		if i % 7 == 0 {
			fmt.Printf("%d ", i);
		} else if i % 9 == 0 {
			fmt.Printf("%d ", i);
		}
	}
}
```

출력 값

```
7 9 14 18 21 27 28 35 36 42 45 49 54 56 63 70 72 77 81 84 90 91 98 99
```

<br />

### 실습 2. 두 수의 차

```js
package main

import "fmt"

func main() {
	var num1, num2, result int;

	fmt.Scanf("%d %d", &num1, &num2);

	if num1 > num2{
		result = num1 - num2
	} else {
		result = num2 - num1
	}

	fmt.Println(result);
}
```

입력 값

```
8 13
```

출력 값

```
5
```
