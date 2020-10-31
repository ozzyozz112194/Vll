package main

import (
	"fmt"
	"sort"
)

type Persona struct {
	nombre string
	edad   uint64
}

type Bynombre []Persona

func (a Bynombre) Len() int           { return len(a) }
func (a Bynombre) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Bynombre) Less(i, j int) bool { return a[i].nombre > a[j].nombre }

type Byage []Persona

func (a Byage) Len() int           { return len(a) }
func (a Byage) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Byage) Less(i, j int) bool { return a[i].edad < a[j].edad }

func main() {

	ps := []Persona{
		Persona{nombre: "Carlos", edad: 40},
		Persona{nombre: "Braulio", edad: 50},
		Persona{nombre: "Alberto", edad: 20},
	}
	sort.Sort(Bynombre(ps))
	fmt.Println("Por nombre :", ps)

	sort.Sort(Byage(ps))
	fmt.Println("Por edad :", ps)

}
