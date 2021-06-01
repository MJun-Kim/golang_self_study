# Chapter 18 Channel(채널)

- 고루틴 간에 서로 값을 주고 받는 통로가 Channel !
- 데이터를 주고 받을때 까지 고루틴을 종료하지 않음 -> 별도의 lock 없이 데이터 동기화

   기본적인 사용법
    - "make(chan 데이터타입)" 형식으로 생성
    - 채널의 데이터 송/수신은 '<-' 연산자 이용
    - 체널에 값을 보낼 때는 채널 <- 데이터, 채널에서 값을 받을 때는 <- 채널
    - 값을 받을 때는 :=이나 =을 이용해 변수에 바로 값을 대입 가능
    - 채널에서 값을 받을 때까지만 대기, 가져오면 바로 다음 코드를 실행

~~~go
package main

import "fmt"

func main() {
	var str = "Hello Goorm!"
	done := make(chan bool)
	
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(str, i)	
		}
		
		done <- true //채널에 true를 송신함
	}()

	<- done //수신함으로써 대기를 끝냄
}
~~~

---

## 비동기 체널과 버퍼

- Deadlock (데드락) : 둘 이상의 프로세스(함수)가 서로 가진 한정된 자원을 요청하는 경우 발생하는 것

ex ) 채널에 송신했는데 수신자가 없거나 수신자가 있으나 송신하는 데이터가 없다? -> Deadlock 발생

~~~go
package main
 
import "fmt"
 
func main() {
	c := make(chan string)
	
	c <- "Hello goorm!"
	
	fmt.Println(<-c)
}
~~~

---

## 비동기 채널 버퍼

- 채널은 송/수신자 1대1일 대응! -> 이를 편하게 하기 위한 중재자 버퍼
- 중간에 데이터를 잠깐 저장하는 공간으로 생각
- 송신자는 모두 송신하면 종료, 마찬가지로 수신자도 모두 수신하면 송신 루틴에 상관없이 자신의 일을 수행하고 종료

~~~go
package main

import (
	"fmt"
)

func main() {
	done := make(chan bool, 6)

	go func() {
		for i := 0; i < 6; i++ {
			done <- true
			fmt.Println("고루틴 : ", i)
		}
	}()

	for i := 0; i < 6; i++ {
		<-done
		fmt.Println("메인 함수 : ", i)
	}

}

~~~
---

## 동기 채널

- 처음 사용했던 그대로!
- 동기 채널은 단순히 송/수신 채널이 여러개여서 송신 루틴과 수신 루틴이 번갈아가면서 실행되는 것


---

## 채널 닫기

- 루틴에서 채널을 생성하고 통신하려면 수진 부위가 명확해야함
- 송신 데이터가 없는 수신시에는 무한 대기 상태가 발생 -> Deadlock!<br>
    -> 이를 방지하고자 채널을 닫으면 수신은 가능 -> 오류 발생 x

- 모두 수신후 close된 채널 수신하면 nil 값
- close된 채널에 송신하면 "send on closed channel" panic 발생
- <- 수신자 에서 반환값은 2개 : (수신값) , (개폐 여부)


~~~go
package main
 
import "fmt"

func main() {
	c := make(chan string, 2)

	c <- "Hello"
	c <- "goorm"

	close(c)

	val, open := <- c
	fmt.Println(val, open)
	val, open = <- c
	fmt.Println(val, open)
	val, open = <- c
	fmt.Println(val, open)
	val, open = <- c
	fmt.Println(val, open)
}
~~~


### 채널 range 문


- 채널의 개패 여부를 이용 하여 for 문을 이용 하여 송/수신 개수 맞춰사용
- range문 이용시 간략하게 작성가능

*close안할시에 에러(Deadlock)!
for 만 이용시
~~~go
package main
 
import "fmt"

func main() {
	c := make(chan int, 10)

	for i := 0; i<10; i++ {
		c <- i
	}    
	close(c)
	
	for {
		if val, open := <-c; open { // 표현식; 조건식 형태
			// open이 true면 실행
			fmt.Println(val, open)
		} else {
			break
		}
	}
}
~~~

range문 이용시
~~~go
package main
 
import "fmt"

func main() {
	c := make(chan int, 10)

	for i := 0; i<10; i++ {
		c <- i
	}
	close(c)
	
	for val := range c { // <- c를 사용하지 않음
		fmt.Println(val)
	}
}
~~~

