package httpcheck

import (
	"errors"
	"strconv"
)

type Tests []Test

// A Scenario describes multiple tests executed in-order.
// Cookies are preserved within a scenario, so tests can depend on earlier tests.
type Scenario struct {
	Name  string
	Tests Tests "test,flow"
}

// Test runs the tests in this scenario in order, stopping at the first error.
// It returns an error if one of the tests in the scenario has errors, or nil otherwise.
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
