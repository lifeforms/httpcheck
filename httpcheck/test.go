package httpcheck

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Method int

// A single HTTP action, with the expected results.
type Test struct {
	Url           string            // A fully specified URL including the protocol
	Content       string            ",omitempty" // Expected content as a regexp, e.g. "Hello World"
	Code          int               ",omitempty" // Expected HTTP response code
	Method        string            ",omitempty" // HTTP method, i.e. "GET" (default) or "POST"
	Type          string            ",omitempty" // Optional value for Content-Type header
	Data          string            ",omitempty" // Optional post data
	Headers       map[string]string ",omitempty" // Optional headers to add to the request
	SkipSSLVerify bool              ",omitempty" // If true, SSL server verification is skipped
}

// Test makes a HTTP call and checks the response
func (t Test) Test() (err error) {
	err = t.Validate()

	var code int
	var body string
	if err == nil {
		code, body, err = t.DoRequest()
	}
	if err == nil {
		err = t.CheckCode(code)
	}
	if err == nil {
		err = t.CheckContent(body)
	}

	// Log an error
	if Verbose {
		if err == nil {
			fmt.Println(t.String() + " OK")
		} else {
			fmt.Println(err)
		}
	}
	return
}

func (t Test) Validate() error {
	scheme := regexp.MustCompile("^https?:")
	if !scheme.MatchString(t.Url) {
		return errors.New("URL is not absolute, must specify a base URL (-u): " + t.Url)
	}
	return nil
}

// DoRequest uses the global http object to send a HTTP request
func (t Test) DoRequest() (code int, body string, err error) {
	req, err := http.NewRequest(t.MethodName(), t.Url, strings.NewReader(t.Data))
	if err != nil {
		return 0, "", err
	}

	req.Header.Add("User-Agent", version)

	if t.ContentType() != "" {
		req.Header.Add("Content-Type", t.ContentType())
	}

	for k, v := range t.Headers {
		req.Header.Add(k, v)
	}

	client.Timeout = time.Duration(int(RequestTimeout)) * time.Second
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

// CheckCode inspects the received HTTP response code
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

// CheckContent inspects the returned HTTP response body
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
	return errors.New(t.String() + " FAIL: " + message)
}

func (t Test) String() string {
	return strings.Title(strings.ToLower(t.MethodName())) + " " + t.Url
}

func (t Test) MethodName() string {
	if t.Method == "" {
		return "GET"
	}
	return t.Method
}

func (t Test) ContentType() string {
	if t.Type != "" {
		return t.Type
	}
	if t.Method == "POST" {
		return "application/x-www-form-urlencoded"
	}
	return ""
}
