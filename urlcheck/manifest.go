package urlcheck

import "errors"
import "gopkg.in/yaml.v1"

// Manifest contains multiple Servers, each having Tests and Scenarios.
type Manifest []Server

// Test runs tests on all servers in the manifest.
// It returns an error if one or more server has errors, or nil otherwise.
// In case there are multiple errors, the error contains the concatenated messages.
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

func FromYAML(y []byte) (Manifest, error) {
	var manifest Manifest
	err := yaml.Unmarshal(y, &manifest)
	if err != nil {
		return nil, err
	}
	return manifest, nil
}
