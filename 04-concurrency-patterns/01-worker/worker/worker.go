package worker

type Work interface {
	Task()
}

type Worker struct {
	/*  */
}

func New(maxTasks int) *Worker {
	/*  */
}

func (worker *Worker) Run(w Work) {
	/*  */
}

func (worker *Worker) Shutdown() {
	//wait for the assigned tasks to be completed
	//exit
}
