## CH3. 연산자

---

## 연산자 종류

- 변수와 상수는 값을 저장하는 기능을 하고 저장한 값(저장하지 않은 값)을 가지고 연산을 하기 위해 연산자 사용
- 다른 언어들과 거의 비슷하지만 몇 가지 다른점이 있으니 천천히 살펴보면 좋음

<br>

### 3.1 수식 연산자

- 두 개의 피연산자를 요구하는 이항 연산자(binary operator)
- 다른언어들과 똑같이 사칙연산, 나머지값 반환하는 연산자 존재

| 종류 | 기능                                           | example    |
| ---- | ---------------------------------------------- | ---------- |
| +    | 피연산자의 값을 더한다                         | 1 + 2 = 3  |
| -    | 왼쪽 피연산자 값에 오른쪽 피연산자 값을 뺀다   | 1 - 2 = -1 |
| \*   | 피연산자의 값을 곱한다                         | 2 - 1 = 1  |
| /    | 왼쪽 피연산자 값에 오른쪽 피연산자 값을 나눈다 | 7 / 2 = 3  |
| %    | 피연산자의 값을 곱한다                         | 7 % 2 = 1  |

`주의할 점은 + 연산자는 문자열 결합도 가능` ex) "abc" + "def" = "abcdef"

```js
package main

import "fmt"

func main() {
	num1, num2 := 17, 5
	str1, str2 := "Hello", "goorm!"

	fmt.Println("num1 + num2 =", num1+num2) // 22
	fmt.Println("str1 + str2 =", str1+str2) // Hellogoorm!
	fmt.Println("num1 - num2 =", num1-num2) // 12
	fmt.Println("num1 * num2 =", num1*num2) // 85
	fmt.Println("num1 / num2 =", num1/num2) // 3
	fmt.Println("num1 % num2 =", num1%num2) // 2
}
```

<br>

### 3.2 증감 연산자

- 값을 1만큼 감소시키거나 증가시키는 연산자
- Go에서 증감연산자 사용시 주의사항 2가지
  - 증감연산자 사용과 동시에 대입 불가능. <b>ex) num := count++</b>
  - 전위 연산 불가능. <b>ex) ++count</b>

| 종류 | 기능               | example |
| ---- | ------------------ | ------- |
| ++   | 값을 1 증가시킨다. | count++ |
| --   | 값을 1 감소시킨다. | count-- |

```js
package main

import "fmt"

func main() {
	count1, count2 := 1, 10.4

	count1++
	count2--

	fmt.Println("count1++ :", count1) // 2
	fmt.Println("count2-- :", count2) // 9.4
}
```

<br>

### 3.3 할당 연산자

- 할당 연산자는 값을 단순히 대입하는 대입 연산자와 연산 후 값을 바로 대입시키는 복합 대입 연산자가 존재함

| 종류 | 기능                                              | example                               |
| ---- | ------------------------------------------------- | ------------------------------------- |
| =    | 변수나 상수에 값을 대입한다(변수끼리도 대입 가능) | a = 1                                 |
| :=   | 변수를 선언 및 대입한다                           | a := 1 &nbsp;<==>&nbsp; var a int = 1 |
| +=   | 값을 더한 후 대입한다                             | a += 1 &nbsp;<==>&nbsp; a = a + 1     |
| -=   | 값을 뺀 후 대입한다                               | a -= 1 &nbsp;<==>&nbsp; a = a - 1     |
| \*=  | 값을 곱한 후 대입한다                             | a \*= 1 &nbsp;<==>&nbsp; a = a \* 1   |
| /=   | 값을 나눈 후 몫을 대입한다                        | a /= 1 &nbsp;<==>&nbsp; a = a / 1     |
| %=   | 값을 나눈 후 나머지를 대입한다                    | a %= 1 &nbsp;<==>&nbsp; a = a % 1     |
| &=   | 값의 AND 비트 연산 후 나머지를 대입한다           | a &= 1 &nbsp;<==>&nbsp; a = a & 1     |
| \|=  | 값의 OR 비트 연산 후 나머지를 대입한다            | a \|= 1 &nbsp;<==>&nbsp; a = a \| 1   |
| ^=   | 값의 XOR 비트 연산 후 나머지를 대입한다           | a ^= 1 &nbsp;<==>&nbsp; a = a ^ 1     |
| &^=  | 값의 AND NOT 비트 연산 후 나머지를 대입한다       |                                       |
| <<=  | 비트를 왼쪽으로 이동 후 대입한다.                 |                                       |
| >>=  | 비트를 오른쪽으로 이동 후 대입한다.               |                                       |

