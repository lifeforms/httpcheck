package main

import (
	"flag"
	"fmt"
	"github.com/lifeforms/urlcheck/urlcheck"
	"io/ioutil"
	"os"
)

func parseArgs() (manifestfile string, verbose bool) {
	flagi := flag.String("i", "urls.yaml", "Input file with YAML manifest")
	flagv := flag.Bool("v", false, "Verbose, prints the result of each test")
	flag.Parse()

	manifestfile = *flagi
	verbose = *flagv
	return
}

func readManifest(manifestfile string) (manifest urlcheck.Manifest, err error) {
	y, err := ioutil.ReadFile(manifestfile)
	if err != nil {
		return nil, err
	}

	manifest, err = urlcheck.FromYAML(y)
	if err != nil {
		return nil, err
	}

	return
}

func main() {
	manifestfile, verbose := parseArgs()

	urlcheck.Verbose = verbose
	manifest, err := readManifest(manifestfile)
	if err == nil {
		err = manifest.Test()
	}

	if err != nil {
		fmt.Println("Failures:", err)
		os.Exit(1)
	} else {
		fmt.Println("OK")
	}
}
