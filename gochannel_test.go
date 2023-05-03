package golanggoroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

//----------------------------------------------test channel biasa---------------------------------------
func TestCreateChannel(t *testing.T) {
	Kanal1 := make(chan string) //make(chan....) adalah channel yang kita set, di sini sy jadikan variabel kanal1

	// untuk menutup kanal di setiap function
	// bisa juga untuk mematikan dengan defer
	defer close(Kanal1)
	// close(Kanal1)

	// buat anonymous function seperti di bawah syntaxnya
	go func ()  {
		DatatoKanal1:="Contoh data 1"
		time.Sleep(1 * time.Second)
		Kanal1 <- DatatoKanal1 //mengirim data ke kanal1 (channel yang sudah kita set di atas)
		fmt.Println("Selesai Mengirim", DatatoKanal1, "ke Kanal1")
		// fmt.Println("Selesai Mengirim", "ke Kanal1")
	}()

	DatafromKanal1:= <- Kanal1
	fmt.Println(DatafromKanal1)

	time.Sleep(3 * time.Second)
}


//----------------------------------------------channel as parameter---------------------------------------
func BerikanRespon(Kanal2 chan string){
	InputanKanal2:="Contoh Inputan Kanal 2"
	time.Sleep(2 * time.Second)
	Kanal2 <- InputanKanal2
}

func TestChannelAsParameter(t *testing.T) {

	Kanal2 := make(chan string) //make(chan....) adalah channel yang kita set, di sini sy jadikan variabel kanal1

	// untuk menutup kanal di setiap function
	// bisa juga untuk mematikan dengan defer
	

	go BerikanRespon(Kanal2)

	DatafromKanal2:= <- Kanal2
	fmt.Println(DatafromKanal2)

	time.Sleep(3 * time.Second)

}

//----------------------------------------channel in only / out only---------------------------------------



// channel untuk menerima data kedalam kanal saja, dengan ditandai dengan tanda <-
func DataMasukKeKanal(Kanal3 chan<- string)  {
	
	time.Sleep(2 * time.Second)
	InputanKanal3:="Contoh Inputan Kanal 3" //perhatikan ini adalah data yang dimasukkan ke kanal, 
	Kanal3 <- InputanKanal3
	fmt.Println("func Data Masuk Ke Kanal Running")
	}

// channel untuk mengirim data keluar kanal saja, dengan ditandai dengan tanda <-
func DataKeluarDariKanal(Kanal3 <-chan string)  { 

	DatafromKanal3 := <- Kanal3
	fmt.Println(DatafromKanal3) //dan ini data yang kita masukkan tadi , di print isinya sama
	fmt.Println("func Data Keluar Dari Kanal Running")
	}

func TestInandOut(t *testing.T) {
	Kanal3 := make(chan string) //make(chan....) adalah channel yang kita set, di sini sy jadikan variabel kanal1
	
	// untuk menutup kanal di setiap function
	// bisa juga untuk mematikan dengan defer
	defer close(Kanal3)
	// close(Kanal1)
	// buat anonymous function seperti di bawah syntaxnya
	
	go DataMasukKeKanal(Kanal3)
	go DataKeluarDariKanal(Kanal3)
	
time.Sleep(5 * time.Second)
}

//--------------------------------------------buffered channel--------------------------------------------

func TestBufferedChannel(t *testing.T)  {
	Kanal4:=make (chan string, 3) 	//ini adalah variabel kanal yang sudha kita set panjangnya 3
	defer close(Kanal4)				//agar tidak ngeloop kita close

	InputanKanal4_1 := "Contoh Inputan Kanal 4_1"
	InputanKanal4_2 := "Contoh Inputan Kanal 4_2"
	InputanKanal4_3 := "Contoh Inputan Kanal 4_3"
	Kanal4 <- InputanKanal4_1			
	Kanal4 <- InputanKanal4_2
	Kanal4 <- InputanKanal4_3
									//bisa dilihat tanpa kita ambil data dari kanal, ternyata bisa, 
									//itu karena data tersebut dimasukkan ke buffer, bukan ke kanal 
									//secara keseluruhan, bisa dicoba jika buffer number dihapus, 
									//maka akan blocking seperti contoh channel in out sebelumnya

	fmt.Println("\n",<- Kanal4,"\n",<- Kanal4,"\n",<- Kanal4)

	fmt.Println("test buffered channel test sudah selesai")
}


//-----------------------------------------------range channel---------------------------------------------


func TestRangeChannel(t *testing.T)()  {
	Kanal5:=make (chan string)

	go func ()  {
		JumlahDataDikirim:=10
		for i := 1; i < JumlahDataDikirim; i++ {
			Kanal5 <- "Perulangan ke-" + strconv.Itoa(i) 
		}
		defer close(Kanal5)
	}()

	for JumlahData := range Kanal5 {
		fmt.Println("Mencari data",JumlahData)
	}
	fmt.Println("Selesai")
}

//-----------------------------------------------select channel---------------------------------------------

func BerikanRespon2(Kanal6 chan string){
	InputanKanal6:="Contoh Inputan Kanal"
	time.Sleep(2 * time.Second)
	Kanal6 <- InputanKanal6
}

func TestSelectChannel(t *testing.T)   {
	Kanal6_1 := make(chan string)
	Kanal6_2 := make(chan string)
	defer close(Kanal6_1)
	defer close(Kanal6_2)


	go BerikanRespon2(Kanal6_1)		//BerikanResponse2 adalah function untuk menerima data inputan 
	go BerikanRespon2(Kanal6_2)		//ke kanal

	i:=0							//untuk memberikan nilai awal counter pada looping 
									//pencarian data dari kanal
	for {
		select{
		case AmbilDataKanal:= <- Kanal6_1:
			fmt.Println("Data dari kanal6_1 = ", AmbilDataKanal)
			i++						
		case AmbilDataKanal:= <- Kanal6_2:
			fmt.Println("Data dari kanal6_2 = ", AmbilDataKanal)
			i++
		}
		if i==2{
			break
			}
}
}

//-------------------------------------------default select channel----------------------------------------

func BerikanRespon3(Kanal7 chan string){
	InputanKanal7:="Contoh Inputan Kanal"
	time.Sleep(2 * time.Second)
	Kanal7 <- InputanKanal7
}

func TestDefaultSelectChannel(t *testing.T) {
	Kanal7_1 := make(chan string)
	Kanal7_2 := make(chan string)
	defer close(Kanal7_1)
	defer close(Kanal7_2)


	go BerikanRespon3(Kanal7_1)		//BerikanResponse2 adalah function untuk menerima data inputan 
	go BerikanRespon3(Kanal7_2)		//ke kanal

	i:=0							//untuk memberikan nilai awal counter pada looping 
									//pencarian data dari kanal
	for {
		select{
		case AmbilDataKanal2:= <- Kanal7_1:
			fmt.Println("Data dari kanal7_1 = ", AmbilDataKanal2)
			i++						
		case AmbilDataKanal2:= <- Kanal7_2:
			fmt.Println("Data dari kanal7_2 = ", AmbilDataKanal2)
			i++
		default:
			fmt.Println("Standby, Menunggu data")

		}
		if i==2{
			break
			}
}
}