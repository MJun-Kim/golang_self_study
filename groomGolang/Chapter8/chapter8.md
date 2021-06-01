# Chapter8 Switch문

사용문법!

```GO
switch 태그/표현식{
case 라벨/표현식:
	수행구문
case 라벨/표현식:
	수행구문
	...
default 라벨/표현식:
	수행구문
}
```

##특징!

- 비교하기 위한 태그값 말고도 표현식이 사용이 가능하다!
- 다른 언어와 다르게 break가 필요가 없음!

  example

```Go
package main

import "fmt"

func main() {
	var num int
	fmt.Print("정수 입력:")
	fmt.Scanln(&num)

	//기본적인 사용! break가 없어도됩
	switch num {
	case 0:
		fmt.Println("영")
	case 1:
		fmt.Println("일")
	case 2:
		fmt.Println("이")
	default:
		fmt.Println("모르겠어요.")
	}

	//표현식 사용!
	switch num / 10 {
	case 1:
		result = "A"
	case 2:
		result = "B"
	case 3:
		result = "C"
	default:
		fmt.Println("모르겠어요.")
	}

	var a, b int

	fmt.Print("정수 a와 b를 입력하시오:")
	fmt.Scanln(&a, &b)

	//case 에서도 표현식을 사용
	switch {
	case a > b:
		fmt.Println("a가 b보다 큽니다.")
	case a < b:
		fmt.Println("a가 b보다 작습니다.")
	case a == b:
		fmt.Println("a와 b가 같습니다.")
	default:
		fmt.Println("모르겠어요.")
	}

}
```

---

연습 문제

안좋은 계산기

```Go
package main

import "fmt"

func main() {
	var sel int;
	var num1 , num2 , result float32;

	fmt.Scan(&sel)
	fmt.Scan( &num1, &num2)

	switch sel {
		case 1:
			result = num1 + num2;
		case 2:
			result = num1 - num2;
		case 3:
			result = num1 * num2;
		case 4 :
			result = num1 / num2;
		default :
			fmt.Print("잘못된입력입니다.")
		return ;

	}

	fmt.Printf("%.1f\n", result)
}
```
