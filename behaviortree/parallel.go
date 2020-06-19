package behaviortree

// 並列実行したchildrenが全て成功したらSuccessを返す
func Parallel(children []Node) (Status, error) {
	s := make(chan Status, len(children))
	e := make(chan error, len(children))
	done := make(chan struct{})

	go func(children []Node, s chan<- Status, e chan<- error) {
		for _, c := range children {
			status, err := c.Tick()
			if err != nil {
				e <- err
				s <- Failure
			}
			if status == Failure {
				e <- nil
				s <- Failure
			}
			e <- nil
			s <- Success
		}

		close(s)
		close(e)
		close(done)
	}(children, s, e)

	for {
		select {
		case err := <-e:
			if err != nil {
				return Failure, err
			}
		case status := <-s:
			if status == Failure {
				return Failure, nil
			}
		case <-done:
			return Success, nil
		}
	}
}
