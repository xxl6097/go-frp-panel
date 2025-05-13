package utils

type AutoReplaceQueue struct {
	ch chan interface{} // 通道实现自动覆盖
}

func NewAutoReplaceQueue(capacity int) *AutoReplaceQueue {
	return &AutoReplaceQueue{
		ch: make(chan interface{}, capacity),
	}
}

// Push 添加元素，通道满时自动丢弃旧元素
func (q *AutoReplaceQueue) Push(item interface{}) {
	select {
	case q.ch <- item: // 通道未满时正常写入
	default:
		// 通道已满时，先取出一个旧元素再写入新元素
		<-q.ch
		q.ch <- item
	}
}

// Pop 取出元素
func (q *AutoReplaceQueue) Pop() interface{} {
	select {
	case item := <-q.ch:
		return item
	default:
		return nil
	}
}

// FixedQueue 泛型固定长度队列结构体
type FixedQueue[T any] struct {
	items    []T // 队列元素
	size     int // 当前大小
	capacity int // 队列容量
	head     int // 队首索引
	tail     int // 队尾索引
}

// NewFixedQueue 创建一个新的固定长度队列
func NewFixedQueue[T any](capacity int) *FixedQueue[T] {
	return &FixedQueue[T]{
		items:    make([]T, capacity),
		size:     0,
		capacity: capacity,
		head:     0,
		tail:     0,
	}
}

// Enqueue 入队操作
func (q *FixedQueue[T]) Enqueue(item T) {
	if q.size < q.capacity {
		q.size++
	} else {
		// 队列已满，移动队首指针
		q.head = (q.head + 1) % q.capacity
	}
	q.items[q.tail] = item
	q.tail = (q.tail + 1) % q.capacity
}

// Dequeue 出队操作
func (q *FixedQueue[T]) Dequeue() (T, bool) {
	var zero T
	if q.size == 0 {
		return zero, false
	}
	item := q.items[q.head]
	q.head = (q.head + 1) % q.capacity
	q.size--
	return item, true
}

// Size 返回队列当前大小
func (q *FixedQueue[T]) Size() int {
	return q.size
}

// Capacity 返回队列容量
func (q *FixedQueue[T]) Capacity() int {
	return q.capacity
}

// Items 返回队列所有元素
func (q *FixedQueue[T]) Items() []T {
	items := make([]T, 0, q.size)
	for i := 0; i < q.size; i++ {
		idx := (q.head + i) % q.capacity
		items = append(items, q.items[idx])
	}
	return items
}
