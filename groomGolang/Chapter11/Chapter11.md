# 함수

---

## 함수란

- 특정 기능을 위해 만든 여러 문장을 묶어서 실행하는 코드 블록 단위
- 쉽게 말해, 프로그램의 특정 기능들을 기능별로 묶어 구현해 놓은 것
- 설계의 가장 기본이자 전부라고 할 수 있음

<br />

## Go언어에서 함수

- 다양한 방법으로 활용할 수 있도록 쓰임새롤 유연하게 만듬
- 하지만 앞서 배운 Go였듯, 개발자들의 가독성 좋은 코드를 위해 몇 가지 지켜야 할 문법 존재
- 기본적인 형태의 함수 선언은 `func 함수이름 (매개변수이름 매개변수형) 반환형`

<br />

### Go 언어에서의 기본적인 특징

1. 함수를 선언할 때 쓰는 키워드는 'func'이다
2. `반환형`이 괄호`()` 뒤에 명시된다. 물론 `매개변수형`도 `매개변수이름` 뒤에 명시된다.
3. 함수는 패키지 안에서 정의되고 호출되는 함수가 꼭 호출하는 함수 앞에 있을 필요는 없다.
4. '반환값'이 여러 개일 수 있다. 이럴 때는 '반환형'을 괄호로 묶어 개수만큼 입력해야한다. ((반환형1, 반환형2)형식, 두 형이 같더라도 두 번 써야 한다)
5. 블록 시작 브레이스 `{`가 함수 선언과 동시에 첫 줄에 있어야 한다(모든 용법을 이렇게 쓰는 것이 좋습니다).

- 특이한 점으로는 <b>Go언어에서는 반환 값도 여러 개일 수 있음</b>
- 함수는 기본적으로 매개변수와 리턴 값의 유, 무에 따라 네 개의 형태로 나눌 수 있음
  ```
  - 매개변수가 있고, 반환 값도 있는 형태
  - 매개변수가 있고, 반환 값이 없는 형태
  - 매개변수가 없고, 반환 값이 있는 형태
  - 매개변수가 없고, 반환 값이 없는 형태
  ```
- 예제

```js
package main

import "fmt"

/*기능들만 모아놓은 함수들*/
func guide() { //매개변수 X 반환 값 X
	fmt.Println("두 정수를 입력받고 곱한 결과를 출력하는 프로그램입니다.\n두 정수를 차례로 띄어 써주세요.")
	fmt.Print("두 정수를 입력해주세요 :")
}

func input() (int, int) { //매개변수 X 반환 값 O(두 개)
	var a, b int
	fmt.Scanln(&a, &b)
	return a, b
}

func multi(a, b int) int { //매개변수 O, 반환 값 O
	return a * b
}

func printResult(num int) { //매개변수 O, 반환 값 X
	fmt.Printf("결과값은 %d입니다. 프로그램을 종료합니다.\n", num)
}

func main() { //간결해진 main문
	guide()
	num1, num2 := input()
	result := multi(num1, num2)
	printResult(result)
}
```

실행결과

```
두 정수를 입력받고 곱한 결과를 출력하는 프로그램입니다.
두 정수를 차례로 띄어 써주세요.
두 정수를 입력해주세요 :3 5
결과값은 15입니다. 프로그램을 종료합니다.
```

<br />

## 전역변수와 지역변수

- `지역변수`와 `전역변수`는 선언되는 위치에 따라 그 유형이 결정됨
- `지역변수`는 중괄호`{}` 안에서 선언된 변수를 말함
  - 이렇게 선언된 `지역변수`는 선언된 지역 내에서만 유효
- 전역변수는 특정 지역(중괄호) 밖에서 선언된 변수를 말힘
  - 이 변수는 지역과 관계없이 어느 곳에든 유효합니다
- 따라서 이 두 변수는 두 가지의 차이점 존재
  - 메모리에 존재하는 시간
    - 지역변수는 해당 지역에서 선언되는 순간 메모리가 생성되고 해당 지역을 벗어나면 자동으로 소멸
    - 전역변수는 코드가 시작되어 선언되는 순간 메모리가 생성되고 코드 전체가 끝날때까지 메모리를 차지
  - 변수에 접근할 수 있는 범위

### 지역변수 예제

