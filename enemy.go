package main

import (
	. "github.com/10antz-inc/sanshiro-okazaki/behaviortree/behaviortree"
)

type Enemy struct {
	HP int
}

func NewEnemy(hp int) *Enemy {
	return &Enemy{
		HP: hp,
	}
}

func (e *Enemy) Dead() Node {
	return New(
		func(children []Node) (Status, error) {
			if e.HP <= 0 {
				return Success, nil
			}
			return Failure, nil
		},
	)
}

func (e *Enemy) Alive() Node {
	return New(
		func(children []Node) (Status, error) {
			if e.HP >= 0 {
				return Success, nil
			}
			return Failure, nil
		},
	)
}
