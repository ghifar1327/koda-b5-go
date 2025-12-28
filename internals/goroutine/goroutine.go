package goroutine

import (
	"fmt"
	"sync"
	"time"
)

func Rutinitas() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func ()  {
		defer wg.Done()
		mandi := []string{"mandi......", "mulai", "basah", "selesai"}
		for i := range mandi{
			time.Sleep(time.Millisecond * 1000)
			fmt.Println(mandi[i])
		}
	}()
	wg.Wait()
	wg.Add(1)
	go func ()  {
		defer wg.Done()
		buatKopi := []string{"buat kopi.....", "mulai", "ambil gelas","ambil kopi","buat kopi" , "selesai"}
		for i := range buatKopi{
			time.Sleep(time.Millisecond * 1000)
			fmt.Println(buatKopi[i])
		}
	}()
	wg.Wait()
	wg.Add(1)
	go func ()  {
		defer wg.Done()
		sarapan := []string{"menyiapkan sarapan......", "mulai", "masak atau pesen", "selesai"}
		for i := range sarapan{
			time.Sleep(time.Millisecond * 1000)
			fmt.Println(sarapan[i])
		}
	}()
	wg.Wait()
	wg.Add(1)
	go func ()  {
		defer wg.Done()
		rapihkanKamar := []string{"rapihkan kamar.....", "mulai","menata bantal", "melipat seliput", "selesai"}
		for i := range rapihkanKamar{
			time.Sleep(time.Millisecond * 1000)
			fmt.Println(rapihkanKamar[i])
		}
	}()
	wg.Wait()
	fmt.Println("Berangkat..........")
}