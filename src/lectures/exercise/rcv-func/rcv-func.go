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

import (
	"fmt"
	"math"
)

type Player struct {
	name      string
	maxHealth int
	health    int
	maxEnergy int
	energy    int
}

type Existence Player

func (existence Existence) printStatus() {
	fmt.Println("Existence")
	fmt.Println("{name =", existence.name)
	fmt.Println(" maxHealth =", existence.maxHealth)
	fmt.Println(" health =", existence.health)
	fmt.Println(" maxEnergy =", existence.maxEnergy)
	fmt.Println(" energy =", existence.energy, "}")
}

func (player Player) printStatus() {
	fmt.Println("Player")
	fmt.Println("{name =", player.name)
	fmt.Println(" maxHealth =", player.maxHealth)
	fmt.Println(" health =", player.health)
	fmt.Println(" maxEnergy =", player.maxEnergy)
	fmt.Println(" energy =", player.energy, "}")
}

func (player *Player) modifyHealth(diff int) {
	newValue := player.health + diff
	player.health = int(math.Max(0, math.Min(float64(newValue), float64(player.maxHealth))))
	player.printStatus()
}

func (player *Player) modifyEnergy(diff int) {
	newValue := player.energy + diff
	player.energy = int(math.Max(0, math.Min(float64(newValue), float64(player.maxEnergy))))
	player.printStatus()
}

func (existence *Existence) modifyHealth(diff int) {
	newValue := existence.health + diff
	existence.health = int(math.Max(0, math.Min(float64(newValue), float64(existence.maxHealth))))
	existence.printStatus()
}

func (existence *Existence) modifyEnergy(diff int) {
	newValue := existence.energy + diff
	existence.energy = int(math.Max(0, math.Min(float64(newValue), float64(existence.maxEnergy))))
	existence.printStatus()
}

func main() {
	boss := Existence{name: "Nenuphar, the Grand Inquisitor of Vedic Temples", maxHealth: 300, health: 300, maxEnergy: 800, energy: 800}
	boss.printStatus()
	player := Player{name: "Dan", maxHealth: 200, health: 200, maxEnergy: 200, energy: 200}
	player.printStatus()
	fmt.Println("Field Boss:", boss.name, "has started reciting the Vedic scriptures...")
	fmt.Println("Player:", player.name, ", your soul is aching!")
	player.modifyEnergy(-60)
	fmt.Println("Player:", player.name, "has offered his blood to an evil spirit in order to gain energy! Ew")
	player.modifyHealth(-50)
	fmt.Println("Player:", player.name, ", the spirit has accepted your offer and gave you energy !")
	player.modifyEnergy(+60)
	fmt.Println("Field Boss:", boss.name, "has made a call to arms to the daevas of the Temples...")
	fmt.Println("Player:", player.name, ", the spirit has opened a portal to come and ward you from the Holy Powers!")
	spirit := Existence{name: "Botchulaz, Nurgle's Great Unclean One", maxHealth: 1700, health: 1700, maxEnergy: 1000, energy: 950}
	fmt.Println("Player:", player.name, ", the spirit manifested its presence! It turned out to be", spirit.name)
	spirit.printStatus()
	fmt.Println("Player:", player.name, ", the Daevas have arrived...")
	daevas := []Existence{
		{name: "Holy Daeva", maxHealth: 500, health: 500, maxEnergy: 400, energy: 400},
		{name: "Lesser Daeva 1", maxHealth: 400, health: 400, maxEnergy: 200, energy: 200},
		{name: "Lesser Daeva 2", maxHealth: 400, health: 400, maxEnergy: 200, energy: 200}}

	for _, daeva := range daevas {
		daeva.printStatus()
	}
	fmt.Println("Field Boss:", boss.name, "smites", spirit.name, "with divine power!")
	boss.modifyEnergy(-400)
	spirit.modifyHealth(-600)
	fmt.Println("Chaos Daemon:", spirit.name, "corrupts the body and soul of", boss.name, " with its Plague Sword!")
	boss.modifyHealth(-200)
	fmt.Println("Chaos Daemon:", spirit.name, "the plague further rots", boss.name, "to the core!")
	boss.modifyHealth(-100)
	fmt.Println("Field Boss:", boss.name, "dies, but the daevas look even angrier!")
	for _, daeva := range daevas {
		fmt.Println("Daeva:", daeva.name, "pierces through", spirit.name, "with flaming arrows!")
		spirit.modifyHealth(-300)
	}
	fmt.Println("Player:", player.name, ", counts to three and throws a plutonium grenade to daevas")
	for i, daeva := range daevas {
		fmt.Println("Daeva:", daeva.name, "gets swept by the pale blue glow impact!")
		daevas[i].modifyHealth(-350)
	}
}
