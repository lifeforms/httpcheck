package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/lifeforms/urlcheck/urlcheck"
	"io/ioutil"
	"os"
)

func parseArgs() (manifestfile string, server string, timeout uint, verbose bool) {
	flagi := flag.String("i", "manifest.yaml", "Input file with YAML manifest")
	flags := flag.String("s", "", "Only check this server, default check all servers in manifest")
	flagt := flag.Uint("t", 5, "Timeout for each HTTP request in seconds, 0 for no timeout")
	flagv := flag.Bool("v", false, "Verbose, prints the result of each test")
	flag.Parse()

	manifestfile = *flagi
	server = *flags
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

func filterManifest(in urlcheck.Manifest, server string) (out urlcheck.Manifest, err error) {
	if server == "" {
		return in, nil
	}

	for _, s := range in {
		if s.Name == server {
			out = append(out, s)
		}
	}
	if len(out) == 0 {
		return nil, errors.New("Server not in manifest: " + server)
	}
	return
}

func main() {
	manifestfile, server, timeout, verbose := parseArgs()

	urlcheck.RequestTimeout = timeout
	urlcheck.Verbose = verbose
	manifest, err := readManifest(manifestfile)

	if err == nil {
		manifest, err = filterManifest(manifest, server)
	}

	if err == nil {
		err = manifest.Test()
	}

	if err != nil {
		fmt.Println("Failures:")
		fmt.Println(err)
		os.Exit(1)
	} else {
		if verbose {
			fmt.Println("OK")
		}
	}
}
