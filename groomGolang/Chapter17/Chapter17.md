# 고루틴(Goroutine)

---

## 비동기 프로세스의 기본

- `고루틴`은 여러 함수를 동시에 실행할 수 있는 논리적 가상 스레드.
- `멀티 프로세스`는 동시에 여러 프로그램을 실행하는 것
- `멀티 스레드`는 프로그램 안에서 다양한 기능을 동시에 실행하는 것입니다.

### 고랭에서의 비동기 프로세스

자바와 같은 언어에서는 함수를 호출할 때 멀티스레드를 사용하지만, Go언어에서는 스레드보다 훨씬 가벼운 비동기 동시 처리를 구현해 각각의 일에 대해 스레드와 1대 1로 대응하지 않고, `훨씬 적은 스레드를 사용함`

- 메모리 측면에서 스레드가 1MB의 스택을 갖을 때, 고루틴은 KB 단위의 훨씬 작은 스택을 가지고 필요시에 동적으로 증가 (이는 굉장히 효율적이고 가벼운 기능으로서 비동기 프로세스를 구현할 때 Go언어의 장점이 극대화)

> 고 언어에서는 함수에 고루틴을 선언함으로써 함수를 비동기적으로 실행할 수 있으며, 문맥에서 알 수 있듯이 '비동기'적이라는 말은 한 번에 여러 일을 실행함을 의미

#### 비동기로 실행되는 간단한 예제

```go
package main

import "fmt"

func testGo() {
	fmt.Println("Hello goorm!")
}

func main() {
	go testGo() //고루틴으로 비동기 실행
    fmt.Scanln()
}
```

실행결과

```
Hello groom!
입력대기...
```

#### 고루틴을 이용한 난수 생성 예제

난수를 생성하기 위해 "math/rand" 패키지, 시간 출력을 위해 "time" 패키지를 import

- rand.Intn()는 정수형 난수를 생성하는 함수
- `time`에서 쓰이는 시간은 모두 Duration형. 이는 `time` 패키지 안에 선언된 구조체로서 int64형
- time.Sleep()은 프로그램에 대기 시간을 주는 함수. 괄호 안은 Duration형을 써야함

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func hello(n int) {
	r := rand.Intn(3) // 0부터 3까지 난수 생성
	time.Sleep(time.Duration(r) * time.Second)
	// 난수를 Dration형으로 형변환 후 second로 계산
	fmt.Println(n)
}

func main() {
	for i := 0; i < 100; i++ {
		go hello(i)        // 고루틴 100개 생성 비동기 실행
        // hello(i)        // 동기 실행(주어진 일을 순서대로 실행)
	}

	fmt.Scanln()
}
```

<br />

## 고루틴의 활용

고루틴이 끝날때까지 대기하는 기능 `sync` 패키지의 `WaitGroup`

> `WaitGroup`은 `sync`패키지에 선언되어있는 구조체로서 고루틴이 완료될 때까지 대기. 이를 변수로 선언해 메서드로 활용 가능

- Add() : 기다릴 고루틴의 수 설정
- Done() : 고루틴이 실행된 함수 내에서 호출함으로써 함수 호출이 완료되었음을 알림
- Wait() : 고루틴이 모두 끝날때까지 차단

#### `WaitGroup`을 이용한 예제

```go
package main

import (
    "fmt"
    "math/rand"
    "time"
    "sync"
)

func hello(n int, w *sync.WaitGroup) {
    defer w.Done() //끝났음을 전달

    r := rand.Intn(3)

    time.Sleep(time.Duration(r) * time.Second)

    fmt.Println(n)
}

