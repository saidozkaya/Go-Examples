package main

import (
	"os"
	"log"
)

var (
	newFile *os.File
	err     error
)

func main() {
	newFile, err = os.Create("demo.txt")
	iff err!= nil{
		log.Fatal(err)
	}
}


type Human struct {
	Firstname string
	Surname   string
	Age       int
}

func newHuman() *Human {
	h := new(Human)
	return h
}
func newHumanPar(firstName, lastName string, age int) *Human {
	h := new(Human)
	h.Firstname = firstName
	h.Surname = lastName
	h.Age = age
	return h
}

func createMethods() {
	//1.Yol
	human1 := Human{Firstname: "Toms", Surname: "Richer", Age: 22}
	fmt.Println(human1.Firstname)
	human1.Firstname = "Toms Rich"
	fmt.Println(human1)
	//2.Yol
	human2 := new(Human)
	human2.Firstname = "Graphicso"
	human2.Surname = "Picture"
	human2.Age = 16
	fmt.Println(human2)
	//3.Yol
	human3 := newHuman()
	human3.Firstname = "Hakan"
	human3.Surname = "Shelby"
	human3.Age = 65
	fmt.Println(human3)
	//4. Yol
	human4 := newHumanPar("Ahmet", "Gelby", 23)
	fmt.Println(human4)
	//Hepsini tek tek okuma
	oku := human4.Firstname + " " + human4.Surname + " " + strconv.Itoa(human4.Age)
	fmt.Println(oku)
}

