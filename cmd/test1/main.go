package main

import (
	"github.com/xxl6097/glog/glog"
	"reflect"
)

func DynamicSelect[T any](t []T, fun func(T) T) T {
	ch := make(chan T, len(t)) // 缓冲大小等于协程数量
	for _, v := range t {
		go func(t T, c chan<- T) {
			c <- fun(t)
		}(v, ch)
	}
	var ret T
	for range ch {
		_, value, ok := reflect.Select([]reflect.SelectCase{{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		}})
		ret = value.Interface().(T)
		if ok {
			return ret
		}
	}

	//for i := 0; i < len(t); i++ {
	//	_, value, ok := reflect.Select([]reflect.SelectCase{{
	//		Dir:  reflect.SelectRecv,
	//		Chan: reflect.ValueOf(ch),
	//	}})
	//	ret = value.Interface().(T)
	//	if ok {
	//		return ret
	//	}
	//}
	return ret
}

func sender(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		glog.Printf("Sent %d to the channel\n", i)
	}
	close(ch)
}

func test1() {
	ch := make(chan int, 1)
	go sender(ch)

	for num := range ch {
		glog.Printf("Received %d from the channel\n", num)
	}
	glog.Println("Channel is closed")
}
func main() {
	test1()

	//strs := []string{"a", "b", "c", "d", "e", "f", "g"}
	//r := DynamicSelect[string](strs, func(s string) string {
	//	t := time.Duration(rand.Intn(10))*time.Second + time.Second
	//	fmt.Println(s, t)
	//	time.Sleep(t)
	//	return fmt.Sprintf("Worker-%v 完成 %v", s, t)
	//})
	//fmt.Println(r)
}
