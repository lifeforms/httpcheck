package urlcheck

import "errors"

type Server struct {
	Name      string
	Scenarios []Tester
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
