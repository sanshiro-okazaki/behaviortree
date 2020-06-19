package behaviortree

// 1つ実行
func Action(children []Node) (Status, error) {
	status, err := children[0].Tick()
	if err != nil {
		return Failure, err
	}

	return status, nil
}
