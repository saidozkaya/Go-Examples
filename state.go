package main

import (
	"fmt"
	"strings"
)

type heck []string

func (d heck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
func deal(d heck, deger int) (heck, heck) {
	return d[:deger], d[deger:]
}
func newheck() heck {

	family := heck{}

	names := []string{"John", "Robbert", "Cuci"}
	surnames := []string{"Summary", "Susamy", "Susy"}

	for _, name := range names {
		for _, surname := range surnames {
			family = append(family, name+" of "+surname)

		}
	}
	return family
}
func (d heck) toString() string {
	return strings.Join([]string(d), " ,")
}
