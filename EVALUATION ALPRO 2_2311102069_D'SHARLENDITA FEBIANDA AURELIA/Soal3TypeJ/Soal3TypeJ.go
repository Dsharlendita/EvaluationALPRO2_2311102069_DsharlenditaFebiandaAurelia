//D'sharlendita Febianda Aurelia
//2311102069

package main

import "fmt"

func main(){
	var n, m, hasil float64
	fmt.Print("Masukan Bilangan Pembilang: ")
	fmt.Scan(&n)

	fmt.Print("Masukan Bilangan Penyebut: ")
	fmt.Scan(&m)

	hasil = n/m
	fmt.Print("Hasil dari pembagian ", n, " dengan ", m, " adalah: ", hasil)
}