package golanggoroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxprocs(t *testing.T) {
group:=sync.WaitGroup{}

for i := 0; i < 10000; i++ {
	group.Add(1)
	go func ()  {
		time.Sleep(5 * time.Second)
		group.Done()
	}()
}

totalcpu:=runtime.NumCPU()
fmt.Println(totalcpu)

totalthread:=runtime.GOMAXPROCS(-1)
fmt.Println(totalthread)

totalgorutin:=runtime.NumGoroutine()
fmt.Println(totalgorutin)

group.Wait()
}
