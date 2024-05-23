package worker

type Task struct {
	Spec string
	Exec func()
}

func NewTask(spec string, exec func()) *Task {
	return &Task{
		Spec: spec,
		Exec: exec,
	}
}
