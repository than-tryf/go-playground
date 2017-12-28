package main

import (
	"fmt"
	"time"
	"sync"
)

type Job struct {
	i int
	max int
	text string
}

/*
	 i) Added a a pointer arguement of goGroup
	ii) Send Done message --> goGroup.Done()
*/
func outText(j *Job, goGroup *sync.WaitGroup){
	for j.i < j.max {
		time.Sleep(10*time.Millisecond)
		fmt.Println(j.text)
		j.i++
	}
	goGroup.Done()
}

func main()  {

	//Create a Waitgroup struct that will receive note that goroutines
	//have completed x number of times before allowing the program to exit
	goGroup := new(sync.WaitGroup)

	fmt.Println("Starting...")

	hello := new(Job)
	world := new(Job)

	hello.text = "hello"
	hello.i=0
	hello.max=3

	world.text = "world"
	world.i = 0
	world.max = 5


	go outText(hello, goGroup)
	go outText(world, goGroup)

	//Add method specifies how many Done messages goGroup should receive
	//before saisfying its wait
	goGroup.Add(2)

	//Tell Waitgroup to wait
	goGroup.Wait()
}



