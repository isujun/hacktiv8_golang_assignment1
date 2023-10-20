package main

import (
	"fmt"
	"os"
	"strings"
)

type Person struct {
	ID        int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

// FindPersonByName searches for a person by name and returns their details.
func FindPersonByName(name string, people []Person) (*Person, bool) {
	for _, person := range people {
		if strings.EqualFold(person.Nama, name) {
			return &person, true
		}
	}
	return nil, false
}

func main() {
	people := []Person{
		{
			ID:        0,
			Nama:      "Thomas",
			Alamat:    "Kota A",
			Pekerjaan: "Programmer",
			Alasan:    "Suka makan ice cream",
		},
		{
			ID:        1,
			Nama:      "Ilham",
			Alamat:    "Kota B",
			Pekerjaan: "Programmer",
			Alasan:    "Suka dunia tech-tech",
		},
		{
			ID:        2,
			Nama:      "Johnny",
			Alamat:    "Kota C",
			Pekerjaan: "HR",
			Alasan:    "Lagu Johnny yes papa",
		},
		{
			ID:        3,
			Nama:      "Parker",
			Alamat:    "Kota D",
			Pekerjaan: "Kang Parkir",
			Alasan:    "Cuan gede",
		},
		{
			ID:        4,
			Nama:      "Genjih",
			Alamat:    "Kota E",
			Pekerjaan: "Kang Siomay",
			Alasan:    "Tiada ada kata seindah \"bang beli bang\"",
		},
	}

	if len(os.Args) < 2 {
		fmt.Println("input nama yang mau diabsen !")
		return
	}

	nameToFind := os.Args[1]
	person, found := FindPersonByName(nameToFind, people)
	if found {
		fmt.Println("ID:", person.ID)
		fmt.Println("Nama:", person.Nama)
		fmt.Println("Alamat:", person.Alamat)
		fmt.Println("Pekerjaan:", person.Pekerjaan)
		fmt.Println("Alasan:", person.Alasan)
	} else {
		fmt.Printf("Nama %s tidak ditemukan!\n", nameToFind)
	}
}
