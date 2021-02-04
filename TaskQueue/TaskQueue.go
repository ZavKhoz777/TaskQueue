package TaskQueue

import (
	"sync"
)
/* Главная структура очереди */
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
	task func()
}

func newTask(identifier int, task func()) *Task {
	return &Task{identifier: identifier, task: task}
}
/* Добавление тасков */
func (queue *TaskQueue) AddTask (goroutine int,f func()){
	queue.mx.Lock()
	task := newTask(goroutine, f)
	queue.tasks[goroutine] = append(queue.tasks[goroutine], *task)
	queue.mx.Unlock()
}
/** Выполнение тасков и удаление выполненных */
func (queue *TaskQueue) DoTasks (){
	queue.mx.Lock()
	 for _,v := range queue.tasks{
		for n, f := range v{
			f.task()
			if len(queue.tasks[n]) > 1 {
				queue.tasks[n] = queue.tasks[n][1:]
			} else {
				delete(queue.tasks, n)
			}
		}
	}
	queue.mx.Unlock()
}
