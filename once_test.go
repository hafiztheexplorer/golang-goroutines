package golanggoroutines

import (
	"fmt"
	"sync"
	"testing"
)

var i = 0

func OnlyOnce() {
	i++
}

func TestOnce(t *testing.T) {
	Once := sync.Once{}
	Group := sync.WaitGroup{}

	for j := 0; j < 100; j++ {
	go func ()  {
		Group.Add(1)
		Once.Do(OnlyOnce)
		// OnlyOnce()
		Group.Done()
	}()	
		
	}

	Group.Wait()
	fmt.Println("counternya =",i)
}