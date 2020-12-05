package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	//	*********----------------STRING AND DEFER------------------
	defer fmt.Println(strings.Contains("text", "te"))
	fmt.Println(strings.Count("text", "t"))
	fmt.Println(strings.HasPrefix("text.rar", "text"))
	fmt.Println(strings.HasSuffix("text.rar", "rar"))
	defer fmt.Println(strings.Index("text", "e"))

	file, _ := os.Open("dosyam.txt")
	println(file.Name)

	//-------------------------CONTİNUE & BREAK------------
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		} else if i == 7 {
			break
		}

		println(i, ". Sıra")

	}

	// ----------Normal değerleri çevirme-------------

	x := 12.124124
	e := int(x)
	println(e)

	// ------------String İnte çevirme--------------

	sonuc, _ := strconv.Atoi(girdiAl())
	switchCaseBloklari(sonuc)

	//--------------------------------MAP'S------------------

	myKeys := make(map[string]string)
	myKeys["Canvas"] = "1234"
	myKeys["Udemy"] = "1234as"
	myKeys["Esogubs"] = "1245613"
	myKeys["LinkedIn"] = "1234as23"
	myKeys["Computer"] = "123654789"

	fmt.Println(myKeys)
	//Sadece seçili şehiri getir

	udemy := myKeys["Udemy"]
	fmt.Println("İstenilen şifre: ", udemy)

	//silme
	delete(myKeys, "Udemy")

	//forla yazdırma

	for j, i := range myKeys {
		fmt.Println(j, i)
	}

}
