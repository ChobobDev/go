package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main 함수가 시작하였습니다", time.Now())

	go long()
	go short()

	time.Sleep(5 * time.Second)
	fmt.Println("main 함수가 종료되었습니다", time.Now())
}

func long() {
	fmt.Println("long 함수를 시작합니다 ", time.Now())
	time.Sleep(3 * time.Second)
	fmt.Println("long 함수를 종료합니다 ", time.Now())
}

func short() {
	fmt.Println("Short 함수를 시작합니다 ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("Short 함수를 종료합니다 ", time.Now())

}
