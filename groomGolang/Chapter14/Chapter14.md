# Chapter 14 인터페이스 Interface

- structure -> 변수를 모아둔것
- 아주 간략하게! -> 메소드를 모아둔것!
- 매개변수로 인터페이스를 사용한다는 것은 구조체에 관계없이 인터페이스에 포함된 메소드를 사용하겠다는 뜻

```go
package main

import (
	"fmt"
	"math"
)

type geometry interface { //인터페이스 선언 Reat와 Circle 메도스의 area를 모두 포함
	area() float64
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}


func main() {
	r1 := Rect{10, 20}
	c1 := Circle{10}
	r2 := Rect{12, 14}
	c2 := Circle{5}

	printMeasure(r1, c1, r2, c2)
}

func printMeasure(m ...geometry) { //인터페이스를 가변 인자로 하는 함수
	for _, val := range m { //가변 인자 함수의 값은 슬라이스형
		fmt.Println(val.area()) //인터페이스의 메소드 호출
	}
}
```

## 빈 인터페이스 ( Empty Interfaece )

- 인터페이스는 내용을 선언하지 않아도 사용할 가능
- 하나의 형( 변수 타입 )으로 매개변수로 사용 가능
- 어떠한 타입도 담을 수 있는 Dynamic Type ( 굳이 해석하자면 동적 타입 )

ex )

```go
package main

import "fmt"

func printVal(i interface{}) {
    fmt.Println(i)
}

func main() {
    var x interface{} //빈 인터페이스 선언

	x = 1           //정수를 넣었다가
	printVal(x)

	x = "test"      //문자열을 넣었다가 like javascript!
    printVal(x)
}
```

## Type Assertion

- Dynamic Type 이기 떄문에 확실한 형을 명시할때 사용
- 인터페이스 형으로 선언 후 nil값이면 에러 발생

```go
package main

import "fmt"

func main() {
    var num interface{} = 10

    a := num
    b := num.(int)
    c := num

	fmt.Printf("%T,%d\n",a,a)
    printtest(b)
    printtest(c)
}

func printtest (i interface{}){
	fmt.Printf("%T,%d\n",i,i)
}
```

## 예제문제

```go
package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float32
	volume() float32
}

type Cylinder struct {
	r float32
	h float32
}

type Cuboid struct {
	a float32
	b float32
	c float32
}

func (c Cylinder) volume() float32 {
	return c.r * c.r * math.Pi * c.h
}

func (c Cuboid) volume() float32 {
	return c.a * c.b * c.c
}

func (c Cylinder) area() float32 {
	return c.r*c.r*math.Pi*2.0 + 2.0*math.Pi*c.r*c.h
}

func (c Cuboid) area() float32 {
	return 2*c.a*c.b + 2*c.b*c.c + 2*c.c*c.a
}

func main() {

	cy1 := Cylinder{10, 10}
	cy2 := Cylinder{4.2, 15.6}
	cu1 := Cuboid{10.5, 20.2, 20}
	cu2 := Cuboid{4, 10, 23}

	printMeasure(cy1, cy2, cu1, cu2)
}

func printMeasure(g ...geometry) {
	for _, value := range g {
		fmt.Printf("%.2f, %.2f\n", value.area(), value.volume())
	}
}

```
