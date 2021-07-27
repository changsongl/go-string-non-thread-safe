package main

import (
	"fmt"
	"time"
)

var a = "0"

func main(){
	ch := make(chan string)
	go func() {
		i := 1
		for {
			if i%2 == 0{
				a = "0"
			}else{
				a = "bb"
			}

			time.Sleep(1*time.Millisecond)
			i++
		}
	}()

	go func() {
		for{
			b := a
			if b != "0" && b != "bb"{
				ch <- b
			}
		}
	}()

	for i:=0; i< 10; i++{
		fmt.Println(<-ch)
	}
}