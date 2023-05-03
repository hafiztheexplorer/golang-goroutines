package golanggoroutines

import (
	"fmt"
	"math"
	"testing"
	"time"
)

//contoh function yang dipanggil oleh unit test di bawah
func ContohFunction1() {
	fmt.Println("contoh function 1 sudah dijalankan")

}

//function dengan menggunakan goroutines
func TestDenganGoroutines1(t *testing.T) {
	go ContohFunction1() //goroutines dengan menggunakan "go" sebelum memanggil function
	fmt.Println("contoh goroutines create 1")
	time.Sleep(1 * time.Second)
	}

//function dengan tanpa goroutines	
func TestTanpaGoroutines1(t *testing.T) {
	ContohFunction1()
	fmt.Println("contoh goroutines create 1")
	time.Sleep(1 * time.Second)
	}

// contoh ringannya Go-Routines
func DisplayNumber(number int) {
	perhitungan:=math.Log2(float64(number))
	fmt.Println("Display ", perhitungan)
	}

func TestRingannyaGoroutines(t *testing.T){
	for i := 0; i < 100000; i++ {

		go DisplayNumber(i)
	}
	time.Sleep(10 * time.Second)
}