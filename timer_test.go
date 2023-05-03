package golanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

//------------------------------------test channel times--------------------------------

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <- timer.C
	fmt.Println(time)
}

//------------------------------------test channel time After--------------------------------

func TestTimeAfter(t *testing.T){
	kanal := time.After(5 * time.Second)
	fmt.Println(time.Now())

	time := <- kanal
	fmt.Println(time)
}


//------------------------------------test After Function--------------------------------

func TestAfterFunction(t *testing.T)()  {
	group:=sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(5*time.Second, func() {
		fmt.Println(time.Now())
		group.Done()
	})
	fmt.Println(time.Now())

	group.Wait()
}

//------------------------------------test Time Ticker--------------------------------

func TestTimeTicker(t *testing.T)()  {
	ticker:=time.NewTicker(1*time.Second)

	go func ()  {
		time.Sleep(5*time.Second)
		ticker.Stop()
	}()

	for time:=range ticker.C {
		fmt.Println(time)
	}

	
}

//------------------------------------test Time Tick--------------------------------

func TestTimeTick(t *testing.T)()  {
	kanal:=time.Tick(1*time.Second)

	// go func ()  {
	// 	time.Sleep(5*time.Second)
	// 	ticker.Stop()
	// }()

	for time:=range kanal {
		fmt.Println(time)
	}

	
}