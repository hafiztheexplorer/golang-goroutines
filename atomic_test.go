package golanggoroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)



func TestAtomic(t *testing.T) {
	var x int64 = 0
	Group:=sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		
		go func ()  {	
			Group.Add(1)					
			for j := 1; j <= 100; j++ {		
				atomic.AddInt64(&x,1)					
			}
			Group.Done()
		}()
		
	}
	
	Group.Wait()
	fmt.Println("counter = ",x)
}