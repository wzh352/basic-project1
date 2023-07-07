package main

import (
	"fmt"
	"time"
)

func main(){
	//咋还侮辱人呢，不礼貌，差评！
	s := []string{"I","am","stupid","and","weak"}
	changeStringList(s)
	fmt.Println(s)


	c := make(chan int, 10)
	t := time.NewTicker(time.Second)

	go func() {
		for{
			select {
			case  <-c:
				fmt.Println("receive value")

			case  <-t.C:
				fmt.Println(time.Now())
				c <- 6
				fmt.Println("input value six")
			}
		}
	}()
	time.Sleep(30*time.Second)
}

func changeStringList(a []string)  {
	for i,_ := range a {
		if i == 2{
			a[i] = "smart"
		}
		if i == len(a)-1{
			a[i] = "strong"
		}
	}
}