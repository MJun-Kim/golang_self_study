package main

import (
	"fmt"
)

type item struct {
	name   string
	price  int
	amount int
}

type buyer struct {
	point          int
	shoppingBucket map[string]int
}

func newBuyer() *buyer {
	d := buyer{}
	d.point = 1000000
	d.shoppingBucket = map[string]int{}
	return &d
}

func main() {
	items := make([]item, 5)
	buyer := newBuyer()

	items[0] = item{"텀블러", 10000, 30}
	items[1] = item{"롱패딩", 500000, 20}
	items[2] = item{"투미 백팩", 400000, 20}
	items[3] = item{"나이키 운동화", 150000, 50}
	items[4] = item{"빼빼로", 1200, 500}

	for {
		menu := 0

		fmt.Println("1. 구매")
		fmt.Println("2. 잔여 수량 확인")
		fmt.Println("3. 잔여 마일리지 확인")
		fmt.Println("4. 배송 상태 확인")
		fmt.Println("5. 장바구니 확인")
		fmt.Println("6. 프로그램 종료")
		fmt.Print("실행할 기능을 입력하시오 : ")

		fmt.Scanln(&menu)
		fmt.Println()

		if menu == 1 {
			fmt.Println("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			fmt.Scanln()
		} else if menu == 2 {
			fmt.Println("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			fmt.Scanln()
		} else if menu == 3 {
			fmt.Println("현재 잔여 마일리즌 %d점입니다.\n", buyer.point)
			fmt.Println("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			fmt.Scanln()
		} else if menu == 4 {
			fmt.Println("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			fmt.Scanln()
		} else if menu == 5 {
			fmt.Println("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			fmt.Scanln()
		} else if menu == 6 {
			fmt.Println("프로그램을 종료합니다.")
			return
		} else {
			fmt.Println("잘못된 입력입니다. 다시 입력해주세요. \n")
		}

	}
}
