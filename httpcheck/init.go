package httpcheck

import (
	"golang.org/x/net/publicsuffix"
	"log"
	"net/http"
	"net/http/cookiejar"
)

type Tester interface {
	Test() error
}

var RequestTimeout uint = 5  // Timeout for each HTTP request (seconds)
var ServerTimeout uint = 120 // Timeout for a server's tests (seconds)
var Verbose = false          // If true, every request is printed to standard output

var client http.Client
var version = "httpcheck/2.1"

// Set up a HTTP client with a cookie jar
func init() {
	options := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}

	jar, err := cookiejar.New(&options)
	if err != nil {
		log.Fatal(err)
	}

	client = http.Client{
		Jar: jar,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			req.Header.Add("User-Agent", version)
			return nil
		},
	}
}
