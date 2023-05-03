package golanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

//----------------------------------Contoh Race Conditions dengan Mutex------------------------------

func TestMutex(t *testing.T) {
	x := 0
	var Mutex sync.Mutex								//struct sync.Mutex digunakan dan dijadikan var agar 
														//bisa dipanggil dalam func di bawahnya, 
														//pada perulangan
	for i := 1; i <= 1000; i++ {
		go func ()  {									//coba dirubah tanpa go, 
														//hasilnya akan menjadi 100.000
			for j := 1; j <= 100; j++ {					//kenapa berbeda, karena jika pakai go , maka 
				//mutex lock di awal penambahan nilai
				Mutex.Lock()				
				x = x + 1								//penambahan kadang2 bermula dari nilai yang sama
				//mutex lock di awal penambahan nilai	//maka dari itu dibutuhkan mutex
				Mutex.Unlock()
			}
		}()
	}

	time.Sleep(5*time.Second)
	fmt.Println("Penghitung = ", x)
}

//---------------------------------Contoh Race Conditions dengan RWMutex-----------------------------

type BankAccount struct {								//mutex di sini lebih mendetail, karena terpisah
RWMutex sync.RWMutex									//dengan menggunakan struct sync.RWmutex
Balance int
}

func (Account *BankAccount) AddBalance(Amount int) {
	Account.RWMutex.Lock()								//sama seperti di function sebelumnya, mutex ini
	Account.Balance = Account.Balance + Amount			//ditempatkan di sebelum dan sesudah proses
	Account.RWMutex.Unlock()							//penambahan data
}

func (Account *BankAccount) GetBalance() int {
	Account.RWMutex.Lock()
	Balance2:=Account.Balance							//atau kalau di sini pada proses pengambilan
	Account.RWMutex.Unlock()							//data
	return Balance2
}

func TestRWMutex(t *testing.T){
Account := BankAccount{}

for i := 1; i <= 1000; i++ {
	go func ()  {
		for j := 1; j <= 100; j++ {
			Account.AddBalance(1)
			fmt.Println(Account.GetBalance())
		}
	}()
}
	time.Sleep(5 * time.Second)
	fmt.Println("total balance adalah = ", Account.GetBalance())
}

//---------------------------------Contoh Deadlock Conditions dengan Mutex-----------------------------

type UserBalance struct {
sync.Mutex										//bisa juga ditulis dengan sync.Mutex, tapi,
Nama string												//di methodnya dipanggilnya sebagai Mutex saja
Balance int												//jadi tidak perlu Simutex, cukup "Mutex"
}														//don't ask why

func (User *UserBalance) Lock() {						//Method untuk lock
	User.Mutex.Lock()
}

func (User *UserBalance) Unlock() {						//Method untuk unlock
	User.Mutex.Unlock()
}

func (User *UserBalance) Change(Amount int) {			//method untuk menambah balance
	User.Balance = User.Balance + Amount
}

func Transfer(User1 *UserBalance, User2 *UserBalance, Amount int)  {		//prinsipnya seperti ini, jadi dari user1 transfer ke user2
																			//maka dari itu nanti amount yang di user1 ditransfer dikurangi
	User1.Lock()															//lalu di user2 ditambah amountnya karena dapat trasferan dari user1
	fmt.Println("Lock engaged User1=", User1.Nama)
	User1.Change(-Amount)

	time.Sleep(1 * time.Second)
	
	User2.Lock()
	fmt.Println("Lock engaged User2=", User2.Nama)
	User2.Change(Amount)

	time.Sleep(1 * time.Second)

	User1.Unlock()
	User2.Unlock()
	
}

func TestDeadLock(t *testing.T) {
	User1:=UserBalance{
		Nama:"Contoh User 1",
		Balance: 100000,
	}
	User2:=UserBalance{
		Nama:"Contoh User 2",
		Balance: 100000,
	}

	go Transfer(&User1, &User2, 10000)
	go Transfer(&User2, &User1, 10000)

	time.Sleep(2 * time.Second)

	fmt.Println("User=",User1.Nama,", Saldonya=",User1.Balance)
	fmt.Println("User=",User2.Nama,", Saldonya=",User2.Balance)
}
