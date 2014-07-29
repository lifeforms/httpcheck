package main

import "fmt"
import "io/ioutil"
import "github.com/lifeforms/urlcheck/urlcheck"

func main() {
	y, err := ioutil.ReadFile("urls.yaml")
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	manifest, err := urlcheck.FromYAML(y)
	if err != nil {
		fmt.Println("Error parsing:", err)
		return
	}

	err = manifest.Test()
	if err != nil {
		fmt.Println("Failures:", err)
	} else {
		fmt.Println("OK")
	}
}
