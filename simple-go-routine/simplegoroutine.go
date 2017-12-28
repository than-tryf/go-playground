package main

import (
	"fmt"
	"time"
)

type Job struct {
	i int
	max int
	text string
}


func outText(j *Job){
	for j.i < j.max {
		time.Sleep(10*time.Millisecond)
		fmt.Println(j.text)
		j.i++
	}
}

func main()  {
	hello := new(Job)
	world := new(Job)

	hello.text = "hello"
	hello.i=0
	hello.max=3

	world.text = "world"
	world.i = 0
	world.max = 5

	/*
		Being asynchronous,when a goroutine is invoked,it
		waits for the blocking code to complete before concurrency
		begins.
	*/

	go outText(hello)
	outText(world)

	/*
		If instead:
			go outText(hello)
			go outText(world)
		was run you will get no output because the main function ends
		while the asynchronous goroutines are running.
	*/

}



