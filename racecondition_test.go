package golanggoroutines

import (
	"fmt"
	"testing"
	"time"
)

//--------------------------------------Contoh Race Conditions----------------------------------

func TestRaceCondition(t *testing.T) {
	x := 0
	for i := 1; i <= 1000; i++ {
		go func ()  {						//coba dirubah tanpa go, hasilnya akan menjadi 100.000
			for j := 1; j <= 100; j++ {		//kenapa berbeda, karena jika pakai go , maka 
				x = x + 1					//penambahan kadang2 bermula dari nilai yang sama
			}
		}()
	}

	time.Sleep(5*time.Second)
	fmt.Println("Penghitung = ", x)
}

