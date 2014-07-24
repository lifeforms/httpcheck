package urlcheck

import (
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

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

func (t Test) Test() (err error) {
	code, body, err := t.DoRequest()
	if err != nil {
		return err
	}

	if err == nil {
		err = t.CheckCode(code)
	}
	if err == nil {
		err = t.CheckContent(body)
	}
	return
}

func (t Test) DoRequest() (code int, body string, err error) {
	req, err := http.NewRequest(t.MethodName(), t.Url, strings.NewReader(t.Data))
	if err != nil {
		return 0, "", err
	}

	for k, v := range t.Headers {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, "", err
	}

	rcvdbytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, "", err
	}
	body = string(rcvdbytes)
	code = resp.StatusCode

	defer resp.Body.Close()

	return code, body, nil
}

func (t Test) CheckCode(code int) error {
	expect := t.Code
	if expect == 0 {
		expect = 200
	}
	if code != expect {
		return t.NewError("Expected status code " + strconv.Itoa(expect) + ", received " + strconv.Itoa(code))
	}
	return nil
}

func (t Test) CheckContent(body string) error {
	if t.Content == "" {
		return nil
	}

	match, err := regexp.MatchString(t.Content, body)
	if err != nil {
		return err
	}

	if !match {
		return t.NewError("Expected content '" + t.Content + "', not found in response (" + strconv.Itoa(len(body)) + " bytes)")
	}

	return nil
}

func (t Test) NewError(message string) error {
	return errors.New(strings.Title(strings.ToLower(t.MethodName())) + " " + t.Url + ": " + message)
}

func (t Test) MethodName() string {
	switch t.Method {
	case GET:
		return "GET"
	case POST:
		return "POST"
	default:
		return ""
	}
}
