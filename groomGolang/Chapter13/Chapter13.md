# 구조체

---

- 하나 이상의 변수를 묶어서 새로운 자료형을 정의하는 `Custom data type`
- Go언어에서의 구조체는 필드들의 집합체이며 필드들의 컨테이너입니다
- `함수` 챕터에서 익명 함수를 배울 때 함수의 원형을 정의하기 위해 type문을 사용한 적이 있는데, 구조체도 이와 같은 `Custom data type`이기 때문에 `type`문을 사용해서 구조체를 정의

```js
type person struct {
	name string
	age int
	contact string
}
```

- 구조체는 선언 후 객체를 생성해서 사용할 수 있는데 지난 강의에서 배운 '컬렉션 - 슬라이스, 맵'과 선언 방식이 유사
  - 생성방법 `객체이름 := 구조체이름{저장할값}`
  - 빈 객체로 생성했다면 `객체이름.필드이름 = 저장할값`
  - `.(dot)`을 이용하면 필드 값을 저장하는 기능만 있는 것이 아니라 필드 값을 접근할 수 있는 용법

### 구조체 용법 예제

```js
package main

import "fmt"

type person struct {
	name    string
	age     int
	contact string
}

func main() {
	var p1 = person{}
	fmt.Println(p1)

	p1.name = "kim"
	p1.age = 25
	p1.contact = "01000000000"
	fmt.Println(p1)

	p2 := person{"nam", 31, "01022220000"} // 필드 이름을 생력할 시 순서대로 저장함
	fmt.Println(p2)

	p3 := person{contact: "01011110000", name: "park", age: 23} // 필드 이름을 명시할 시 순서와 상관 없이 저장할 수 있음
	fmt.Println(p3)

	p3.name = "ryu" //필드에 저장된 값을 수정할 수 있음
	fmt.Println(p3)

	fmt.Println(p3.contact) //필드 값의 개별 접근도 가능함
}
```

실행결과

```
{ 0 }
{kim 25 01000000000}
{nam 31 01022220000}
{park 23 01011110000}
{ryu 23 01011110000}
01011110000
```

- 초기화 할 때 값을 나열하면 구조체에 선언한 필드 순서대로 저장되지만 하지만 필드 이름에 값을 지정한다면 순서에 상관없이 해당 필드에 값이 저장
- Go언어의 구조체는 기본적으로 'mutable' 개체로서 필드 값이 변화할 경우 별도로 새 개체를 만들지 않고 해당 개체 메모리에서 직접 변경가능

<br />

## 구조체 포인터

- 함수(메소드)에서 같이 값을 복사해서 지역 변수로 사용하는 경우가 아니라 원래 값의 주소를 참조해 값이 저장된 주소에 직접 접근 하는 경우에 포인터를 썼음(매개변수에 '&'을 붙여서 Pass by reference를 한 것)
- 구조체도 마찬가지로 '구조체 포인터'를 생성할 수 있음. 두가지가 존재
  - `new(구조체이름)`을 사용하여 객체를 생성하기
  - 구조체 이름 앞에 `&` 붙이기
- 주의할 점은, 다른 자료형의 포인터들은 역참조를 위해 `*` 연산자를 사용했지만 `포인터 구조체는 선언하면 자동으로 역참조`. 따라서 함수 안에서 `*` 연산자를 사용할 필요가 없습니다.

### 포인터 구조체 예제

```js
package main

import "fmt"

type person struct {
	name    string
	age     int
	contact string
}

func addAgeRef(a *person) { //Pass by reference
	a.age += 4 //자동 역참조 * 생략
}

func addAgeVal(a person) { //Pass by value
	a.age += 4
}

func main() {
	var p1 = new(person) //포인터 구조체 객체 생성
	var p2 = person{}    // 빈 구조체 객체 생성

	fmt.Println(p1, p2)

	p1.age = 25
	p2.age = 25

	addAgeRef(p1) //&을 쓰지 않아도 됨
	addAgeVal(p2)

	fmt.Println(p1.age)
	fmt.Println(p2.age)

	addAgeRef(&p2) //&을 붙이면 pass by reference 가능
	fmt.Println(p2.age)
}
```

실행결과

```
&{ 0 } { 0 }
29
25
29
```

<br />

## 생성자(constructor) 함수

