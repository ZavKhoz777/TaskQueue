package TaskQueue

import (
	"sync"
	"time"
)

type TaskQueue struct {
	mx sync.RWMutex
	writers int
	tasks map[int][]Task

}

func NewTaskQueue(writers int) *TaskQueue {
	response := &TaskQueue{writers: writers, tasks: map[int][]Task{}}
	return response
}

type Task struct {
	identifier int
	timeStamp time.Time
	task func()
}

func newTask(identifier int, task func()) *Task {
	return &Task{identifier: identifier, task: task}
}


func (queue TaskQueue) AddTask (goroutine int,f func()){
	queue.mx.Lock()
	task := newTask(goroutine, f)
	queue.mx.Unlock()
	task.timeStamp = time.Now()
	queue.tasks[goroutine] = append(queue.tasks[goroutine], *task)
}

func (queue TaskQueue) DoTasks (){
	 for _,v := range queue.tasks{
		for _, f := range v{
			f.task()
		}
	}
}
