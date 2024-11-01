package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("%f\n", math.Sqrt(25.0))
	fmt.Print("정수 입력 : ")
	in := bufio.NewReader(os.Stdin)
	number, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	number = strings.TrimSpace(number) // 줄바꿈, 띄어쓰기, 탭 등 제거
	n, err := strconv.Atoi(number)

	if err != nil {
		log.Fatal(err)
	}

	var isPrime bool = true
	if n <= 1 {
		isPrime = false
	} else if n == 2 {
		isPrime = true
	} else if n%2 == 0 { // 2를 제외한 모든 짝수는 소수가 아님.
		isPrime = false
	} else { // 3이상의 홀수
		i := 3
		// for i <= int(math.Sqrt(float64(n))) {
		for i*i <= n {
			if n%i == 0 {
				isPrime = false
				break
			}
			fmt.Printf("%d ", i) // 반복 횟수 확인용 코드
			i = i + 2
		}
	}

	if isPrime {
		fmt.Printf("%d는(은) 소수입니다.\n", n)
	} else {
		fmt.Printf("%d는(은) 소수가 아닙니다.\n", n)
	}
}
