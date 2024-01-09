package main

import (
	"fmt"
	"time"

	"github.com/thanhquy1105/GolangPattern/singleton"
)

func main() {

	// s1 := singleton.GetInstance()
	// s2 := singleton.GetInstance()
	// fmt.Println(s1.AddOne())
	// fmt.Println(s2.AddOne())
	// fmt.Println(s1.AddOne())
	// fmt.Println(s2.AddOne())
	// fmt.Printf("%p\n", s1)
	// fmt.Printf("%p\n", s2)
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Printf("%p\n", singleton.GetInstance())
		}()
	}
	time.Sleep(time.Second * 10)
}