```js
package main

import "fmt"

func main() {
	a := 2
	var num int

	num = a
	fmt.Println("num = a :", num) // num = a : 2
	num += 4
	fmt.Println("num += 4 :", num) // num += 4 : 6
	num -= 2
	fmt.Println("num -= 2 :", num) // num -= 2 : 4
	num *= 5
	fmt.Println("num *= 5 :", num) // num *= 5 : 20
	num /= 2
	fmt.Println("num /= 2 :", num) // num /= 2 : 10
	num %= 3
	fmt.Println("num %= 3 :", num) // num %= 3 : 1

	num = 3  //00000011
	num &= 2 //00000010
	fmt.Printf("num &= 2 : %08b, %d\n", num, num) // num &= 2 : 00000010, 2
	num |= 5 //00000101
	fmt.Printf("num |= 5 : %08b, %d\n", num, num) // num |= 5 : 00000111, 7
	num ^= 4 //00000100
	fmt.Printf("num ^= 4 : %08b, %d\n", num, num) // num ^= 4 : 00000011, 3
	num &^= 2 //00000010
	fmt.Printf("num &^= 2 : %08b, %d\n", num, num) // num &^= 2 : 00000001, 1
	num <<= 9 //00001001
	fmt.Printf("num <<= 9 : %08b, %d\n", num, num) // num <<= 9 : 1000000000, 512
	num >>= 8 //00001000
	fmt.Printf("num >>= 8 : %08b, %d\n", num, num) // num >>= 8 : 00000010, 2
}
```

<br>

### 3.4 논리 연산자

- 익숙하게 알고 있는 다른 언어들과 똑같음
- Go에서 논리 연산자 사용시 bool형의 선인 및 사용만이 가능함
  `var a int = 10, b := 1`일 때,

  - fmt.Println(!a) (X)
  - fmt.Println(!b) (X)

  | 종류 | 기능                                                           | example                                  |
  | ---- | -------------------------------------------------------------- | ---------------------------------------- |
  | &&   | 연산하는 A와 B모두 ‘참'이면 연산 결과로 ‘참'을 반환            | 1 && 0 = 0, 1 && 1 = 1, 0 && 0=0         |
  | \|\| | 연산하는 A와 B 둘 중 하나라도 ‘참'이면 연산 결과로 ‘참’을 반환 | 1 \|\| 0 = 1, 1 \|\| 1 = 1, 0 \|\| 0 = 0 |
  | !    | 연산하는 A가 ‘참'이면 ‘거짓', ‘거짓'이면 ‘참'을 반환           | !1 = 0, !0 = 1                           |

```js
package main

import "fmt"

func main() {
	var a bool = true
	b := false

	fmt.Println("0 && 0 : ", b && b) // 0 && 0 :  false
	fmt.Println("0 && 1 : ", b && a) // 0 && 1 :  false
	fmt.Println("1 && 1 : ", a && a) // 1 && 1 :  true
	fmt.Println("0 || 0 : ", b || b) // 0 || 0 :  false
	fmt.Println("0 || 1 : ", b || a) // 0 || 1 :  true
	fmt.Println("1 || 1 : ", a || a) // 1 || 1 :  true

	fmt.Println("!1 ", !true) // !1  false
	fmt.Println("!0 ", !false) // !0 true
}
```

<br>
7
### 3.5 관계 연산자

- 두 값의 대소와 동등의 관계를 따지는 연산자
- 연산 조건 만족 시 true, 만족하지 못할 시 false 반환

| 종류 | 기능                           |
| ---- | ------------------------------ |
| ==   | 두 값이 같은지 비교            |
| !=   | 두 값이 다른지 비교            |
| <    | 오른쪽 값이 큰지 비교          |
| <=   | 오른쪽 값이 크거나 같은지 비교 |
| >    | 왼쪽 값이 큰지 비교            |
| >=   | 왼쪽 값이 크거나 같은지 비교   |

