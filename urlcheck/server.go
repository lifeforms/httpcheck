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
		for _, err := range allerrors {
			errorstr += err.Error() + "\n"
		}
		return errors.New(errorstr)
	}
	return nil
}
