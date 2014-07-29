package urlcheck

import (
	"code.google.com/p/go.net/publicsuffix"
	"log"
	"net/http"
	"net/http/cookiejar"
)

type Tester interface {
	Test() error
}

var Timeout uint = 5
var Verbose = false
var Version = "urlcheck/2.0"

var client *http.Client

// Set up a HTTP client with a cookie jar
func init() {
	options := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}

	jar, err := cookiejar.New(&options)
	if err != nil {
		log.Fatal(err)
	}

	client = &http.Client{
		Jar: jar,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			req.Header.Add("User-Agent", Version)
			return nil
		},
	}
}
