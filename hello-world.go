package main

type Player struct {
	name  string
	hp    int
	level int
	Xp    int
}
type Monster struct {
	name string
	hp   int
	xp   int
}

func newPlayer(name string) *Player {
	p := Player{name: name}
	p.hp = 100
	p.level = 1
	p.Xp = 0
	return &p
}
func newMonster(name string) *Monster {
	m := Monster{name: name}
	m.hp = 30
	m.xp = 5
	return &m
}

func main() {

}
