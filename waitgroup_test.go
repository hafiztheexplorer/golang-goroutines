package golanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

//---------------------------------Wait Group-----------------------------

func RunAsync(Group *sync.WaitGroup) {
	defer Group.Done()

	Group.Add(1)

	fmt.Println("Starting Hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
WaitGroup:=&sync.WaitGroup{}

for i := 0; i < 100; i++ {
	go RunAsync(WaitGroup)
}
WaitGroup.Wait()
fmt.Println("it's done")
}