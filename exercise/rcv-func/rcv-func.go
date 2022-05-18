//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Player struct {
	name              string
	health, maxHealth uint
	energy, maxEnergy uint
}

func (player *Player) addHealth(amount uint) {
	player.health += amount
	if player.health > player.maxHealth {
		player.health = player.maxHealth
	}
	fmt.Println(player.name, "add", amount, "health ->", player.health)
}

func (player *Player) applyDamage(amount uint) {
	if player.health-amount > player.health {
		player.health = 0
		fmt.Println(player.name, "are dead")
	}
	player.health -= amount
	fmt.Println(player.name, "received", amount, "damage, reamining health ->", player.health)
}

func (player *Player) addEnergy(amount uint) {
	if player.energy+amount > player.maxHealth {
		player.energy = player.maxEnergy
		fmt.Println(player.name, "already have the max energy")
	}
	player.energy += amount
	fmt.Println(player.name, "add", amount, "energy ->", player.energy)
}

func (player *Player) comsumeEnergy(amount uint) {
	if player.energy-amount > player.energy {
		fmt.Println(player.name, "does not have enough energy")
	}
	player.energy -= amount
	fmt.Println(player.name, "has consumed", amount, "energy, remaining energy -> ", player.energy)
}

func main() {
	player := Player{
		name:      "Kignth",
		health:    500,
		maxHealth: 500,
		energy:    500,
		maxEnergy: 500,
	}

	player.applyDamage(100)
	player.addHealth(50)
	player.comsumeEnergy(400)
	player.addEnergy(100)

}
