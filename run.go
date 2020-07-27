package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main(){

	runtime.GOMAXPROCS(4)

	waitGroup := new(sync.WaitGroup)
	//
	waitGroup.Add(1)
	//
	//x := func(name string, num int, t time.Time){
	//
	//	defer waitGroup.Done()
	//
	//	for x:=0;x<1<<35; x+=1 {
	//
	//	}
	//
	//	fmt.Println(name, time.Now().Sub(t))
	//}

	//t := time.Now()
	//
	////go x("1st", 20, t)
	////
	////go x("2nd", 20, t)
	////
	////go x("3rd", 20, t)
	////
	////go x("4th", 20, t)

	ch := make(chan string, 100)

	go func() {
		defer func() {
			waitGroup.Add(-1)
			close(ch)
		}()
		for i:=0;i<100;i+=1{
			ch <- fmt.Sprintf("data %d", i)
		}
	}()
	//
	//go func() {
	//	defer waitGroup.Add(-1)
	//	for ;<-ch != ""; {
	//		fmt.Println(fmt.Sprintf("1st %s",<-ch))
	//	}
	//}()
	//
	//go func() {
	//	defer waitGroup.Add(-1)
	//	for ;<-ch != ""; {
	//		fmt.Println(fmt.Sprintf("2nd %s",<-ch))
	//	}
	//}()

	waitGroup.Wait()

}