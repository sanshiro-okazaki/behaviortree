package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	. "github.com/10antz-inc/sanshiro-okazaki/behaviortree/behaviortree"
)

func main() {
	s := os.Getenv("SKILL_A")
	skillA, err := strconv.ParseFloat(s, 64)
	if err != nil {
		err = errors.New("error: SKILL_A value is required")
		panic(err)
	}

	player := NewPlayer(150, skillA, 1.0)

	enemy := NewEnemy(120)

	err = NewTicker(
		New(
			Sequencer,

			// 出発
			New(
				Action,
				Departure(),
			),

			// 敵に寄る
			New(
				Condition,
				player.CanGetClose(),
				player.GetClose(),
			),

			// 友達を呼ぶ
			New(
				Parallel,
				player.CallFriend1(),
				player.CallFriend2(),
			),

			// スキル発動
			Repeater(
				2,
				New(
					Selector,
					New(
						Condition,
						player.DrawSkillA(),
						player.RunSkillA(enemy),
					),
					New(
						Condition,
						player.DrawSkillB(),
						player.RunSkillB(enemy),
					),
				),
			),

			// 終了判定
			New(
				Selector,
				New(
					Condition,
					enemy.Dead(),
					END1(),
				),
				New(
					Condition,
					enemy.Alive(),
					END2(),
				),
			),
		),
	)

	fmt.Printf("%#v %#v", player, enemy)
	if err != nil {
		panic(err)
	}
}

func Departure() Node {
	return New(
		func(children []Node) (Status, error) {
			fmt.Printf("%s\n", "出発")
			return Success, nil
		},
	)
}

func END1() Node {
	return New(
		func(children []Node) (Status, error) {
			fmt.Printf("%s\n", "終了1")
			return Success, nil
		},
	)
}

func END2() Node {
	return New(
		func(children []Node) (Status, error) {
			fmt.Printf("%s\n", "終了2")
			return Success, nil
		},
	)
}
