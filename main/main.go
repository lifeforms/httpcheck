package main

import (
	"flag"
	"fmt"
	"github.com/lifeforms/urlcheck/urlcheck"
	"io/ioutil"
	"os"
)

func parseArgs() (manifestfile string, timeout uint, verbose bool) {
	flagi := flag.String("i", "urls.yaml", "Input file with YAML manifest")
	flagt := flag.Uint("t", 5, "Timeout for each HTTP request in seconds, 0 for no timeout")
	flagv := flag.Bool("v", false, "Verbose, prints the result of each test")
	flag.Parse()

	manifestfile = *flagi
	timeout = *flagt
	verbose = *flagv
	return
}

func readManifest(manifestfile string) (manifest urlcheck.Manifest, err error) {
	y, err := ioutil.ReadFile(manifestfile)
	if err == nil {
		manifest, err = urlcheck.FromYAML(y)
	}
	return
}

func main() {
	manifestfile, timeout, verbose := parseArgs()

	urlcheck.Timeout = timeout
	urlcheck.Verbose = verbose
	manifest, err := readManifest(manifestfile)
	if err == nil {
		err = manifest.Test()
	}

	if err != nil {
		fmt.Println("Failures:", err)
		os.Exit(1)
	} else {
		if verbose {
			fmt.Println("OK")
		}
	}
}
