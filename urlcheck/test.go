package urlcheck

import "fmt"

type Method int

const (
	GET Method = iota
	POST
)

type Test struct {
	Url           string
	Method        Method
	Data          string
	Code          int
	Content       string
	SkipSSLVerify bool
}

func (t Test) Test() error {
	fmt.Println("testing:", t.Url)
	return nil
}
