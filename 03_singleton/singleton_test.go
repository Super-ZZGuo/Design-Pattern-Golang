package singleton

import (
	"sync"
	"testing"
)

func TestSingleton(t *testing.T) {
	singleton1 := GetInstance()
	singleton2 := GetInstance()
	if singleton1 != singleton2 {
		t.Fatal("instance is not equal!")
	}
}

func TestParallelSingleton(t *testing.T) {
	start := make(chan struct{})
	wg := sync.WaitGroup{}
	const parCnt = 100
	wg.Add(parCnt)
	instances := [parCnt]Singleton{}
	for i := 0; i < parCnt; i++ {
		go func(index int) {
			// t.Log(strconv.Itoa(index))
			// 阻塞协程
			<-start
			instances[index] = GetInstance()
			wg.Done()
		}(i)
	}
	// 关闭chan 同步的协程开始执行
	close(start)
	wg.Wait()
	for i := 0; i < parCnt-1; i++ {
		if instances[i] != instances[i+1] {
			t.Fatal("instance is not equal!")
		}
	}
}
