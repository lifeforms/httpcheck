package urlcheck

import (
	"errors"
	"gopkg.in/yaml.v1"
)

// Manifest contains one or more Servers, each having some Scenarios.
type Manifest []Server

// Test runs tests on all servers in the manifest.
// It returns an error if one or more server has errors, or nil otherwise.
// In case there are multiple errors, the error contains the concatenated messages.
func (m Manifest) Test() error {
	if len(m) == 0 {
		return errors.New("Manifest is empty")
	}

	// Start goroutine for every server.Test() call
	// The result of server.Test() is an error or nil. This is passed through a channel.
	var testresults []chan error
	for _, server := range m {
		c := make(chan error)
		testresults = append(testresults, c)
		go func(server Server) {
			c <- server.Test()
		}(server)
	}

	// Read from every channel to collect all errors returned.
	// This blocks until every channel has something to receive so all tests are done.
	errorcount := 0
	errorstr := ""
	for _, c := range testresults {
		err := <-c
		if err != nil {
			errorstr += err.Error()
			if errorcount > 0 {
				errorstr += "\n"
			}
			errorcount++
		}
	}

	if errorcount > 0 {
		return errors.New(errorstr)
	} else {
		return nil
	}
}

func FromYAML(y []byte) (Manifest, error) {
	var manifest Manifest
	err := yaml.Unmarshal(y, &manifest)
	if err != nil {
		return nil, err
	}
	return manifest, nil
}
