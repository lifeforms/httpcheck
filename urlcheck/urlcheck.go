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

type Manifest []Server

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
	client = &http.Client{Jar: jar}
}
