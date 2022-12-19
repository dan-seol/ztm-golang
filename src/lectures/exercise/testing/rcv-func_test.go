// --Summary:
//
//	Copy your rcv-func solution to this directory and write unit tests.
//
// --Requirements:
// * Write unit tests that ensure:
//   - Health & energy can not go above their maximums
//   - Health & energy can not go below 0
//   - If any of your  tests fail, make the necessary corrections
//     in the copy of your rcv-func solution file.
//
// --Notes:
// * Use `go test -v ./exercise/testing` to run these specific tests
package rcv_func

import "testing"

func TestHealthAndEnergyUpperBounds(t *testing.T) {
	player := Player{name: "Dan", maxHealth: 200, health: 150, maxEnergy: 200, energy: 80}
	player.modifyHealth(+100)
	player.modifyEnergy(+100)

	if player.health > player.maxHealth {
		t.Errorf("The health: %v cannot go beyond maxHealth: %v", player.health, player.maxHealth)
	}

	if player.energy > player.maxEnergy {
		t.Errorf("The energy: %v cannot go beyond maxEnergy: %v", player.energy, player.maxEnergy)
	}
}

func TestHealthAndEnergyLowerBounds(t *testing.T) {
	player := Player{name: "Dan", maxHealth: 200, health: 200, maxEnergy: 200, energy: 200}
	player.modifyHealth(-300)
	player.modifyEnergy(-500)

	if player.health < 0 {
		t.Fatalf("The health: %v cannot go below 0", player.health)
	}

	if player.energy < 0 {
		t.Fatalf("The energy: %v cannot go below 0", player.energy)
	}
}
