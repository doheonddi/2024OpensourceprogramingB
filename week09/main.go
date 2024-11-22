package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

func main() {
	// fmt.Printf("%d\n", rand.Intn(6)+1)
	r := fmt.Sprintf("%10d\n", rand.Intn(100+1))
	fmt.Println(reflect.TypeOf(r))
	fmt.Printf("%T\n", 2.1) // %T는 타입을 알아볼 때 사용
 
	i := 1
	for i <= 10 {  // while
		fmt.Printf("%3d점\n", i)
		i = i + 1
	}
}