---

## 송진 전용, 수신 전용 채널

- 특별한 선언이 없으면 채널은 송신, 수신에 자유로움
- 채널을 매개변수로 이용시 "(매개변수이름 chan 채널테이터타입)"
- 매개변수 선언시 설정 or 반환 형식 선언
- "chan<- 채널데이터타입" - 송신 전용 선언
- "<-chan 채널데이터타입" - 수신 전용 선언

~~~go
package main

import "fmt"

func main() {
	numsch := num(10, 5) 
	result := sum(numsch)  
	//채널 result는 수신만 할 수 있음
	fmt.Println(<-result) 
}

func num(num1, num2 int) <-chan int {
	numch := make(chan int)
	
	go func() {
		numch <- num1
		numch <- num2 //송신 후
		close(numch) 
	}()
	
	return numch //수신 전용 채널로 반환
}

func sum(c <-chan int) <-chan int {
	//채널 c는 수신만 할 수 있음
	sumch := make(chan int)
	
	go func() {
		r := 0
		for i := range c { //채널 c로부터 수신
			r = r + i  
		}
		sumch <- r // 송신 후
	}()
	
	return sumch //수신 전용 채널로 반환
}
~~~


# 채널 select문

- switch문과 비슷하게 생김
- 수신자가 여러개 있으면 동시에 여러개를 기다리는 형태?
- case 에 송신자가 있을시에는 다른 case에 수신이 없을시에 송신 수행

~~~go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan string)
	
	go func() {
		for {
			time.Sleep(200 * time.Millisecond)
			c := <- ch3
			fmt.Printf("ch3 데이터 %s 수신\n", c)
		}
	}()

	go func() {
		for {
			time.Sleep(1000 * time.Millisecond)
			ch1 <- 10
		}
	}()

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			ch2 <- 20
		}
	}()

	go func() {
		for {
			select {
				case a := <-ch1:
				fmt.Printf("ch1 데이터 %d 수신\n", a)
				case b := <-ch2:				
				fmt.Printf("ch2 데이터 %d 수신\n", b)
				case ch3 <- "goorm":
			}
		}
	}()
	
	time.Sleep(5 * time.Second)
}
~~~
---

# 연습문제
1. 고루틴 실습2

~~~go
package main

import "fmt"

func add(ch chan int, num1 int, num2 int) {
	ch <- num1 + num2
}

func main() {
	var num1, num2 int
	ch := make(chan int)

	fmt.Scanln(&num1, &num2)

	go add(ch, num1, num2)

	fmt.Println(<-ch)
}
~~~

2. 메세지 전송

- 10초안에 입력 받고 메세지 전송
~~~go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go sendMessage(ch)

	for leftTime := 10; leftTime > 0; leftTime-- {
		select {
		case message := <-ch:
			fmt.Printf("%s 메세지가 발송되었습니다.\n", message)
			return
		default:

			go func() {
				fmt.Printf("%d 초 안에 입력하세요.\n", leftTime)
			}()
		}

		time.Sleep(time.Second)

	}
}

func sendMessage(ch chan string) {
	var test string
	fmt.Scan(&test)

	ch <- test
}
~~~

3. 동기 체널 실습
- "송신 루틴 완료" 텍스트 보기
~~~go
package main

import (
	"fmt"
	"time"
)

func main() {
	var done chan bool = make(chan bool)
	var channalTest chan int

	fmt.Printf("%T\n", done)
	fmt.Printf("%T\n", channalTest)

	go func() {
		for i := 0; i < 20; i++ {
			done <- true
		}
		fmt.Println("송신 루틴 완료")
	}()

	for i := 0; i < 20; i++ {
		<-done
		fmt.Println("수신한 데이터 :", i)
	}

	time.Sleep(time.Second * 3)

}
~~~


4. 비동기 채널 실습

~~~go
package main

import "fmt"

func main() {
	
	ch := make(chan bool, 50)

	go func() {
		for i := 0 ; i < 20 ; i ++{
			ch <- true
		}
		
		fmt.Println("송신 루틴 완료")
	}()
		
	
	for i := 0 ; i < 20 ; i++{
		<- ch
		fmt.Println("수신한 데이터 : ", i)
	}
}
~~~