```js
package main

import "fmt"

func exampleFunc1() {
	var a int = 10 //지역변수 선언

	a++

	fmt.Println("exampleFunc1의 a는 ", a)
}

func exampleFunc2() {
	var b int = 20 //지역변수 선언
	var c int = 30 //지역변수 선언

	b++
	c++

	fmt.Println("b와 c는 ", b, c)
}

func main() {
	var a int = 28 //지역변수 선언

	exampleFunc1()
	exampleFunc2()

	fmt.Println("main의 a는", a)
}
```

실행결과

```
exampleFunc1의 a는  11
b와 c는  21 31
main의 a는 28
```

<br />

### 전역변수 예제

```js
package main

import "fmt"

var a int = 1 //전역변수 선언

func localVar() int { //지역변수로 연산
	var a int = 10 //지역변수 선언

	a += 3

	return a
}

func globalVar() int { //전역변수로 연산
	a += 3

	return a
}

func main() {
	fmt.Println("지역변수 a의 연산: ", localVar())
	fmt.Println("전역변수 a의 연산: ", globalVar())
}
```

실행결과

```
지역변수 a의 연산:  13
전역변수 a의 연산:  4
```

<br />

## 매개변수

- Go에서 매개변수는 `Pass by value`, `Pass by reference`, `가변 인자`에 대해 알면 됨
- `가변인자`는 변수의 접근 범위 내용과 좀 다르지만 매개변수와 관련된 내용이라 이번 강의에서 다룸

### Pass by value

- 인자의 값을 복사해서 전달하는 방식
- 따라서 복사한 값을 함수 안에서 어떠한 연산을 하더라도 원래 값은 변하지 않음
- 함수를 호출할 때는 `함수이름(변수이름)`
- 예제

```js
package main

import "fmt"

func printSqure(a int) {
	a *= a

	fmt.Println(a)
}
func main() {
	a := 4 //지역변수 선언

	printSqure(a)

	fmt.Println(a)
}
```

실행결과

```
16
4
```

<br />

### Pass by reference

- Java를 잘 안다면 이해하기 편함. Java는 객체 지향 언어로서 기본형 데이터 이외에 클래스 변수들은 전부 pass by reference를 기초로 두고있음
- Pass by value와 Pass by reference를 구분해 배우는 것은 Go언어는 분명 객체 지향을 따른다고 했지만 형태와 용법을 보았다시피 C언어와 비슷한 모습
- 따라서 Go언어에서는 C/C++ 언어에서 핵심 개념인 `포인터`라는 개념을 지원
- `&` : 주소, `*` : 직접참조. Go언어는 c언어의 다양하고 복잡한 포인터와 달리 포인터의 핵심 개념만 사용하도록 제공

  - C언어에서는 배열이름 자체가 배열의 첫번째 인덱스 요소의 주솟값인데 Go언어는 그런 것이 없음. 주솟값은 어떤 변수 앞에 &를 붙이는 것만 기억하면 됨
  - C언어에서는 "*(배열이름+인덱스)"는 "배열이름[인덱스]"와 같은 기능을 했는데 Go언어는 그런 것이 없습니다. 직접 참조를 원하면 포인터 변수 앞에 *를 붙이는 것만 기억하면 됩니다
  - 함수 호출시엔 주솟값 전달을 위해 `함수이름(&변수이름)` 입력.
  - 매개변수 이름을 입력할 땐 값을 직접 참조하기 위해 `*을 매개변수형 앞에 붙힘`
  - 그리고 함수 안에서 매개변수 안에 전부 `*`을 붙혀야 함

- 예제

```js
package main

import "fmt"

func printSqure(a *int) {
	*a *= *a

	fmt.Println(*a)
}
func main() {
	a := 4         //지역변수 선언

	printSqure(&a) //참조를 위한 a의 주솟값을 매개변수로 전달

	fmt.Println(a)
}
```

실행결과

```
16
16
```

<br />

### 가변 인자 함수

- 전달하는 매개변수의 개수를 고정한 함수가 아니라 함수를 호출할 때마다 매개변수의 개수를 다르게 전달할 수 있도록 만든 함수
- 기능이 좀 다르지만 비슷한 개념으로 Java에서는 메소드 `오버로딩` 존재
- 오버로딩은 매개변수 형태와 개수에 따라 메소드를 만들어야하지만 Go언어의 가변 인자 함수는 동일한 형의 매개변수를 n개 전달가능
  - 예를 들어 숫자들을 더하는 함수를 만든다고 할 때, 더하는 값의 개수에 따라 각각 함수를 만들 필요가 없다는 것
