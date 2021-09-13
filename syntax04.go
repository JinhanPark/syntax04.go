package main

import(
	"fmt"
	"math/rand"
	"time"
)

func main(){
	seed := time.Now().Unix()
	rand.Seed(seed)

	isPrime := true
	number := rand.Intn(150) + 2
	fmt.Println("임의로 추출된 수 : ", number)
	if isPrime { //비교 연산자 제거}
		fmt.Println(number, "는(은) 소수입니다!")
	
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다~")
	}
}