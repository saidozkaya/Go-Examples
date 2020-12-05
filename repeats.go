package main

import (
	"fmt"
	"time"
)

var a int = 3
var b int = 6

func mains() {

	names := []string{"Ahmet", "Said", "Fatih", "Melih"}
	names = append(names, "Bilal")

	fruits := []string{"Elma", "Armut", "Muz", "Grape", "Lale", "Gül", "Çiçek"}

	fmt.Println(fruits[3:])
}
func ifElseBloklari() {

	if c := 2; c > b && c > a {
		println("C Hepsinden büyüktür.")
	} else if b == a && c == a {
		println("Hepsi Eşitler")
	} else {
		println("İstenilen sonuçlar dışında bir sonuç")
	}
}

func switchCaseBloklari(c int) {

	println("A=3 ve B=6'dır")
	switch {

	case a > b && b > c:
		println("1. Blok ..... En büyük A")
		break
	case b > a:
		println("2. Blok.... En büyük B")

	case c > b && b > a:
		println("3. Blok .... C en büyük ")
		break

	default:
		println("Eğer bloklardan birine uymuyorsa çalışıyor...")
	}
}

func girdiAl() string {
	var girdi string
	println("Sayı giriniz")
	fmt.Scanf("%v", &girdi)
	println("Girdiğiniz değer ", girdi, "'dir")
	return girdi
}

func tarihHesapla() {
	tarih := time.Now()
	tarihe := time.Date(11, time.April, 3, 2, 3, 5, 2, time.UTC)
	fmt.Println(tarihe, "\n", tarih)
}
