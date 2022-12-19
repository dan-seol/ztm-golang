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

package rcv_func

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
