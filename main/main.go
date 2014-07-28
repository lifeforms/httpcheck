package main

import "fmt"
import "github.com/lifeforms/urlcheck/urlcheck"

func main() {
	manifest := urlcheck.Manifest{
		urlcheck.Server{
			Name: "tau",
			Scenarios: []urlcheck.Tester{
				urlcheck.Test{Url: "http://www.lifeforms.nl/nonexistent", Code: 404},
				urlcheck.Scenario{
					Name: "lifeforms",
					Tests: urlcheck.Tests{
						urlcheck.Test{Url: "http://www.lifeforms.nl/", Content: "DEHUMANIZE"},
						urlcheck.Test{Url: "https://www.lifeforms.nl/", Content: "DEHUMANIZo"},
					},
				},
				urlcheck.Scenario{
					Name: "ionica",
					Tests: urlcheck.Tests{
						urlcheck.Test{Url: "http://ionica.nl/", Content: "Ionica Smeets"},
					},
				},
			},
		},
	}

	err := manifest.Test()
	if err != nil {
		fmt.Println("Failures:", err)
	} else {
		fmt.Println("OK")
	}
}
