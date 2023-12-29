package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Player struct {
	name     string
	position Position
	maxhp    int
	hp       int
	level    int
	Xp       int
	mana     int
	maxmana  int
	fighting bool
}
type Monster struct {
	monsterName string
	hp          int
	xp          int
}
type Position struct {
	x, y int
}

var WorldMap = map[Position]*Monster{
	{1, 0}:  newMonster("Goblin"),
	{-1, 0}: newMonster("Troll"),
	{1, 2}:  newMonster("Imp"),
	{2, 2}:  newMonster("Raccoon"),
	{3, 3}:  newMonster("Hobgoblin"),
}

func printmsg(msg string) {
	for _, r := range msg {
		fmt.Print(string(r))
		time.Sleep(50 * time.Millisecond)
	}
}
func printPlayerPosition(p Player) {
	fmt.Printf("Current position: (%d, %d)\n", p.position.x, p.position.y)
}
func movePlayer(p *Player, direction string) {
	switch direction {
	case "north":
		p.position.y++
	case "south":
		p.position.y--
	case "east":
		p.position.x++
	case "west":
		p.position.x--
	default:
		fmt.Println("Invalid direction. Try again.")
	}
}
func newPlayer(name string) *Player {
	p := Player{name: name}
	p.maxhp = 100
	p.hp = 100
	p.level = 1
	p.Xp = 0
	p.maxmana = 20
	p.mana = 20
	p.position = Position{0, 0}
	return &p
}
func newMonster(monsterName string) *Monster {
	m := Monster{monsterName: monsterName}
	m.hp = 30
	m.xp = 5
	return &m
}
func (p *Player) Attack(m *Monster) {
	xpIntoString := strconv.Itoa(m.xp)

	var attackValue int = p.level * 10
	attackValueIntoString := strconv.Itoa(attackValue)
	if m.hp > 0 && p.hp > 0 {
		printmsg("You attack the monster!")
		fmt.Println("You deal " + attackValueIntoString + " damage!")
		m.hp = m.hp - attackValue
		if m.hp <= 0 {
			printmsg("Monster is dead!, You've gained " + xpIntoString + " XP!")
			fmt.Println("")
			p.Xp = p.Xp + m.xp
		}
	}
}
func Heal(p *Player) {
	healValue := 5 * p.level
	healIntoString := strconv.Itoa(healValue)
	if p.hp > 0 && p.mana >= 5 {
		p.hp = p.hp + healValue
		printmsg("You've healed yourself for " + healIntoString + " HP!")
		p.mana = p.mana - 5
	} else if p.mana < 5 {
		printmsg("You have no mana!")
	}

}
func (m *Monster) MonsterAttack(p *Player) {
	var monsterAttack int = rand.Intn(10) + 1
	if m.hp > 0 {
		fmt.Println("Hero's HP: ")
		fmt.Print(p.hp)
		fmt.Println("")
		fmt.Println("Hero's mana: ")
		fmt.Print(p.mana)
		fmt.Println("The monster attacks!")
		p.hp = p.hp - monsterAttack
	}

}
func fight(p *Player, m *Monster) {

	xpThreshhold := 10 + (p.level * 5)
	for p.hp > 0 {
		realUserInput := ""
		fmt.Println("What would you like to do?: ")
		fmt.Println(`1) Attack
2) Heal (5 mana)`)

		fmt.Scan(&realUserInput)
		userInput := strings.TrimSpace(strings.ToLower(realUserInput))
		switch userInput {
		case "attack":
			p.Attack(m)
			m.MonsterAttack(p)
		case "heal":
			Heal(p)
		default:
			fmt.Println("That's not a valid input!")
		}
		if p.Xp >= xpThreshhold {
			printmsg("You've levelled up!")
			p.level++
			p.maxhp = p.maxhp + (5 * p.level)
			p.maxmana = p.maxmana + (5 * p.level)
			p.hp = 100 + (5 * p.level)
			p.mana = 20 + (5 * p.level)
			p.Xp = 0
			printmsg("Hero's new HP: " + strconv.Itoa(p.maxhp))
			printmsg("Hero's new mana pool: " + strconv.Itoa(p.maxmana))
		}
		if m.hp <= 0 {
			break
		}
	}

}
func travel(p *Player) {
	for {
		printPlayerPosition(*p)
		realUserInput := ""
		fmt.Print("Enter a direction to travel (north, south, east, west): ")
		var userInput string
		fmt.Scan(&realUserInput)
		userInput = strings.TrimSpace(strings.ToLower(realUserInput))
		if userInput == "quit" {
			fmt.Println("Goodbye!")
			break
		}

		movePlayer(p, userInput)

		// Check if there's a monster at the current position
		if monster, ok := WorldMap[p.position]; ok {
			printmsg("A wild " + monster.monsterName + " appears!\n")
			fight(p, monster)
			delete(WorldMap, p.position)
		}
	}
}
func main() {
	var playerName string
	printmsg("Welcome to the game, choose your character's name: ")
	fmt.Scan(&playerName)
	player := newPlayer(playerName)
	travel(player)

}
