package urlcheck

import "errors"
import "strconv"

type Scenario []Test

func (s Scenario) Test() error {
	for i, t := range s {
		err := t.Test()
		if err != nil {
			return errors.New("Step " + strconv.Itoa(i+1) + ": " + err.Error())
		}
	}
	return nil
}
