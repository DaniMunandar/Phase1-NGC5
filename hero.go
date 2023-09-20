package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Struct Hero 1
type Weapon struct {
	Attack int
}

type Hero struct {
	Name           string
	BaseAttack     int
	Defence        int
	CriticalDamage int
	HealthPoint    int
	Weapon         Weapon
}

func (h *Hero) CountDamage() int {
	rand.Seed(time.Now().UnixNano()) // Mengatur seed untuk menghasilkan nilai acak yang berbeda setiap kali program dijalankan

	// Menentukan apakah critical damage terjadi (50:50 peluang)
	critical := rand.Intn(2) == 1

	totalDamage := h.BaseAttack + h.Weapon.Attack

	// Jika critical damage terjadi, tambahkan ke total damage
	if critical {
		totalDamage += h.CriticalDamage
	}

	return totalDamage
}

func (h *Hero) isAttackedBy(attacker Hero) {
	totalDamageReceived := attacker.CountDamage() - h.Defence

	if totalDamageReceived < 0 {
		totalDamageReceived = 0
	}

	h.HealthPoint -= totalDamageReceived
}

// Fungsi untuk simulasi pertempuran antara dua Hero
func Battle(attacker Hero, defender Hero) {
	fmt.Printf("%s menyerang %s!\n", attacker.Name, defender.Name)
	fmt.Printf("%s's HealthPoint sebelum serangan: %d\n", defender.Name, defender.HealthPoint)

	defender.isAttackedBy(attacker)

	fmt.Printf("%s menerima %d damage!\n", defender.Name, attacker.CountDamage())
	fmt.Printf("%s's HealthPoint setelah serangan: %d\n\n", defender.Name, defender.HealthPoint)
}
