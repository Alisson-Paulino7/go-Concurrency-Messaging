package main

import (
	"fmt"
	"time"
)

//T1
func main() {

	canal := make(chan int)
		
	go func () {
		for i := 0; i < 10; i++ {
			canal <- i //T2 - Enche o canal
			fmt.Println("Jogou no canal: ", i)
		}
	}()

	go func () {
		for i := 10; i < 20; i++ {
			canal <- i //T2 - Enche o canal
			fmt.Println("Jogou no canal: ", i)
		}
	}()

	// for x := range canal{
	// 	fmt.Println("Recebeu do canal:", x)
	// 	time.Sleep(time.Second)
	// }

	go worker(canal,1)
	worker(canal,2)

	// fmt.Println(<-canal) //Esvaziar o canal

	// time.Sleep(time.Second * 2)
	
}

func worker(canal chan int, workerID int ) {
	for {
		fmt.Println("Recebeu do canal:", <-canal, "no worker", workerID)
		time.Sleep(time.Second)
	}

}