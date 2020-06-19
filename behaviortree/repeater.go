package behaviortree

import "errors"

// n回繰り返し
func Repeater(n int, node Node) Node {

	return New(
		func(children []Node) (Status, error) {
			for i := 0; i < n; i++ {
				tick, children := node()
				if tick == nil {
					return Failure, errors.New("error: cannot with a nil tick")
				}

				status, err := tick(children)
				if err == nil && status == Failure {
					return Failure, nil
				}
			}
			return Success, nil
		},
	)
}
