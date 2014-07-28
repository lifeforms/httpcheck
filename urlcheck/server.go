package urlcheck

import "errors"

type Server struct {
	Name      string
	Scenarios []Scenario
}

func (server Server) Test() error {
	var allerrors []error
	for _, scenario := range server.Scenarios {
		err := scenario.Test()
		if err != nil {
			allerrors = append(allerrors, err)
		}
	}

	if len(allerrors) > 0 {
		errorstr := ""
		for i, err := range allerrors {
			if i > 0 {
				errorstr += "\n"
			}
			errorstr += err.Error()
		}
		return errors.New(errorstr)
	}
	return nil
}
