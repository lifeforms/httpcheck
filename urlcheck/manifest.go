package urlcheck

import "errors"

type Manifest []Server

func (m Manifest) Test() error {
	var allerrors []error
	for _, server := range m {
		err := server.Test()
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
