package main

import (
	"fmt"
	"math/rand"
	"time"

	. "github.com/10antz-inc/sanshiro-okazaki/behaviortree/behaviortree"
)

type Player struct {
	HP     int
	SkillA float64
	SkillB float64
}

func NewPlayer(hp int, sa, sb float64) *Player {
	return &Player{
		HP:     hp,
		SkillA: sa,
		SkillB: sb,
	}
}

func (p *Player) CanGetClose() Node {
	return New(
		func(children []Node) (Status, error) {
			if p.HP >= 100 {
				return Success, nil
			}
			return Failure, nil
		},
	)
}

func (p *Player) GetClose() Node {
	return New(
		func(children []Node) (Status, error) {
			fmt.Printf("%s\n", "敵に寄る")
			return Success, nil
		},
	)
}

func (p *Player) CallFriend1() Node {
	return New(
		func(children []Node) (Status, error) {
			fmt.Printf("%s\n", "友達1")
			return Success, nil
		},
	)
}

func (p *Player) CallFriend2() Node {
	return New(
		func(children []Node) (Status, error) {
			fmt.Printf("%s\n", "友達2")
			return Success, nil
		},
	)
}

func (p *Player) RunSkillA(e *Enemy) Node {
	return New(
		func(children []Node) (Status, error) {
			e.HP -= 50
			fmt.Printf("%s \n", "スキルA発動")
			return Success, nil
		},
	)
}

func (p *Player) RunSkillB(e *Enemy) Node {
	return New(
		func(children []Node) (Status, error) {
			e.HP -= 60
			fmt.Printf("%s \n", "スキルB発動")
			return Success, nil
		},
	)
}

func (p *Player) DrawSkillA() Node {
	return New(
		func(children []Node) (Status, error) {
			rand.Seed(time.Now().UnixNano())
			r := rand.Float64()
			if r > p.SkillA {
				return Failure, nil
			}
			return Success, nil
		},
	)
}

func (p *Player) DrawSkillB() Node {
	return New(
		func(children []Node) (Status, error) {
			rand.Seed(time.Now().UnixNano())
			r := rand.Float64()
			if r > p.SkillB {
				return Failure, nil
			}
			return Success, nil
		},
	)
}
