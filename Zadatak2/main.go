/*2. Napiši strukturu koja predstavlja adresu, koja sadrži polja za grad i ulicu.
Zatim napiši strukturu koja predstavlja osobu, koja sadrži ime, godine i adresu.
Kreiraj metodu koja ispisuje puni opis osobe, uključujući njezinu adresu. (Sva polja) u formatu:
Ime Prezime, 20 godina, živi u Grad, Ulica.*/

package main

import "fmt"

type Adresa struct {
	Grad  string
	Ulica string
}

type Osoba struct {
	Ime    string
	Godine int
	Adresa Adresa
}

func main() {
	osoba := Osoba{"Damir Mihajlović", 24, Adresa{"Mostar", "Jasenica"}}
	osoba.Ispis()
}

func (o Osoba) Ispis() {
	//fmt.Println(o.Ime, ", ", o.Godine, "godina, živi u", o.Adresa.Grad, ", ", o.Adresa.Ulica)
	fmt.Printf("%s, %d, godina, živi u %s, %s", o.Ime, o.Godine, o.Adresa.Grad, o.Adresa.Ulica)
}
