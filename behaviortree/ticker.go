package behaviortree

import (
	"errors"
)

func NewTicker(node Node) error {
	tick, children := node()
	if tick == nil {
		return errors.New("error: cannot with a nil tick")
	}

	status, err := tick(children)
	if err == nil && status == Failure {
		err = errors.New("error: exit on failure")
	}
	return nil
}
