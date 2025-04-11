package pool

type WorkerPool struct {
	tasks chan func()
}

func NewWorkerPool(size int) *WorkerPool {
	pool := &WorkerPool{
		tasks: make(chan func(), size),
	}
	for i := 0; i < size; i++ {
		go pool.worker()
	}
	return pool
}

func (p *WorkerPool) worker() {
	for task := range p.tasks {
		task()
	}
}

func (p *WorkerPool) Submit(task func()) {
	p.tasks <- task
}

func test() {
	// 使用示例
	pool := NewWorkerPool(5) // 5个worker
	pool.Submit(func() {
		//do something
	})
}