- 가변 인자 함수 사용 시 몇 가지 특징
  ```
  1. n개의 동일한 형의 매개변수를 전달함
  2. 전달된 변수들은 슬라이스 형태. 변수를 다룰 때 슬라이스를 다루는 방법과 동일함
  3. 함수를 선언은 'func 함수이름(매개변수이름 ...매개변수형) 반환형' 형식으로 힘. '매개변수형' 앞에 '...'을 붙이면 됨
  4.
  ```
- 예제

```js
package main

import "fmt"

func addOne(num ...int) int {
	var result int

	for i := 0; i < len(num); i++ { //for문을 이용한 num[i] 순차 접근
		result += num[i]
	}

	return result
}

func addTwo(num ...int) int {
	var result int

	for _, val := range num { //for range문을 이용한 num의 value 순차 접근
		result += val
	}

	return result
}

func main() {
	num1, num2, num3, num4, num5 := 1, 2, 3, 4, 5
	nums := []int{10, 20, 30, 40}

	fmt.Println(addOne(num1, num2))
	fmt.Println(addOne(num1, num2, num4))
	fmt.Println(addOne(nums...))
	fmt.Println(addTwo(num3, num4, num5))
	fmt.Println(addTwo(num1, num3, num4, num5))
	fmt.Println(addTwo(nums...))
}
```

실행결과

```
3
7
100
12
13
100
```

<br />

## 반환값(리턴값)

- Go언어에서는 복수개의 반환값을 반환할 수 있음
  - 반환값의 개수만큼 반환형을 명시해야 함. 2개 이상의 반환형을 입력할 때는 괄호`()`안에 명시
  - 동일한 반환형이더라도 모두 명시해야 함. `(int, int, int)`

#### 예제

```js
package main

import "fmt"

func add(num ...int) (int, int) {
	var result int
	var count int

	for i := 0; i < len(num); i++ { //for문을 이용한 num[i] 순차 접근
		result += num[i]
		count++
	}

	return result, count
}

func main() {
	nums := []int{10, 20, 30, 40, 50}

	fmt.Println(add(nums...))
}
```

실행결과

```
150 5
```

<br />

### Named Return Parameter

- 직역하면 `이름이 붙여진 반환 인자` 즉, 이름이 붙혀진 반환 값
- 여러 개의 값을 반환할 때 괄호 안에 반환형을 모두 명시해야 하는데 반환 값이 많고 반환형이 다양하다면 가독성이 좋지 않을 수 있음
- 코드 안에서 return 뒤에 명시하던 리턴 값들을 반환형 앞에 명시하는 것
- Named return parameter의 몇 가지 특징

```
1. (반환값이름1 반환형1, 반환값이름2 반환형2, 반환값이름3 반환형3, ...) 형식으로 입력
2. '반환값이름 반환형' 자체가 변수 선언. 따라서 함수 안에서 따로 선언할 필요가 없고 만약 선언하면 에러 발생
3. 'return'을 생략하면 안되며 반환 값이 있을 때는 반드시 return을 명시
4. 반환 값이 하나라도 반환값이름을 명시했다면 괄호 안에 써야함
```

#### 예제

```js
package main

import "fmt"

func dessertList(fruit ...string) (count int, list []string) { //여기서 이미 선언된 것이다

	for i := 0; i < len(fruit); i++ {
		list = append(list, fruit[i])
		count++
	}

	return //생략하면 안 된다
}

func inputFruit() (list []string) { //Named return parameter는 값이 하나라도 괄호를 써야한다

	for {
		var fruit string
		fmt.Print("과일을 입력하세요:")
		fmt.Scanln(&fruit)

		if fruit != "1" {
			list = append(list, fruit)
		} else {
			fmt.Println("입력을 종료합니다.\n")
			break //반복문을 빠져나간다
		}
	}

	return
}

func main() {
	fmt.Println("디저트로 먹을 과일을 입력하고 출력합니다. \n1을 입력하면 입력을 멈춥니다.\n")
	count, list := dessertList(inputFruit()...) //함수를 변수처럼 사용할 수 있습니다
	fmt.Printf("%d개의 과일을 입력하셨고, 입력한 과일의 리스트는 %s입니다.\n", count, list)
}
```

실행결과

```
디저트로 먹을 과일을 입력하고 출력합니다.
1을 입력하면 입력을 멈춥니다.

과일을 입력하세요: 사과
과일을 입력하세요:바나나
과일을 입력하세요:수박
과일을 입력하세요:1
입력을 종료합니다.

3개의 과일을 입력하셨고, 입력한 과일의 리스트는 [사과 바나나 수박]입니다.
```

