package golanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var Locker = sync.Mutex{}
var cond = sync.NewCond(&Locker)
var group = sync.WaitGroup{}

func WaitCondition(value int)  {
	defer group.Done()
	group.Add(1)
	cond.L.Lock()
	cond.Wait()
	fmt.Println("done ", value)

	cond.L.Unlock()
}

func TestCond(t *testing.T)()  {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	//dijalankan satu satu dengan menggunakan cond.Signal seperti ini
	//perhatikan bahwa outputnya akan keluar satu2
	// go func ()  {
	// 	for i := 0; i < 10; i++ {
	// 		time.Sleep(1*time.Second)
	// 		cond.Signal()
	// 	}	
	// }()

	//dijalankan bersamaan dengan menggunakan cond.Broadcast seperti ini
	//perhatikan bahwa outputnya akan keluar seakan2 bersamaan
	go func ()  {
		for i := 0; i < 10; i++ {
			time.Sleep(1*time.Second)
			cond.Broadcast()
		}
	}()

	group.Wait()
}