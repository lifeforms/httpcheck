package urlcheck

import "errors"

// A Server has a number of Scenarios to test.
type Server struct {
	Name      string "server"
	Scenarios []Scenario
}

// Test runs tests on the scenarios and tests for this server.
// It returns an error if one or more scenarios/tests has errors, or nil otherwise.
// In case there are multiple errors, the error contains the concatenated messages.
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
		if server.Name != "" {
			errorstr = "Server " + server.Name + ": "
		}
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
