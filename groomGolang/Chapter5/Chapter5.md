## CH5. 콘솔 출력과 입력 함수

---

### 콘솔 출력 함수(Print)

- 프로그래밍의 가장 대표적인 함수인 `콘솔 출력 함수`는 c언어에선 `printf`, java에선 `System.out.println()`
- Go언어에서 사용하는 `fmt패키지`의 콘솔 출력함수
  - 기본 콘솔출력함수
    - `Println`
    - `Print`
    - `Printf`
  - 파일출력을 위한 콘솔출력함수 : 파일 출력을 위해 사용되고 사용법이 크게 다르지 않음
    - `Fprintln`
    - `Fprint`
    - `Fprintf`
  - string형으로 반환되는 콘솔 출력 함수 : 반환값을 string형으로 출력하는 함수. 특별한 용도에 사용되고 표준 출력을 하지 않기 때문에 이번챕터에서 다루지 않음
    - `Sprintln`
    - `Sprint`
    - `Sprintf`

---

### import "fmt"

- 코드 상단에 `import "fmt"` 선언하여 사용
- C언어의 입/출력 함수인 `printf`, `scanf`함수같은 유사한 함수들을 형식화한 입/출력 함수를 사용할 수 있다는 뜻
- 쉽게 말해, 콘솔 입력 함수와 출력 함수를 사용하기 위해선 fmt패키지를 import 해야함

### Println, Print

| 선언    | 출력형태                                           |
| ------- | -------------------------------------------------- |
| Println | 개행 자동 추가                                     |
| Print   | 개행 자동 추가하지 않음                            |
| Printf  | 포맷 지정자를 이용하여 개발자가 원하는 형태로 출력 |

- 하나의 인자를 출력할 수도 있고 여러 개의 인자를 콤마(,)로 결합해 열거할 수 있음

### 예제 1

```js
package main

import "fmt"

func main() {
	fmt.Println("안녕 난 길동이야")
	fmt.Println("나이는 24살이야")
	fmt.Println("반가워")

	n := 14
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Print("내 동생은", n, "살이야")
	fmt.Print(arr)
	fmt.Print(n, arr)
}
```

출력결과

```
안녕 난 길동이야
나이는 24살이야
반가워
반가워내 동생은14살이야[1 2 3 4 5]14 [1 2 3 4 5]
```

### 예제 2

```js
package main

import "fmt"

func main() {
	fmt.Print("안녕 난 길동이야")
	fmt.Print("나이는 24살이야\n") //개행을 위해 \n를 입력한다.
	fmt.Print("반가워")

	n := 14
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Print("내 동생은", n, "살이야")
	fmt.Print(arr)
	fmt.Print(n, arr)
}
```

출력결과

```
안녕 난 길동이야나이는 24살이야
반가워내 동생은14살이야[1 2 3 4 5]14 [1 2 3 4 5]
```

---

## Printf

- c언어의 printf와 유사
- 개발자가 포맷을 지정하여 원하는 형태로 출력할 때 사용
- 주의해야할 점은 `반드시 포맷을 지정해줘야 함`
- 예를 들어, `a := 5`로 변수 선언 후 `fmt.Println(a)` 이렇게는 출력 불가
- 한 개의 인자라도 출력하기 위해선 `fmt.Println("%d", a)` 형식으로 사용
- c언어와의 차이점은 배열을 한번에 출력 가능

```js
c언어
int arr[3] = {1, 2, 3};
printf("%d", arr); // 불가능

go언어
var arr [3]int = [3]int{1, 2, 3};
fmt.Println("%d", arr); // [1 2 3]
```

### 예제 1

```js
package main

import "fmt"

func main() {
	age, name := 24, "길동"

	fmt.Printf("안녕 난 %s이야\n", name)
	fmt.Printf("나이는 %d살이야\n", age)
	fmt.Printf("반가워")

	n := 14
	var arr [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Printf("내 동생은 %d살이야\n", n)
	fmt.Printf("%d \n", arr) //배열 한 번에 출력 가능
	fmt.Printf("%d %d", n, arr)
}
```

출력결과

```
안녕 난 길동이야
나이는 24살이야
반가워내 동생은 14살이야
[1 2 3 4 5]
14 [1 2 3 4 5]
```

