/*1. Napiši strukturu koja predstavlja pravokutnik i sadrži polja za širinu i visinu. Kreiraj metode za
izračunavanje površine i opsega pravokutnika, metode moraju biti vezane za novo kreiranu strukturu Pravokutnika.
U main funkciji inicijalizirajte jedan pravokutnik, te pozovite iznad kreirane metode za računanje površine i
opsega.*/

package main

import "fmt"

type Pravokutnik struct {
	sirina float64
	visina float64
}

func (p Pravokutnik) Povrsina() float64 {
	return p.sirina * p.visina
}

func (p Pravokutnik) Opseg() float64 {
	return 2 * (p.sirina + p.visina)
}

func main() {
	pravokutnik := Pravokutnik{sirina: 10, visina: 15}

	povrsina := pravokutnik.Povrsina()
	opseg := pravokutnik.Opseg()

	fmt.Println("Površina = ", povrsina)
	fmt.Println("Opseg = ", opseg)
}
