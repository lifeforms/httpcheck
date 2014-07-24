package urlcheck

import "errors"
import "strconv"

type Tests []Test

type Scenario struct {
	Tests Tests
	Name  string
}

func (s Scenario) Test() error {
	for i, t := range s.Tests {
		err := t.Test()
		if err != nil {
			return errors.New("Step " + strconv.Itoa(i+1) + ": " + err.Error())
		}
	}
	return nil
}
