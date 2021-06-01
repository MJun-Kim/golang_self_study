package main

import "fmt"

func main() {
	c := make([]string, 0, 3) //용량이 3이고 길이가0인 정수형 슬라이스 선언
	// s := make([]int, 0, 3) //용량이 3이고 길이가0인 정수형 슬라이스 선언
	c = append(c, "test", "test", "test", "test", "test", "test", "test")
	// c = append(c, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(len(c), cap(c))

	fmt.Printf("  ")
	// c = append(c, 1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7)
	// fmt.Println(len(c), cap(c))

	// for i := 1; i <= 7; i++ { // 1부터 차례대로 한 요소씩 추가
	// 	c = append(c, i)

	// 	fmt.Println(len(c), cap(c)) // 슬라이스 길이와 용량 확인
	// }

	// l := c[1:3] //인덱스 1요소부터 2요소까지 복사
	// fmt.Println(l)

	// l = c[2:] //인덱스 2요소부터 끝까지 복사
	// fmt.Println(l)

	// l[0] = 6

	// fmt.Println(c) //슬라이스 l의 값을 바꿨는데 c의 값도 바뀜
	// //값을 복사해온 것이 아니라 기존 슬라이스 주솟값을 참조
}

// package main

// import "fmt"

// func main() {
// 	var a map[int]string

// 	if a == nil {
// 		fmt.Println("nil map")
// 	}

// 	var m = map[string]string{ //key:value, 형식으로 초기화한다
// 		"apple":  "red",
// 		"grape":  "purple",
// 		"banana": "yellow",
// 	}

// 	fmt.Println(m, "\nm의 길이는", m["apple"])
// }
