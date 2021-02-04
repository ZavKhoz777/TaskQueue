package main

import (
	"awesomeSolution/TaskQueue"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"sync"
	"time"
)

func main(){

	if len(os.Args) == 0 {
		fmt.Println(len(os.Args))
		return
	}

	writers , _ := strconv.Atoi(os.Args[1])
	arrSize , _ := strconv.Atoi(os.Args[2])
	iters, _ := strconv.Atoi(os.Args[3])

	tskQ := TaskQueue.NewTaskQueue(writers)

	var wg sync.WaitGroup
	wg.Add(writers)

	work := func(id int){
		defer wg.Done()
		for iter := 0; iter < iters; iter++ {
			tskQ.AddTask(id, func() {
				arr := arrGenerator(arrSize)
				sort.Ints(arr)
				insertTime := time.Now().Format("2006-01-02T15:04:05-0700")
				fmt.Printf("{Gorutine № %d} {%s} {min %d} {avg %d} {max %d} \n", id, insertTime, arr[0], arr[arrSize/2], arr[arrSize-1])
			})
		}
	}

	for i := 0; i < writers; i++ {
		go work(i)
	}
	wg.Wait()
	tskQ.DoTasks()
}

func arrGenerator(size int) []int  {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, size)
	for k := range arr{
		arr[k] = rand.Intn(100)

	}

	return arr
}