---

## 서식 문자의 종류와 의미

- Printf함수를 사용하면서 포맷을 지정해줄 때 `%d`와 같은 문자 사용
- 이는 값을 입력하고 출력하는 데 있어서 서식을 지정하는 것
- 예를 들어 `"제 나이는 10진수로 %d살, 16진수로는 %X살 입니다."`라고 포맷 지정 가능
  - %d : 10진수 형태로 출력
  - %X : 16진수 형태로 출력, 알파벳은 대문자로 출력
  - 이러한 %d, %X와 같은 형태를 가리켜 `서식문자`라고함

### 예제 1

```js
package main

import "fmt"

func main() {
	age := 24

	fmt.Printf("제 나이는 10진수로 %d살, 16진수로 %X살입니다.", age, age)

}
```

출력결과

```
제 나이는 10진수로 24살, 16진수로 18살입니다.
```

### 서식문자 종류와 의미

| 서식문자 | 출력 형태                          |
| -------- | ---------------------------------- |
| %t       | bool                               |
| %b       | 2진수 정수                         |
| %c       | 문자                               |
| %d       | 10진수 정수                        |
| %o       | 8진수 정수                         |
| %x       | 16진수 정수, 소문자                |
| %X       | 16진수 정수, 대문자                |
| %f       | 10진수 방식의 고정 소수점 실수     |
| %F       | 10진수 방식의 고정 소수점 실수     |
| %e       | 지수 표현 실수, e                  |
| %E       | 지수 표현 실수, E                  |
| %g       | 간단한 10진수 실수                 |
| %G       | 간단한 10진수 실수                 |
| %s       | 문자열                             |
| %p       | 포인터                             |
| %U       | 유니코드                           |
| %v       | 타입                               |
| %T       | 모든 형식                          |
| %#v      | #을 이용해구분할 수 있는 형식 표현 |

- C언어와 다른점이 몇 가지 있는데 C언어에서는 %u, %x등과 같은 부호의 여부에 따라 표현이 다른 서식문자들이 있음
- Go언어에는 %v문자를 이용해 변수의 타입에 관계없이 출력 가능
- 포인터 서식문자인 %p는 값이 참조하는 주소값을 반환하는 서식문자
- 출력하는 형태를 따로 지정 가능한데, 원하는 출력 폭을 지정할 때는 `%(폭)d`형식으로 입력,
- 지정한 출력 폭에 0을 채워넣고 싶으면 `%0(폭)d`
- 출력할 때 왼쪽부터 출력을 원한다면 `%-(폭)d`
- 출력할 소수점 이하 자리를 지정할 때 `%.(자릿수)f`형식으로 지정 (이건 c언어에도 존재)

### 예제 1

```js
package main

import "fmt"

func main() {
	fmt.Printf("5>6=%b\n", 5 > 6)
	fmt.Printf("15는 2진수로 %b\n", 15)
	fmt.Printf("저의 성은 %c 입니다\n", '김')
	fmt.Printf("19는 10진수로 %d입니다.\n", 19)
	fmt.Printf("19는 8진수로 %o입니다.\n", 19)
	fmt.Printf("19는 16진수로 %x입니다.\n", 19)
	fmt.Printf("19는 16진수로 %X입니다.\n", 19)
	fmt.Printf("19.1234는 고정 소수점으로 %f입니다.\n", 19.1234)
	fmt.Printf("19.1234는 고정 소수점으로 %F입니다.\n", 19.1234)
	fmt.Printf("19.1234의 지수 표현은 %e입니다.\n", 19.1234)
	fmt.Printf("19.1234의 지수 표현은 %E입니다.\n", 19.1234)
	fmt.Printf("19.1234의 간단한 실수 표현은 %g입니다.\n", 19.1234) // 고정 소수점이 아님
	fmt.Printf("19.1234의 간단한 실수 표현은 %G입니다.\n", 19.1234) // 고정 소수점이 아님
	fmt.Printf("문자열: %s\n", "안녕하세요.")

	var num int = 50
	fmt.Printf("num은 %d입니다.\n", num)

	fmt.Printf("num의 메모리 주소 출력: %p\n", &num) //주솟값을 참조하기 위해 &를 쓴다.
	fmt.Printf("num의 유니코드 값은 %U입니다.\n", num)
	fmt.Printf("num의 타입은 %T입니다.\n", num)
	fmt.Printf("num의 타입은 %T입니다.\n", num)
	fmt.Printf("모든 형식으로 출력: %v, %v\n", 54.234, "Hello")
	fmt.Printf("num의 타입은 %T입니다.\n", num)
	fmt.Printf("7이 어떤 형식인지 표시: %d, %#o, %#x\n", 7, 7, 7) //8진수는 앞에 0이 붙고, 16진수는 0x가 붙습니다.
	fmt.Printf("네 칸 차지하는 13: %4d\n", 13)
	fmt.Printf("빈칸은 0으로 채우고 4칸 차지하는 13: %04d\n", 13)
	fmt.Printf("총 네 칸 차지하고 왼쪽으로 정렬되는 13과 15: %-4d%-4d\n", 13, 15)
	fmt.Printf("12.1234를 소수점 둘째 자리까지만 표시하면 %.2f입니다.\n", 12.1234)

}
```

