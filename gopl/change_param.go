package main

import (
	"fmt"
	"log"
	"time"
)

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func bigSlowOperation() {

	log.Printf("enter 11111")
	defer trace("bigSlowOperation")()
	defer pp()
	defer log.Printf("enter eeeeeeeeeeee")
	log.Printf("enter 222")
	//extra parenthess
	//... lots fo work...
	time.Sleep(10 * time.Second)
	//operation by sleeping
}

func pp() {
	log.Printf("xxxxxxxxx")
}
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func f(x int) {
	//if x == 0 {
	//	return
	//}
	defer func() {
		if p := recover(); p != nil {
			//p := recover()
			fmt.Printf("internal error: %v", p)
		}
	}()
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

func main() {
	fmt.Println(sum())
	fmt.Println(sum(0))
	fmt.Println(sum(0, 1))
	fmt.Println(sum(0, 1, 2))
	fmt.Println(sum(0, 1, 2, 3))
	log.Printf("enter 00000")
	bigSlowOperation()
	log.Printf("enter 3333")

	f(3)

}
