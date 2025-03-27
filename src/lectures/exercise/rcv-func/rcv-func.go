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

func NewPlayer(name string, maxHealth, maxEnergy uint) *Player {
	return &Player{
		name:      name,
		health:    maxHealth,
		maxHealth: maxHealth,
		energy:    maxEnergy,
		maxEnergy: maxEnergy,
	}
}

// TakeDamage is a receiver function to reduce player's health
func (p *Player) TakeDamage(amount uint) {
	// Reduce health, ensuring it doesn't go below 0
	oldHealth := p.health
	if amount > oldHealth {
		p.health = 0
	} else {
		p.health -= amount
	}

	fmt.Printf("%s takes %d damage. Health reduced from %d to %d\n",
		p.name, amount, oldHealth, p.health)
}

// Heal is a receiver function to increase player's health
func (p *Player) Heal(amount uint) {
	// Increase health, ensuring it doesn't exceed max health
	oldHealth := p.health
	p.health += amount
	if p.health > p.maxHealth {
		p.health = p.maxHealth
	}

	fmt.Printf("%s heals for %d. Health increased from %d to %d\n",
		p.name, amount, oldHealth, p.health)
}

// ConsumeEnergy is a receiver function to reduce player's energy
func (p *Player) ConsumeEnergy(amount uint) {
	// Reduce energy, ensuring it doesn't go below 0
	oldEnergy := p.energy
	if amount > oldEnergy {
		p.energy = 0
	} else {
		p.energy -= amount
	}

	fmt.Printf("%s consumes %d energy. Energy reduced from %d to %d\n",
		p.name, amount, oldEnergy, p.energy)
}

// RestoreEnergy is a receiver function to increase player's energy
func (p *Player) RestoreEnergy(amount uint) {
	// Increase energy, ensuring it doesn't exceed max energy
	oldEnergy := p.energy
	p.energy += amount
	if p.energy > p.maxEnergy {
		p.energy = p.maxEnergy
	}

	fmt.Printf("%s restores %d energy. Energy increased from %d to %d\n",
		p.name, amount, oldEnergy, p.energy)
}

func main() {
	// Create a new player
	player := NewPlayer("Adventurer", 100, 50)

	// Demonstrate health modifications
	player.TakeDamage(30)
	player.Heal(20)

	// Demonstrate energy modifications
	player.ConsumeEnergy(25)
	player.RestoreEnergy(15)
}