func main() {
    wait := new(sync.WaitGroup) //waitgroup 생성

    wait.Add(100) // 100개의 고루틴을 기다림

    for i := 0; i < 100; i++ {
            go hello(i, wait) //wait을 매개변수로 전달
    }

    wait.Wait() // 고루틴이 모두 끝날때까지 대기
}
```

- 위 예시 코드를 보면 `WaitGroup 생성` - `고루틴 개수 설정` - `끝났음을 전달` - `모두 끝날 때까지 대기(비동기라 같이 진행됨)` - `종료` 순서
- `main()` 함수에서 `new` 키워드를 사용해 `WaitGroup`의 `wait 포인터 변수`를 생성

  > new 키워드로 선언한 변수는 포인터형. 따라서 매개변수로 사용할 때 `&` 연산자를 사용하지 않아도 자동으로 주소를 참조. 일반적으로 var wait sync.WaitGroup 형식으로도 선언 가능하지만 이는 그냥 변수이므로 함수의 매개변수로 전달할 때는 주소를 참조하기 위해 `&` 연산자를 사용해야 함.

- wait.Add(100) 메소드를 호출해 고루틴 100개를 기다리도록 함. <b style="color: red">여기서 중요한 것은 wait은 포인터 변수이기 때문에 매개변수형을 _sync.WaitGroup 형식으로 선언함 </b> 만약 `_` 연산자를 넣지 않으면 hello() 함수 내에 w와 main() 함수의 wait이 다르게 인식되어(call by value) 고루틴이 모두 종료되지 않고 교착상태에 빠지게 됨. 따라서 Call by reference 형식으로 함수 내에서 w.Done() 메소드를 호출하고 고루틴이 모두 종료됐으면 main() 함수에서 호출되어 기다리고있던 wait.Wait() 메소드도 대기를 멈추고 종료

<br />

## 클로저에서의 고루틴

사실 WaitGroup은 클로저에서 많이 활용됨. 위 예시를 보았다시피 wait 변수를 포인터 변수로 선언해 사용하면서 좀 복잡했음. `익명함수 클로저`를 사용하면 클로저를 감싸고있는 함수 내에 선언된 wait을 직접 접근할 수 있기때문에 사용하기 편리

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wait sync.WaitGroup
	wait.Add(102)

	str := "goorm!"

	go func() {
		defer wait.Done()
		fmt.Println("Hello")
	}()

	go func() {
		defer wait.Done()
		fmt.Println(str)
	}()

	for i := 0; i < 100; i++ {
		go func(n int) {
			defer wait.Done()

			fmt.Println(n)
		}(i)
	}

	wait.Wait()
}
```

실행결과

```
Hello
goorm!
22
9
10
1
...
...
```

<br />

## 다중 CPU 병렬 처리

- 고루틴을 아무리 많이 만들어도 한 개의 CPU에서 시분할 처리
- 그치만 요즘 디바이스들은 CPU가 복수개(듀얼 코어 이상)이기 때문에 Go 언어는 고루틴 뿐만이 아니라 다중 CPU를 이용한 병렬 처리를 지원
- 고루틴에서의 동시성(Concurrency)은 독립적으로 실행되는 기능들이고, 다중 CPU의 병렬 처리(Parallelism)는 계산들을 동시 실행하는 것
  > 동시성은 한 번에 많은 것들을 처리하고, 병렬 처리는 한 번에 많은 일을 하는 것에 관한 것

개념은 조금 헷갈릴 수 있어도 Go언어에서 다중 CPU를 사용하는 것은 굉장히 간단한데, `runtime` 패키지에서 제공하는 함수를 이용하면 됨

- `runtime.NumCPU()` : 현재 디바이스의 CPU 개수를 반환
- `runtime.GOMAXPROCS()` : 입력한 수만큼 (Logical)CPU 사용, 입력한 수가 1 미만일 때 현재 설정 값을 반환하고 설정 값은 바꾸지 않음

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//디바이스의 총 CPU 개수를 반환하고 그 값을 CPU 사용 값으로 설정
	fmt.Println(runtime.GOMAXPROCS(0))
	// 현재 설정값 출력, 1미만이기 때문에 설정값 바꾸지 않음
	var wait sync.WaitGroup
	wait.Add(100)

	for i := 0; i<100; i++ {
		go func(n int) {
			defer wait.Done()
			fmt.Println(n)
		}(i)
	}

	wait.Wait()
}
```

<br />

## 실습문제 1 - 고루틴 실습

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func add(a *int, b *int, r *int, w *sync.WaitGroup) int {
	defer w.Done()
	time.Sleep(time.Second)
	*r = *a + *b
	return *r
}


func main() {
	var num1, num2 int = 10, 5
	var result int
	wait := new(sync.WaitGroup)

	wait.Add(1)

	go add(&num1, &num2, &result, wait)


	wait.Wait()
	fmt.Println(result)
}
```

실행결과

```
15
```
