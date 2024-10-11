package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("이름 입력 : ")
	in := bufio.NewReader(os.Stdin)
	name, err := in.ReadString('\n') // err을 하나 더 주었기때문에 ReadString이 오류 나지 않음
	fmt.Println(name)
	fmt.Println(err)
}
