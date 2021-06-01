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
	shoppingbucket map[string]int
}

func newBuyer() *buyer {
	d := buyer{}
	d.point = 1000000
	d.shoppingbucket = map[string]int{}
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

		if menu == 1 { // 물건 구매
			itemchoice := 0
			for i := 0; i < 5; i++ {
				fmt.Printf("물품 %d : %s, 가격 : %d, 잔여 수량 : %d\n", i+1, items[i].name, items[i].price, items[i].amount)
			}
			fmt.Print("구매할 물품을 선택하세요 : ")
			fmt.Scanln(&itemchoice)
			fmt.Println()

			if itemchoice == 1 {
				buying(items, buyer, itemchoice)
			} else if itemchoice == 2 {
				buying(items, buyer, itemchoice)
			} else if itemchoice == 3 {
				buying(items, buyer, itemchoice)
			} else if itemchoice == 4 {
				buying(items, buyer, itemchoice)
			} else if itemchoice == 5 {
				buying(items, buyer, itemchoice)
			} else {
				fmt.Println("잘못된 입력입니다. 다시 입력해주세요.")
			}
			fmt.Println("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			fmt.Scanln()

		} else if menu == 2 { // 잔여 수량 확인
			for i := 0; i < 5; i++ {
				fmt.Printf("%s, 잔여 수량: %d\n", items[i].name, items[i].amount)
			}
			fmt.Println("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			fmt.Scanln()

		} else if menu == 3 { // 잔여 마일리지 확인
			fmt.Printf("현재 잔여 마일리즌 %d점입니다.\n", buyer.point)
			fmt.Println("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			fmt.Scanln()

		} else if menu == 4 { // 배송 상태 확인
			fmt.Println("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			fmt.Scanln()

		} else if menu == 5 { // 장바구니 확인
			bucketmenu := 0
			for {
				emptyBucket(buyer)
				canbuy := requiredPoint(items, buyer) // 살수 있는지 없는지 확인하는 canbuy 선언 및 초기화
				canbuy = excessAmount(items, buyer)

				fmt.Println("1. 장바구니 상품 주문")
				fmt.Println("2. 장바구니 초기화")
				fmt.Println("3. 메뉴로 돌아가기")
				fmt.Scanln(&bucketmenu)
				fmt.Println()

				if bucketmenu == 1 {
					if canbuy {
						bucketBuying(items, buyer)
						break
					} else {
						fmt.Println("구매할 수 없습니다.")
						break
					}
				} else if bucketmenu == 2 {
					fmt.Println()
					break
				} else {
					fmt.Println("잘못된 입력입니다. 다시 입력해주세요.")
				}

			}
			fmt.Println("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
			fmt.Scanln()

		} else if menu == 6 { // 프로그램 종료
			fmt.Println("프로그램을 종료합니다.")
			return

		} else {
			fmt.Println("잘못된 입력입니다. 다시 입력해주세요.")
		}

	}
}

func bucketBuying(itm []item, byr *buyer) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("\n", r, "\n")
		}
	}()

	if len(byr.shoppingbucket) == 0 {
		panic("주문 가능한 목록이 없습니다.")
	} else {
		for index, val := range byr.shoppingbucket {
			for i := range itm {
				if itm[i].name == index {
					itm[i].amount -= val            // 수량 차감
					byr.point -= itm[i].price * val // 포인트 차감
				}
			}
		}
	}

	byr.shoppingbucket = map[string]int{} // 장바구니 초기화
	fmt.Println("주문 접수 되었습니다.")
}

func requiredPoint(itm []item, byr *buyer) (canbuy bool) {
	bucketpoint := 0
	for index, val := range byr.shoppingbucket {
		for i := 0; i < len(itm); i++ {
			if itm[i].name == index {
				bucketpoint += itm[i].price * val
			}
		}
	}
	fmt.Println("필요 마일리지 : ", bucketpoint)
	fmt.Println("보유 마일리지 : ", byr.point)
	fmt.Println()
	if byr.point < bucketpoint {
		fmt.Println("마일리지가 %d점 부족합니다.", bucketpoint-byr.point)
		return false
	}
	return true
}

func excessAmount(itm []item, byr *buyer) (canbuy bool) {
	for index, val := range byr.shoppingbucket {
		for i := 0; i < len(itm); i++ {
			if itm[i].name == index {
				if itm[i].amount < val {
					fmt.Println("%s, %d개 초과", itm[i].name, val-itm[i].amount)
					return false
				}
			}
		}
	}
	return true
}

func emptyBucket(byr *buyer) {
	if len(byr.shoppingbucket) == 0 {
		fmt.Println("장바구니가 비었습니다.")
	} else {
		for index, val := range byr.shoppingbucket {
			fmt.Printf("%s, 수량 : %d\n", index, val)
		}
	}
	fmt.Println()

}

func buying(itm []item, byr *buyer, itmchoice int) {
	defer func() {
		err := recover()
		fmt.Println(err)

	}()
	inputamount := 0

	fmt.Print("수량을 입력하세요 : ")
	fmt.Scanln(&inputamount)
	fmt.Println()

	if inputamount <= 0 {
		panic("올바른 수량을 입력하세요.")
	}
	if byr.point < itm[itmchoice-1].price*inputamount || itm[itmchoice-1].amount < inputamount {
		panic("주문이 불가능합니다.")
	} else {
		for {
			buy := 0
			fmt.Println("1. 바로 주문\n 2. 장바구니에 담기")
			fmt.Print("실행할 기능을 입력하시오 : ")
			fmt.Scanln(&buy)
			fmt.Println()

			if buy == 1 { // 바로 주문
				itm[itmchoice-1].amount -= inputamount
				byr.point -= itm[itmchoice-1].price * inputamount

				fmt.Println("상품이 주문 접수 되었습니다.")
				break
			} else if buy == 2 { // 장바구니에 담기
				checkbucket := false

				for itms := range byr.shoppingbucket { // 물품체크
					if itms == itm[itmchoice-1].name {
						checkbucket = true
					}
				}

				if checkbucket == true { // 장바구니에 중복되는 물품이 있을때
					temp := byr.shoppingbucket[itm[itmchoice-1].name] + inputamount
					if temp > itm[itmchoice-1].amount {
						fmt.Println("물품이 잔여 수량을 초과했습니다.")
						break
					} else {
						byr.shoppingbucket[itm[itmchoice-1].name] += inputamount // 기존 품목에 수량만 더함
					}
				} else {
					byr.shoppingbucket[itm[itmchoice-1].name] = inputamount // 새로 품목 추가
				}
				fmt.Println("상품이 장바구니에 추가되었습니다.")
				break
			} else {
				fmt.Println("잘못된 입력입니다. 다시 입력해주세요.")
			}
		}
	}
}
