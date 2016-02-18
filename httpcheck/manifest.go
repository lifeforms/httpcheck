package httpcheck

import (
	"errors"
	"gopkg.in/yaml.v1"
	"regexp"
)

// Manifest contains one or more Servers, each having some Scenarios.
type Manifest []Server

// SetBaseURL walks through all tests and prefixes any relative URLs (e.g. '/') with a base URL.
func (m *Manifest) SetBaseURL(baseurl string) {
	scheme := regexp.MustCompile("^https?:")
	for i := range *m {
		for j := range (*m)[i].Scenarios {
			for k := range (*m)[i].Scenarios[j].Tests {
				match := scheme.MatchString((*m)[i].Scenarios[j].Tests[k].Url)
				if !match {
					//	This is a relative URL, prefix it with baseurl.
					//	When glueing the URLs together, remove a double '/' to make prettier URLs.
					//	To have a test that starts with '//', add it to the test URL explicitly.
					var url string
					if baseurl[len(baseurl)-1] == '/' && len((*m)[i].Scenarios[j].Tests[k].Url) > 0 &&
						(*m)[i].Scenarios[j].Tests[k].Url[0] == '/' {
						url = baseurl[0:len(baseurl)-1] + (*m)[i].Scenarios[j].Tests[k].Url
					} else {
						url = baseurl + (*m)[i].Scenarios[j].Tests[k].Url
					}
					(*m)[i].Scenarios[j].Tests[k].Url = url
				}
			}
		}
	}
}

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
		defer close(c)
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
			if errorcount > 0 {
				errorstr += "\n"
			}
			errorstr += err.Error()
			errorcount++
		}
	}

	if errorcount > 0 {
		return errors.New(errorstr)
	} else {
		return nil
	}
}

// FromYAML parses YAML input and returns a manifest.
func FromYAML(y []byte) (Manifest, error) {
	var manifest Manifest
	err := yaml.Unmarshal(y, &manifest)
	if err != nil {
		return nil, err
	}
	return manifest, nil
}
