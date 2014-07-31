package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/lifeforms/httpcheck/httpcheck"
	"io/ioutil"
	"os"
)

func parseArgs() (manifestfile string, server string, rt uint, st uint, verbose bool) {
	flagi := flag.String("i", "manifest.yaml", "Input file with YAML manifest")
	flags := flag.String("s", "", "Only check this server, default check all servers in manifest")
	flagr := flag.Uint("r", 5, "Timeout for each HTTP request in seconds, 0 for no timeout")
	flagt := flag.Uint("t", 120, "Timeout for each server in seconds, 0 for no timeout")
	flagv := flag.Bool("v", false, "Verbose, prints the result of each test")
	flag.Parse()

	manifestfile = *flagi
	server = *flags
	rt = *flagr
	st = *flagt
	verbose = *flagv
	return
}

func readManifest(manifestfile string) (manifest httpcheck.Manifest, err error) {
	y, err := ioutil.ReadFile(manifestfile)
	if err == nil {
		manifest, err = httpcheck.FromYAML(y)
	}
	return
}

func filterManifest(in httpcheck.Manifest, server string) (out httpcheck.Manifest, err error) {
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
	manifestfile, server, rt, st, verbose := parseArgs()

	httpcheck.RequestTimeout = rt
	httpcheck.ServerTimeout = st
	httpcheck.Verbose = verbose
	manifest, err := readManifest(manifestfile)

	if err == nil {
		manifest, err = filterManifest(manifest, server)
	}

	if err == nil {
		err = manifest.Test()
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		if verbose {
			fmt.Println("OK")
		}
	}
}
