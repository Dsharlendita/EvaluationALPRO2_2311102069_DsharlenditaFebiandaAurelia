//D'sharlendita Febianda Aurelia
//2311102069

package main

import "fmt"

func main() {

	var jumlahPeserta int
	var nomorTiket int
	var hadiahUtama, hadiahSembako, hadiahKonsol int

	fmt.Print("Masukkan Jumlah Peserta: ")
	fmt.Scan(&jumlahPeserta)

	for i := 1; i <= jumlahPeserta; i++ {
		fmt.Print("Masukkan nomor tiket peserta: ")
		fmt.Scan(&nomorTiket)

		if nomorTiket %2 == 0 && nomorTiket == nomorTiket{
			fmt.Println("Hadiah Utama")
			hadiahUtama++
		} else if nomorTiket %2 != 0 {
			fmt.Println("Hadiah Sembako")
			hadiahSembako++
		} else {
			fmt.Println("Hadiah Konsol")
			hadiahKonsol++
		}
	}
	fmt.Println("Jumlah Hadiah:")
	fmt.Printf("Hadiah Utama: %d\n", hadiahUtama)
	fmt.Printf("Hadiah Sembako: %d\n", hadiahSembako)
	fmt.Printf("Hadiah Konsol: %d\n", hadiahKonsol)
}

