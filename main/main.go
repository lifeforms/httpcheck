package main

import "fmt"
import "github.com/lifeforms/urlcheck/urlcheck"

func main() {
	server := urlcheck.Server{
		Name: "tau",
		Scenarios: []urlcheck.Scenario{
			urlcheck.Scenario{
				Name: "lifeforms",
				Tests: urlcheck.Tests{
					urlcheck.Test{Url: "http://www.lifeforms.nl/", Content: "DEHUMANIZE"},
					urlcheck.Test{Url: "https://www.lifeforms.nl/", Content: "DEHUMANIZE"},
				},
			},
			urlcheck.Scenario{
				Name: "ionica",
				Tests: urlcheck.Tests{
					urlcheck.Test{Url: "http://ionica.nl/", Content: "Ionica Smeets"},
				},
			},
		},
	}
	err := server.Test()
	if err != nil {
		fmt.Println("Failures:", err)
	} else {
		fmt.Println("OK")
	}
}
