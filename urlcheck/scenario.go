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
			return errors.New(s.String() + " step " + strconv.Itoa(i+1) + ": " + err.Error())
		}
	}
	return nil
}

func (s Scenario) String() string {
	if s.Name == "" {
		return "Unnamed scenario"
	} else {
		return "Scenario " + s.Name
	}
}