- 구조체를 사용하기 위해서는 우선 객체를 생성해야 사용할 수 있지만, 때로는 구조체의 필드 자체가 사용 전에 초기화되어야 하는 경우가 있음
- 예를 들어, 구조체의 필드가 'map' 형일 경우 구조체를 초기화할 때마다 맵 필드도 같이 초기화해야 하는 번거로움이 있을 수 있음
- 따라서 사전에 미리 초기화를 해 놓으면 외부 구조체 사용자가 매번 맵을 초기화해야 한다는 것을 기억할 필요가 없음
- 이러한 목적을 위해 '생성자 함수'를 사용. 생성자 함수는 호출하면 구조체 객체 생성 및 초기화, 입력한 필드 생성 및 초기화함과 동시에 구조체를 반환

```js
type mapStruct struct{ //맵 형태의 data필드를 가지는 "mapStruct"를 정의함
	data map[int]string
}

func newStruct() *mapStruct { //포인터 구조체를 반환함
	d := mapStruct{} //구조체 객체를 생성하고 초기화함
	d.data = map[int]string{} //data 필드의 맵을 초기화함
	return &d //초기화 한 포인터 구조체를 반환함
}
```

- 위 생성자 함수는 구조체 객체를 포인터와 함께 반환. 포인터 값이 없는 객체를 생성하는 생성자를 만들려면 반환형에 구조체 이름 앞에 붙은 포인터 연산자를 없애면 됨

### 생성자 함수를 이용한 구조체 예제

```js
package main

import "fmt"

type mapStruct struct {
	data map[int]string
}

func newStruct() *mapStruct { //포인터 구조체를 반환함
	d := mapStruct{}
	d.data = map[int]string{}
	return &d
}

func main() {
	s1 := newStruct() // 생성자 호출
	s1.data[1] = "hello"
	s1.data[10] = "world"

	fmt.Println(s1)

	s2 := mapStruct{}
	s2.data = map[int]string{}
	s2.data[1] = "hello"
	s2.data[10] = "world"

	fmt.Println(s2)
}
```

실행결과

```
&{map[1:hello 10:world]}
{map[1:hello 10:world]}
```

- 위 예시에서 `s1` 객체는 생성자 함수로 data 필드의 맵을 초기화했기 때문에 바로 data 필드에 값을 저장할 있지만, `s2` 객체는 구조체만 생성했기 때문에 data 필드에 값을 저장하기 위해 선언이 필요한 맵은 따로 초기화해야 됨
- 이렇게 생성자 함수를 사용하면 구조체의 사용이 훨씬 수월해질 수 있음

<br />

## 메소드

#### - 삼각형의 넓이를 계산하는 프로그램 (`함수`와 `구조체`를 활용)

```js
package main

import "fmt"

type triangle struct {
    width, height float32
}

func triArea(s *triangle) float32 { //'new'로 생성한 구조체 객체는 포인터값 반환
    return s.width * s.height / 2 //포인터 구조체는 자동 역참조 "*" 생략
}

func main() {
    tri1 := new(triangle)
    tri1.width = 12.5
    tri1.height = 5.2

    triarea := triArea(tri1)
    fmt.Printf("삼각형 너비:%.2f, 높이:%.2f 일때, 넓이:%.2f ", tri1.width, tri1.height, triarea)
}
```

```
삼각형 너비:12.50, 높이:5.20 일때, 넓이:32.50
```

- 구조체는 함수와 마찬가지로 만드는 데 기준과 목적이 있는데, 구조체도 변수가 많다는 이유로 단순히 값들을 묶지 않음
- 특정 속성들의 기능을 수행하기 위해 만들어진 특별한 함수를 `메소드` 라고 함
- `Java` 에서는 이들을 한 곳에 묶은 클래스 안에 필드와 메소드가 있지만 `Go언어`에서는 구조체 내에서 메소드를 선언하지 않고 일반 함수처럼 별도로 분리되어 선언
- 쉽게 말하자면 메소드는 구조체의 필드들을 이용해 특정 기능을 하는 특별한 함수. 특별한 함수인 만큼 선언 방법도 특이함
  - 기본적으로 메소드는 `func (매개변수이름 구조체이름) 메소드이름() 반환형 {` 형식으로 선언. 매개변수 이름은 구조체 변수명으로서 메소드 내에서 매개변수처럼 사용
  - `함수이름`은 입력하지 않고 구조체이름 뒤에 메소드 이름을 입력. 본문에서 메소드를 이용하기 위해 이름을 사용

### 메소드를 사용한 삼각형 넓이 계산 프로그램

```js
package main

import "fmt"

type triangle struct {
	width, height float32
}

func (s triangle) triArea() float32 { //value receiver
	return s.width * s.height / 2
}

func main() {
	tri1 := new(triangle)
	tri1.width = 12.5
	tri1.height = 5.2

	triarea := tri1.triArea()
	fmt.Printf("삼각형 너비:%.2f, 높이:%.2f 일때, 넓이:%.2f ", tri1.width, tri1.height, triarea)
}
```