출력결과

```
5>6=%!b(bool=false)
15는 2진수로 1111
저의 성은 김 입니다
19는 10진수로 19입니다.
19는 8진수로 23입니다.
19는 16진수로 13입니다.
19는 16진수로 13입니다.
19.1234는 고정 소수점으로 19.123400입니다.
19.1234는 고정 소수점으로 19.123400입니다.
19.1234의 지수 표현은 1.912340e+01입니다.
19.1234의 지수 표현은 1.912340E+01입니다.
19.1234의 간단한 실수 표현은 19.1234입니다.
19.1234의 간단한 실수 표현은 19.1234입니다.
문자열: 안녕하세요.
num은 50입니다.
num의 메모리 주소 출력: 0xc82000a348
num의 유니코드 값은 U+0032입니다.
num의 타입은 int입니다.
num의 타입은 int입니다.
모든 형식으로 출력: 54.234, Hello
num의 타입은 int입니다.
7이 어떤 형식인지 표시: 7, 07, 0x7
네 칸 차지하는 13:   13
빈칸은 0으로 채우고 4칸 차지하는 13: 0013
총 네 칸 차지하고 왼쪽으로 정렬되는 13과 15: 13  15
12.1234를 소수점 둘째 자리까지만 표시하면 12.12입니다.
```

---

## 콘솔 입력 함수(Scan)

- c언어에선 `scanf`, java에선 `new Scanner(system.in)`
- Go에선 c언어와 크게 다르지 않음
- `Scanln`, `Scan`, `Scanf`가 존재

### Scanln, Scan, Scanf

| 선언   | 입력 형태                                          |
| ------ | -------------------------------------------------- |
| Scanln | 공백으로 구분하여 입력                             |
| Scan   | 공백과 개행으로 구분하여 입력                      |
| Scanf  | 포맷 지정자를 이용하여 개발자가 원하는 형태로 입력 |

### Scanln 예제

```js
package main

import "fmt"

func main() {
	var name string
	var gen string
	var school string

	fmt.Print("이름, 성별, 학교를 입력하세요.")
	fmt.Scanln(&name, &gen, &school)

	fmt.Println("이름은 ", name, " 성별은 ", gen, "학교는", school)
}
```

입력값

```
홍길동 남자 홍익대
```

### Scan 예제

```js
package main

import "fmt"

func main() {
	var name string
	var gen string
	var school string

	fmt.Print("이름, 성별, 학교를 입력하세요.")
	fmt.Scan(&name, &gen, &school)

	fmt.Println("이름은 ", name, " 성별은 ", gen, "학교는", school)
}
```

입력값

```
홍길동
남자
홍익대
```

### Scanf 예제

```js
package main

import "fmt"

func main() {
	var i, j int

	fmt.Print("주민등록번호를 -를 이용해 입력하세요 :")
	fmt.Scanf("%d-%d", &i, &j)

	fmt.Printf("주민등록번호는 %d-%d\n", i, j)
	fmt.Println("앞자리는", i)
	fmt.Println("뒷자리는", j)
}
```

입력값

```
960822-1111111
```
