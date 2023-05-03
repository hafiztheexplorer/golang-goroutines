package golanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	Pool := sync.Pool{
		New:func()interface{} {
		return "new"
		},
	}

	Pool.Put("Nama 1")
	Pool.Put("Nama 2")
	Pool.Put("Nama 3")

	count:=10
	for i := 0; i < count; i++ {
		go func ()  {
			data:=Pool.Get()
			fmt.Println(data)
			time.Sleep(1*time.Second)
			Pool.Put(data)
		}()
	}

	time.Sleep(11*time.Second)
	fmt.Println("all done sire")
}