```js
package main

import "fmt"

func main() {
	fmt.Println("13 == 13 : ", 13 == 13) // 13 == 13 :  true
	fmt.Println("13 == 23 : ", 13 == 23) // 13 == 23 :  false
	fmt.Println("13 != 13 : ", 13 != 13) // 13 != 13 :  false
	fmt.Println("3 != 5 : ", 3 != 5) // 3 != 5 :  true
	fmt.Println("0 < 1 : ", 0 < 1) // 0 < 1 :  true
	fmt.Println("0 > 1 : ", 0 > 1) // 0 > 1 :  false
	fmt.Println("0 >= 1 : ", 0 >= 1) // 0 >= 1 :  false
	fmt.Println("0 <= 1 : ", 0 <= 1) // 0 <= 1 :  true
}
```

<br>

### 3.6 비트 연산자

- 비트 단위의 연산을 진행하는 연산자
- 기계에 좀 더 친화적인 연산자지만 다른 영역에도 사용돼 효율성을 높이고 연산의 수를 줄이는 요인이 되기도 함
- 활용적 측면을 이해하기엔 큰 부담이 따르기 때문에 연산자 자체의 기능을 이해하는데 초점을 맞추어 진행

| 종류 | 기능                             |
| ---- | -------------------------------- |
| &    | 두 값을 비트 단위로 AND연산      |
| \|   | 두 값을 비트 단위로 OR연산       |
| ^    | 두 값을 비트 단위로 XOR 연산     |
| &^   | 두 값을 비트 단위로 AND NOT 연산 |
| <<   | 값의 비트 열을 왼쪽으로 이동     |
| >>   | 값의 비트 열을 오른쪽으로 이동   |

```js
package main

import "fmt"

func main() {
	num1 := 15 //00001111
	num2 := 20 //00010100

	fmt.Printf("num1 & num2 : %08b, %d\n", num1&num2, num1&num2) // num1 & num2 : 00000100, 4
	fmt.Printf("num1 | num2 : %08b, %d\n", num1|num2, num1|num2) // num1 | num2 : 00011111, 31
	fmt.Printf("num1 ^ num2 : %08b, %d\n", num1^num2, num1^num2) // num1 ^ num2 : 00011011, 27
	fmt.Printf("num1 &^ num2 : %08b, %d\n", num1&^num2, num1&^num2) // num1 &^ num2 : 00001011, 11

	fmt.Printf("num1 << 4 : %08b, %d\n", num1<<4, num1<<4) // num1 << 4 : 11110000, 240
	fmt.Printf("num2 >> 2 : %08b, %d\n", num2>>2, num2>>2) // num2 >> 2 : 00000101, 5
}
```

<br>

### 3.7 채널 연산자

- 채널을 사용할 때 쓰는 연산자
- Go에는 채널이라는 개념이 존재하는데 간단히 말해 채널이랑 고루틴끼리 데이터를 주고 받고 실행 흐름을 제어하는 기능
- 채널과 고루틴관련은 관련챕터가서 상세히 다룰예정

| 종류 | 기능                   | 설명                        |
| ---- | ---------------------- | --------------------------- |
| <-   | 채널의 수신을 연산한다 | 채널에 값을 보내거나 가져옴 |

```js
package main

import "fmt"

func main() {
	ch := make(chan int) //정수형 채널 생성

	go func() {
		ch <- 10
	}() //채널에 10을 보냄

	result := <-ch //채널로부터 10을 전달받음
	fmt.Println(result) // 10
}
```

<br>

### 3.8 포인터 연산자

- C, C++언어와 같이 &, \*연산자를 통해 메모리에 접근할 수 있음
- Go언어는 포인터 연산자는 제공하지만 포인터 산술 즉, `포인터에 더하고 빼는 기능은 제공하지 않음`

| 종류 | 기능                                               |
| ---- | -------------------------------------------------- |
| &    | 변수의 메모리 주소를 참조함                        |
| \*   | 포인터 변수에 저장된 메모리에 접근하여 값을 참조함 |

```js
package main

import "fmt"

func main() {
	var num int = 5
	var pnum = &num

	fmt.Println("num : ", num)   // num :  5
	fmt.Println("pnum :", pnum)  // pnum : 0xc82000a2d0, num의 메모리 주소
	fmt.Println("pnum :", *pnum) // pnum : 5, num의 주소로 메모리에 할당돼있는 값 접근

	*pnum++
	fmt.Println("num : ", num) // num :  6
	fmt.Println("pnum :", *pnum) // pnum : 6
	//포인터 연산자를 이용한 값 변경
}
```

---

## 연산자 우선순위

