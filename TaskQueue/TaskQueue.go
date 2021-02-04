package TaskQueue

import (
	"sync"
)

type TaskQueue struct {
	mx sync.RWMutex
	writers int
	tasks sync.Map[int][]Task

}

func NewTaskQueue(writers int) *TaskQueue {
	response := &TaskQueue{writers: writers, tasks: map[int][]Task{}}
	return response
}

type Task struct {
	identifier int
	task func()
}

func newTask(identifier int, task func()) *Task {
	return &Task{identifier: identifier, task: task}
}


func (queue TaskQueue) AddTask (goroutine int,f func()){

	queue.mx.Lock()
	task := newTask(goroutine, f)
	queue.tasks[goroutine] = append(queue.tasks[goroutine], *task)
	queue.mx.Unlock()
}

func (queue TaskQueue) DoTasks (){
	 for _,v := range queue.tasks{
		for _, f := range v{
			f.task()
		}
	}
}
