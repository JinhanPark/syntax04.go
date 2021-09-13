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

}