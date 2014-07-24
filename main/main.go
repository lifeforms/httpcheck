package main

import "fmt"
import "github.com/lifeforms/urlcheck/urlcheck"

func main() {
	sim := urlcheck.Test{Url: "http://sim.dt.lfms.nl/"}
	lfms := urlcheck.Test{Url: "https://www.lifeforms.nl/"}
	sc := urlcheck.Scenario{sim, lfms}
	err := sc.Test()
	fmt.Println("Test result:", err)
}
