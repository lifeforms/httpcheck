package urlcheck

import "fmt"

type Method int

const (
	GET Method = iota
	POST
)

type Test struct {
	Url           string
	Content       string
	Code          int
	Method        Method
	Data          string
	Headers       map[string]string
	SkipSSLVerify bool
}

func (t Test) Test() error {
	fmt.Println("testing:", t.Url)
	return nil
}
