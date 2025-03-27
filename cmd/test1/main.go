package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func DynamicSelect[T any](t []T, fun func(T, chan<- T)) T {
	ch := make(chan T, len(t)) // 缓冲大小等于协程数量
	for _, v := range t {
		go fun(v, ch)
	}
	_, value, _ := reflect.Select([]reflect.SelectCase{{
		Dir:  reflect.SelectRecv,
		Chan: reflect.ValueOf(ch),
	}})

	//fmt.Printf("返回: %v\n", value)
	return value.Interface().(T)
}

func dynamicSelect(channels chan any) any {
	//cases := make([]reflect.SelectCase, len(channels))
	//for i, ch := range channels {
	//	cases[i] = reflect.SelectCase{
	//		Dir:  reflect.SelectRecv,
	//		Chan: reflect.ValueOf(ch),
	//	}
	//}
	for {
		chosen, value, ok := reflect.Select([]reflect.SelectCase{{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(channels),
		}})
		fmt.Printf("ok:%v,通道 %d 返回: %v\n", ok, chosen, value)
		if ok {
			//return value
		}
	}

}

func worker(id int, result chan<- any) {
	// 模拟耗时计算
	t := time.Duration(rand.Intn(4))*time.Second + time.Second
	time.Sleep(t)
	result <- fmt.Sprintf("Worker%d 完成 %v", id, t)
}

func main() {
	//resultChan := make(chan any, 5) // 缓冲大小等于协程数量
	//// 启动多个计算协程
	//for i := 0; i < 5; i++ {
	//	go worker(i, resultChan)
	//}
	//fmt.Println("resultChan.size", len(resultChan))
	//v := dynamicSelect(resultChan)
	//fmt.Printf("------>%v\n", v)

	strs := []string{"a", "b", "c", "d", "e", "f", "g"}
	r := DynamicSelect[string](strs, func(s string, ch chan<- string) {
		t := time.Duration(rand.Intn(10))*time.Second + time.Second
		fmt.Println(s, t)
		time.Sleep(t)
		ch <- fmt.Sprintf("Worker-%v 完成 %v", s, t)
	})
	fmt.Println(r)
}
