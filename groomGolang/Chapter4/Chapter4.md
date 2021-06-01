# Chapter4 자료형

- 어떤 데이터를 저장할 지 표현하는 것!

- 크게 부울린, 문자열, 정수형, 실수, 복소수, 기타

:= 이용시 정수형은 int로 실수는 float32로 자동 할당됨!

---

## Boolean 타입

- 참/거짓 할당 가능한 자료형
- 오로지 'true' 와 'false'
- 1byte로 할당

---

## Integer 정수형 타입

- uintptr : 포인터의 주소를 할당할 때 사용

<table class="tg">
<thead>
  <tr>
    <th class="tg-0lax">자료형</th>
    <th class="tg-0lax">선언</th>
    <th class="tg-0lax">크기(byte)  </th>
  </tr>
</thead>
<tbody>
  <tr>
    <td class="tg-0lax" rowspan="5">정수형(음수포함)</td>
    <td class="tg-0lax">int</td>
    <td class="tg-0lax">n비트 시스템에서 n비트</td>
  </tr>
  <tr>
    <td class="tg-0lax">int8</td>
    <td class="tg-0lax">1</td>
  </tr>
  <tr>
    <td class="tg-0lax">int16</td>
    <td class="tg-0lax">2</td>
  </tr>
  <tr>
    <td class="tg-0lax">int32</td>
    <td class="tg-0lax">4</td>
  </tr>
  <tr>
    <td class="tg-0lax">int64</td>
    <td class="tg-0lax">8</td>
  </tr>
  <tr>
    <td class="tg-0lax" rowspan="6">정수형(0, 양수)</td>
    <td class="tg-0lax">uint</td>
    <td class="tg-0lax">n비트 시스템에서 n비트</td>
  </tr>
  <tr>
    <td class="tg-0lax">uint8</td>
    <td class="tg-0lax">1</td>
  </tr>
  <tr>
    <td class="tg-0lax">uint16</td>
    <td class="tg-0lax">2</td>
  </tr>
  <tr>
    <td class="tg-0lax">uint32</td>
    <td class="tg-0lax">4</td>
  </tr>
  <tr>
    <td class="tg-0lax">uint64</td>
    <td class="tg-0lax">8</td>
  </tr>
  <tr>
    <td class="tg-0lax">uintptr</td>
    <td class="tg-0lax">8</td>
  </tr>
</tbody>
</table>

- 64bit 운영체제??? 그래서 8byte!

---

## 실수 및 복소수 타입

- 복소수 선언은 3+4i 처럼 선언 할 수있다!

- 부동소수점수 -> 정확한 수는 표현하기 어려움<br>
  ex) 1.01 - 0.99 의 연산 결과 -> 0.02000000000000000018 != 0.02 근접하지만 일치하지는 않음.<br>

  따라서 연산한 값을 직접 비교하면 위험? 주의?

- 크기가 클수록 표현할 수 있는 숫자의 자리수가 커지기 때문에 정확도가 높아진다
<table class="tg">
<thead>
  <tr>
    <th class="tg-0pky">자료형</th>
    <th class="tg-0pky">선언</th>
    <th class="tg-0pky">크기(byte)</th>
  </tr>
</thead>
<tbody>
  <tr>
    <td class="tg-0pky" rowspan="2">실수</td>
    <td class="tg-0pky">float32</td>
    <td class="tg-0pky">4</td>
  </tr>
  <tr>
    <td class="tg-0pky">float64</td>
    <td class="tg-0pky">8</td>
  </tr>
  <tr>
    <td class="tg-0pky" rowspan="2">복소수</td>
    <td class="tg-0pky">complex64</td>
    <td class="tg-0pky">8</td>
  </tr>
  <tr>
    <td class="tg-0pky">complex128</td>
    <td class="tg-0pky">16</td>
  </tr>
</tbody>
</table>

---

## String 문자열 타입

- ""와 같이 비어있을 수 있고, 다른언어의 null과 Go에서 사용되는 nil이 아닐수 있다.

  > nil : 포인터의 zero value

- string으로 선언한 문자열 타입은 immutable 타입으로서 값을 수정 할 수 없다.

```go
var str string = "hello"
str[2] = 'a'    //오류!
```

---

## 기타 타입

- byte : 바이트 값을 8비트 부호없는 정수 값(int8)과 구별하는 데 사용. (byte 단위로 데이터를 처리할때)
- rune : 관례상 문자 값(string)을 정수 값과 구별하기 위해 사용. (유니코드를 표현할때)

---

## 문자열의 표현

2가지 방법!

- Back Quote(`)을 이용 ( Raw String Literal 이라고 부름)

```Go
var testStr string = `\n도 그냥 문자 취급해서 나옵니다.`
```

- 이중인용부호(쌍따움표?) (Interpreterd String Literal)

```Go
var testStr2 string = "개행문자로 나옴 그래서 한줄 띄워짐 \n"
```

두방법 모두 + 연산자로 결합해 표현 가능.

---

## 자료형의 변환

- 형변환에는 크게 2가지 종류가있는데
  '자동 형 변환'이 안되고
  '강제 형 변환'만된다!

```Go
package main

func main() {

    var integerNumber int = 10;
    var changeFloat float32 = float32(integerNumber)    //명시적 형변환
	var num1, num2 int = 3, 4

	var result float32 = num1 / num2    //자동 형 변환이 되지 않아 오류가 발생한다.
}
```

---

## 실습

```go
package main

import "fmt"

func main() {
var num1 , num2 ,num3 int;
fmt.Scanln( &num1 , &num2 , &num3);
fnum, unum, int64num := float32(num1) , uint( num2 ) , int64( num3 );
fmt.Printf( "%[1]T, %[1]f\n" , fnum) //%T 변수의 타입을 출력 '[n]' n번째 파라미터 출력
fmt.Printf( "%[1]T, %[1]d\n" , unum)
fmt.Printf( "%[1]T, %[1]d\n" , int64num)
}

```

출력결과

```
입력 > 1 -1 1
float32, 1.000000
uint, 18446744073709551615
int64, 1
```
