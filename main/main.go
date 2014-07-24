package main

import "fmt"
import "github.com/lifeforms/urlcheck/urlcheck"

func main() {
	sc := urlcheck.Scenario{
		urlcheck.Test{Url: "http://sim.dt.lfms.nl/", Content: "sim"},
		urlcheck.Test{Url: "http://www.lifeforms.nl/", Content: "DEHUMANIZE"},
		urlcheck.Test{Url: "https://www.lifeforms.nl/", Content: "DEHUMANIZE"},
	}
	err := sc.Test()
	fmt.Println("Test result:", err)
}