실행결과

```
삼각형 너비:12.50, 높이:5.20 일때, 넓이:32.50
```

- 위 코드에서 (s triangle)은 어떤 구조체를 전달 받는지 명시하는 `receiver`. 구조체 객체 자체를 전달받는 것이 아니라 구조체 객체 정보를 전달 받고 메소드의 기능을 수행하는 것
- 물론 메소드에서도 값을 복사해서 받는 것이 아닌 `포인터 receiver`도 존재

<br />

## Value Receiver와 Pointer Receiver

- 메소드를 호출할 때는 다른 점이 없지만 메소드의 receiver 부분에서 주솟값을 참조하는 연산자인 `*`를 구조체 이름 앞에 붙여주면 됨

```js
package main

import "fmt"

type triangle struct {
	width, height float32
}

func (s triangle) triAreaVal() float32 { //Value Receiver
	s.width += 10
	return s.width * s.height / 2
}

func (s *triangle) triAreaRef() float32 { //Pointer Reciver
	s.width += 10
	return s.width * s.height / 2
}

func main() {
	tri1 := new(triangle)
	tri1.width = 12.5
	tri1.height = 5.2

	triarea_val := tri1.triAreaVal()
	fmt.Printf("삼각형 너비:%.2f, 높이:%.2f 일때, 넓이:%.2f \n", tri1.width, tri1.height, triarea_val)

	triarea_ref := tri1.triAreaRef()
	fmt.Printf("삼각형 너비:%.2f, 높이:%.2f 일때, 넓이:%.2f ", tri1.width, tri1.height, triarea_ref)
}
```

실행결과

```
삼각형 너비:12.50, 높이:5.20 일때, 넓이:58.50
삼각형 너비:22.50, 높이:5.20 일때, 넓이:58.50
```

<br />

### 실습문제 1 - 성적 저장 프로그램

```js
package main

import "fmt"

type student struct {
	name string
	gender string
	data map[string]int
}

func newStudent() *student {
	d := student{}
	d.data = map[string]int{}
	return &d
}

func main() {
	var studentNumber, subjectNumber, score int
	var name, gender, subject string

	fmt.Scanln(&studentNumber, &subjectNumber)

	s := make([]student, studentNumber)

	for i := 0 ; i < studentNumber ; i++ {
		object := newStudent()
		fmt.Scanln(&name, &gender)

		object.name = name
		object.gender = gender

		for j := 0 ; j < subjectNumber ; j++ {
			fmt.Scanln(&subject, &score)
			object.data[subject] = score
		}
		s[i] = *object
	}

	for i := 0 ; i < studentNumber ; i++ {
		fmt.Println("----------")
		fmt.Println(s[i].name, s[i].gender)

		for index, val := range s[i].data {
			fmt.Println(index, val)
		}

	}
	fmt.Println("----------")
}
```

입력값

```
54
kimF
kor57
math90
eng60
science50
leeF
kor78
math45
eng80
science100
yunM
kor67
math40
eng83
science97
namM
kor88
math91
eng80
science31
kimM
kor42
math29
eng100
science100
```

실행결과

```
----------
kimF
science50
kor57
math90
eng60
----------
leeF
kor78
math45
eng80
science100
----------
yunM
science97
kor67
math40
eng83
----------
namM
math91
eng80
science31
kor88
----------
kimM
kor42
math29
eng100
science100
----------
```

<br />

### 실습문제 2 - 역학적 에너지 2

```js
package main

import "fmt"

var g float32 = 9.8

type object struct {
	m, v, h, ke, pe float32
}

func (o object)Ep() float32 {
	return o.m * g * o.h
}

func (o object)Ek() float32 {
	return o.m*o.v*o.v/2
}


func main() {
	var n int
	var m, v, h float32
	objects := []object{}

	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&m, &v, &h)
		o := object{m, v, h, 0, 0}
		o.ke = o.Ek()
		o.pe = o.Ep()
		objects = append(objects, o)
	}

	for _, o := range objects {
		fmt.Println(o.ke, o.pe, o.ke+o.pe)
	}
}
```

입력값

```
5
30 2 10
10 2 40
4 5 1
30 2 33
19 21 9
```

실행결과

```
60 2940 3000
20 3920 3940
50 39.2 89.2
60 9702 9762
4189.5 1675.7999 5865.3
```
