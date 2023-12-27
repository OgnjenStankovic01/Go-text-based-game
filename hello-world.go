package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Player struct {
	name    string
	maxhp   int
	hp      int
	level   int
	Xp      int
	mana    int
	maxmana int
}
type Monster struct {
	monsterName string
	hp          int
	xp          int
}

func printmsg(msg string) {
	for _, r := range msg {
		fmt.Print(string(r))
		time.Sleep(50 * time.Millisecond)
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
	if m.hp > 0 && p.hp > 0 {
		printmsg("You attack the monster!")
		m.hp = m.hp - attackValue
		if m.hp <= 0 {
			printmsg("Monster is dead!, You've gained " + xpIntoString + " XP!")
			fmt.Println("")
			p.Xp = p.Xp + m.xp
			m.hp = 30
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
func runGame(p *Player, m *Monster) {
	xpThreshhold := 10 + (p.level * 5)
	for p.hp > 0 {
		realUserInput := ""
		printmsg("A monster appears before you!")
		fmt.Println("")
		fmt.Println("What would you like to do?: ")
		fmt.Println("1) Attack")
		fmt.Println("2) Heal (5 mana)")
		fmt.Scan(&realUserInput)
		userInput := strings.TrimSpace(strings.ToLower(realUserInput))
		if userInput == "attack" || userInput == "1" {
			p.Attack(m)
			m.MonsterAttack(p)
		} else if userInput == "heal" || userInput == "2" {
			Heal(p)
		} else {
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
	}
}

func main() {
	var playerName string
	printmsg("Welcome to the game, choose your character's name: ")
	fmt.Scan(&playerName)
	player := newPlayer(playerName)
	monster := newMonster("Goblin")
	runGame(player, monster)
}