<br />

## 익명함수

함수를 만드는 것에 대한 단점이 존재하는데 `프로그램 속도 저하` 입니다. 왜냐하면

- 함수 선언 자체가 프로그래밍 전역으로 초기화되면서 메모리를 잡아먹기 때문
- 기능을 수행할 때마다 함수를 찾아서 호출해야하기 때문입니다.
  이러한 단점을 보완하기 위해 `익명함수`가 필요

#### 간단한 예제

```js
package main

import "fmt"

func main() {
	func() {
		fmt.Println("hello")
	}()

	func(a int, b int) {
		result := a + b
		fmt.Println(result)
	}(1, 3)

	result := func(a string, b string) string {
		return a + b
	}("hello", " world!")
	fmt.Println(result)

	i, j := 10.2, 20.4
	divide := func(a float64, b float64) float64 {
		return a / b
	}(i, j)
	fmt.Println(divide)
}
```

실행결과

```
hello
4
hello world!
0.5
```

- 위 예시 코드에서 확인할 수 있는 익명 함수의 가장 큰 특징은 그 자리에서 만들고 그 자리에서 바로 실행하는 것
- 익명함수의 기본적인 2가지 형태
  - 함수의 이름만 없고 그 외에 형태는 동일
  - 함수의 블록 마지막 브레이스`}` 뒤에 괄호`()`를 사용해 함수를 바로 호출합니다. 이때, 괄호 안에 매개변수를 넣을 수 있음
- 익명 함수는 함수의 '기능적인 요소'만 쏙 빼와서 어디서든 가볍게 활용하기 위해 사용하는 것

- 선언 함수는 반환 값을 변수에 초기화함으로써 변수에 바로 할당이 가능. 익명 함수도 똑같은 기능을 하는데, 여기서 차이점은 변수에 초기화한 익명 함수는 변수 이름을 함수의 이름처럼 사용할 수 있음
  Example )
  ```js
  package main

      import "fmt"

      func addDeclared(nums ...int) (result int) {
      	for i := 0; i < len(nums); i++ {
      		result += nums[i]
      	}
      	return
      }

      func main() {
      	var nums = []int{10, 12, 13, 14, 16}

      	addAnonymous := func(nums ...int) (result int) {
      		for i := 0; i < len(nums); i++ {
      			result += nums[i]
      		}
      		return
      	}

      	fmt.Println(addAnonymous(nums...))
      	fmt.Println(addDeclared(nums...))
      }
      ```

      실행결과
      ```
      65
      65
      ```

  <br />

- 선언 함수와 익명 함수는 프로그램 내부적으로 읽는 순서가 다름.
- 선언 함수는 프로그램이 시작됨과 동시에 모두 읽지만 익명 함수는 위 예시들처럼 그 자리에서 실행되기 때문에 해당 함수가 실행되는 곳에서 읽음. `즉, 선언 함수보다 익명 함수가 나중에 읽힘`

<br />

### 일급 함수(First-Class Function)

- 익명 함수의 사용은 Go언어에서의 함수가 '일급 함수'이기 때문에 가능한 것
- 일급 함수라는 의미는 함수를 기본 타입과 동일하게 사용할 수 있어 함수 자체를 다른 함수의 매개변수로 전달하거나 다른 함수의 반환 값으로 사용될 수 있다는 것
- <b>따라서 함수는 다른 타입들과 비교했을 때 높은 수준의 용법이 아니라 같은 객체로서 사용될 수 있음 </b>

#### 예제

```js
package main

import "fmt"

func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}

func main() {
	multi := func(i int, j int) int {
		return i * j
	}

	r1 := calc(multi, 10, 20)
	fmt.Println(r1)

	r2 := calc(func(x int, y int) int { return x + y }, 10, 20)
	fmt.Println(r2)
}
```

실행결과

```
200
30
```

<br />

### type문을 사용한 함수 원형 정의

- 전달 받는 함수가 매개변수가 5개고 반환형이 6개일 때는 그 함수를 매개변수로 사용할 때마다 그만큼을 명시해야 됨
- 따라서 이를 극복하기 위해 Go언어에서는 'type'문을 사용해 함수의 원형을 정의하고 사용자가 정의한 이름을 형으로써 사용
- 이러한 사용자의 Custom Type은 C언어의 '구조체' 개념과 유사

