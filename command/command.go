package main

type Command interface {
	Execute()
}

type TaskQueue struct {
	commands []Command
}

func (q *TaskQueue) Add(cmd Command) {
	q.commands = append(q.commands, cmd)
}

func (q *TaskQueue) Process() {
	for _, cmd := range q.commands {
		cmd.Execute()
	}
	q.commands = nil
}
