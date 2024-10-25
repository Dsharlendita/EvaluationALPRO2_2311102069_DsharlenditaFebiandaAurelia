//D'sharlendita Febianda Aurelia
//2311102069

package main

import "fmt"

func main(){
	var a, b, c int
	var n int

	fmt.Print("Masukkan bilangan bulat positif lebih dari 100: ")
	fmt.Scan(&a, &b, &c)

	fmt.Print("Hasil Pemotongan: ")
	fmt.Println(a, " ", b,  " ", c)
	

	n = (a + b + c)/3

	fmt.Print("Rata-Rata dari ketiga bilangan: ", n)
}

