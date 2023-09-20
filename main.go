package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
	Job  string
}

func (p *Person) GetInfo() {
	fmt.Printf("Name: %s\nAge: %d\nJob: %s\n", p.Name, p.Age, p.Job)
}

func (p *Person) AddYear() {
	p.Age++
	if p.Age >= 50 {
		p.Job = "Retired"
	}
}

func main() {
	// Membuat sebuah instance Person
	person := &Person{
		Name: "Bambang",
		Age:  45,
		Job:  "Gambler",
	}

	// Menampilkan informasi awal
	fmt.Println("Informasi awal:")
	person.GetInfo()

	// Menambahkan usia 5 kali
	for i := 0; i < 5; i++ {
		person.AddYear()
	}

	// Menampilkan informasi setelah penambahan usia
	fmt.Println("\nInformasi setelah penambahan usia:")
	person.GetInfo()

	// Membuat sebuah slice of Person
	persons := []Person{
		{Name: "Bambang", Age: 20, Job: "Gambler"},
		{Name: "Citra", Age: 30, Job: "Engineer"},
		{Name: "Dewi", Age: 25, Job: "Teacher"},
	}

	// Menampilkan informasi setiap Person dalam slice
	for _, person := range persons {
		person.GetInfo()
	}

	// Membuat sebuah instance Hero
	hero := Hero{
		Name:           "Superman",
		BaseAttack:     50,
		Defence:        10,
		CriticalDamage: 30,
		HealthPoint:    100,
		Weapon: Weapon{
			Attack: 20,
		},
	}

	// Menghitung total damage saat Hero menyerang
	totalDamage := hero.CountDamage()
	fmt.Printf("%s menyerang dan memberikan total damage %d.\n", hero.Name, totalDamage)

	superman := Hero{
		Name:           "Superman",
		BaseAttack:     50,
		Defence:        10,
		CriticalDamage: 30,
		HealthPoint:    100,
		Weapon: Weapon{
			Attack: 20,
		},
	}

	batman := Hero{
		Name:           "Batman",
		BaseAttack:     40,
		Defence:        15,
		CriticalDamage: 20,
		HealthPoint:    90,
		Weapon: Weapon{
			Attack: 25,
		},
	}

	Battle(superman, batman)
	Battle(batman, superman)
}
