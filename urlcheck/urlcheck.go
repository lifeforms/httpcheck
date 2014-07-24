package urlcheck

import "net/http"

type Testable interface {
	Test() error
}

type Manifest []Server

var client = &http.Client{}
