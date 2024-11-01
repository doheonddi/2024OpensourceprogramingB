package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
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

	counts := 0

	if n <= 1 {
		counts = -1
	} else {
		i := 2
		for i < n {
			if n%i == 0 {
				counts++
			}
			i++
		}
	}

	if counts == 0 {
		fmt.Printf("%d는(은) 소수입니다.\n", n) // 수정된 부분: 콤마 추가
	} else {
		fmt.Printf("%d는(은) 소수가 아닙니다.\n", n) // 수정된 부분: 콤마 추가
	}
}