- 한 수식안에 여러 연산자가 들어가있을 경우 어떤 연산자를 우선적으로 처리해야되는지에 대한 순위
- 간단한 예로, `덧셈, 뺄셈보다 곱셈, 나눗셈 먼저 계산해야한다`와 같은 것이 <b>연산자 우선순위</b>
- 같은 순위의 연산자를 어느방향에서부터 처리해야하는지를 나타내는 것이 결합방향

<table>
    <th>순위</th>
    <th>연산기호</th>
    <th>연산자</th>
    <th>결합방향</th>
    <tr>
        <td rowspan="4">1</td>
        <td>()</td>
        <td>함수호출</td>
        <td rowspan="4">--></td>
    </tr>
    <tr>
        <td>[]</td>
        <td>인덱스</td>
    </tr>
    <tr>
        <td>-></td>
        <td>간접지정</td>
    </tr>
    <tr>
        <td>++, --</td>
        <td>증가 및 감소</td>
    </tr>
    <tr>
        <td rowspan="7">2</td>
        <td>+, -</td>
        <td>부호 연산(음수와 양수의 표현)</td>
        <td rowspan="7"><--</td>
    </tr>
    <tr>
        <td>!</td>
        <td>논리 NOT</td>
    </tr>
    <tr>
        <td>~</td>
        <td>비트 단위 NOT</td>
    </tr>
    <tr>
        <td>(type)</td>
        <td>타입 변환</td>
    </tr>
    <tr>
        <td>*</td>
        <td>간접 지정 연산</td>
    </tr>
    <tr>
        <td>&</td>
        <td>주소 연산</td>
    </tr>
    <tr>
        <td>sizeof</td>
        <td>바이트 단위 크기 계산</td>
    </tr>
    <tr>
        <td>3</td>
        <td>*, /, %</td>
        <td>곱셈, 나눗셈 관련 연산</td>
        <td>--></td>
    </tr>
    <tr>
        <td>4</td>
        <td>+, 0</td>
        <td>덧셈, 뺄셈</td>
        <td>--></td>
    </tr>
    <tr>
        <td>5</td>
        <td><<, >></td>
        <td>비트 이동</td>
        <td>--></td>
    </tr>
    <tr>
        <td>6</td>
        <td><, <=, >, >=</td>
        <td>대소 비교</td>
        <td>--></td>
    </tr>
    <tr>
        <td>7</td>
        <td>==, !=</td>
        <td>동등 비교</td>
        <td>--></td>
    </tr>
    <tr>
        <td>8</td>
        <td>&</td>
        <td>비트 AND</td>
        <td>--></td>
    </tr>
    <tr>
        <td>9</td>
        <td>^</td>
        <td>비트 XOR</td>
        <td>--></td>
    </tr>
    <tr>
        <td>10</td>
        <td>|</td>
        <td>비트 OR</td>
        <td>--></td>
    </tr>
    <tr>
        <td>11</td>
        <td>&&</td>
        <td>논리 AND</td>
        <td>--></td>
    </tr>
    <tr>
        <td>12</td>
        <td>||</td>
        <td>논리 OR</td>
        <td>--></td>
    </tr>
    <tr>
        <td>13</td>
        <td>? :</td>
        <td>조건 연산</td>
        <td><--</td>
    </tr>
    <tr>
        <td>14</td>
        <td>=, +=, *=, ...</td>
        <td>대입 연산</td>
        <td><--</td>
    </tr>
    <tr>
        <td>15</td>
        <td>,</td>
        <td>콤마 연산</td>
        <td>--></td>
    </tr>
</table>

---

## 콘솔 입력 함수의 기본

- 콘솔 입력 함수를 이용하면 프로그램 사용자가 값을 입력할 수 있음
- fmt패키지를 이용한 콘솔 입력 함수에는 `Scanf`, `Scan`, `Scanln` 등이 있음
- `Scanln`은 여러 값을 동시에 받을 수 있음. `빈칸(스페이스바)으로 값을 구분하고 엔터(개행)를 입력하면 입력이 종료. 입력받는 변수에 '&' 연산자를 붙여 입력받음`
- "콘솔 출력과 입력 함수"에서 자세히 다룰 예정

```js
package main

import "fmt"

func main() {
	var num1, num2, num3 int

	fmt.Print("정수 3개를 입력하세요 :")
	fmt.Scanln(&num1, &num2, &num3) // 사용자가 3개의 정수를 스페이스바로 구분하여 입력
	fmt.Println(num1, num2, num3) // 사용자가 입력받은 3개의 정수 그대로 출력
}
```