- 간단한 예제

```js
package main

import "fmt"

//함수 원형 정의
type calculatorNum func(int, int) int
type calculatorStr func(string, string) string

func calNum(f calculatorNum, a int, b int) int {
	result := f(a, b)
	return result
}

func calStr(f calculatorStr, a string, b string) string {
	sentence := f(a, b)
	return sentence
}

func main() {
	multi := func(i int, j int) int {
		return i * j
	}
	duple := func(i string, j string) string {
		return i + j + i + j
	}

	r1 := calNum(multi, 10, 20)
	fmt.Println(r1)

	r2 := calStr(duple, "Hello", " Golang ")
	fmt.Println(r2)
}
```

실행결과

```
200
Hello Golang Hello Golang
```

- type문은 함수 원형 정의 뿐만이 아니라 구조체, 인터페이스 등을 정의하기 위해 사용

#### 익명 함수는 가볍게 사용할 수 있고 선언 함수의 쓰임새를 확장했으며 이를 이용해 익명 함수는 뒤에 나오는 챕터인 `클로저`, `defer`, `Go루틴`에서 많이 다루게 됨

<br />

## 실습문제

### 오름차순 정렬

```js
package main

import "fmt"

func bubbleSort(nums []int) {
	var temp int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				temp = nums[i];
				nums[i] = nums[j];
				nums[j] = temp;
			}
		}
	}
}

func inputNums() []int{
	var n, num int;
	nums := make([]int, 0);
	fmt.Scanf("%d", &n);

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &num);
		nums = append(nums, num);
	}
	return nums;
}

func outputNums(nums []int) {
	for i := 0; i < len(nums); i++ {
		fmt.Print(nums[i], " ");
	}
}

func main() {
	nums := inputNums()
	bubbleSort(nums)
	outputNums(nums)
}
```

입력값

```
5
6
3
4
2
56
```

실행결과

```
2 3 4 6 56
```

<br />

### 2. 아이패드를 사주는 조건

```js
package main

import "fmt"

func inputNums() []int {
	var score int;
	scores := make([]int, 0);

	for i := 0; i<5; i++{
		fmt.Scanf("%d", &score);
		scores = append(scores, score);
	}

	return scores;
}


func calExam(arr []int) (int, int, int) {
	var sum, more, less int;
	for i := 0; i < len(arr); i++ {
		sum += arr[i];
		if arr[i] >= 90 {
			more++;
		} else if(arr[i] < 70) {
			less++;
		}
	}

	return sum, more, less;
}



func printResult (sum, more, less int) {
	var result bool = true

	if sum < 400 {
		fmt.Println("총점이 400점 미만입니다.");
		result = false;
	}
	if more < 2 {
		fmt.Println("90점 이상 과목 수가 2개 미만입니다.");
		result = false;
	}
	if less != 0 {
		fmt.Println("70점 미만 과목이 있습니다.");
		result = false;
	}

	if result == false {
		fmt.Println("아이패드를 살 수 없습니다.");
	}else{
		fmt.Println("아이패드를 살 수 있습니다.");
	}
}
func main() {
	scores := inputNums();
	sum, more, less := calExam(scores);
	printResult(sum, more, less);
}
```

입력값

```
70
100
89
89
89
```

실행결과

```
90점이상과목수가2개미만입니다.
아이패드를살수없습니다.
```

<br />

### 3.역학적 에너지

```js
package main

import "fmt"

const g float32 = 9.8;

type kinEnergy func(float32, float32) float32
type potEnergy func(float32, float32) float32

func calKEnergy(kin kinEnergy, m float32, v float32) float32 {
	result := kin(m, v);
	return result;
}
func calPEnergy(pot potEnergy, m float32, h float32) float32 {
	result := pot(m, h);
    return result;
}

func main() {
	var m, v, h float32
	fmt.Scanln(&m, &v, &h);

	kinEnergy := func(m float32, v float32) float32 {
		return 0.5 * (m * v * v);
	}
	potEnergy := func(m float32, h float32) float32 {
		return m * g * h;
	}

	ke := calKEnergy(kinEnergy, m, v);
	pe := calPEnergy(potEnergy, m ,h);
	sum := ke + pe;
	fmt.Println(ke, pe, sum);
}
```

입력값

```
10.5 50.22 2.5
```

실행결과

```
13240.754 257.25 13498.004
```
