package main

import (
	"fmt"
	"time"
)

func getIntSlowly() int {
	time.Sleep(time.Millisecond * 500)
	return 123
}

func funcWithChanResult() chan int {
	chanint := make(chan int)
	go func() {
		chanint <- getIntSlowly()
	}()
	return chanint
}

func funcWithNonChanResult() int {
	var result int
	ch := funcWithChanResult()
	result = <-ch
	return result
}

func main() {
	fmt.Println("Received first int:", <-funcWithChanResult())
	fmt.Println("Received second int:", funcWithNonChanResult())
}
