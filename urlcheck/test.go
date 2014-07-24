package urlcheck

import "errors"
import "io/ioutil"
import "net/http"
import "regexp"
import "strconv"
import "strings"

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
	resp, err := t.DoRequest()
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err == nil {
		err = t.CheckCode(resp)
	}
	if err == nil {
		err = t.CheckContent(resp)
	}
	return
}

func (t Test) DoRequest() (resp *http.Response, err error) {
	req, err := http.NewRequest(t.MethodName(), t.Url, strings.NewReader(t.Data))
	if err != nil {
		return nil, err
	}

	for k, v := range t.Headers {
		req.Header.Add(k, v)
	}

	resp, err = client.Do(req)
	return resp, err
}

func (t Test) CheckCode(resp *http.Response) error {
	code := t.Code
	if code == 0 {
		code = 200
	}
	if resp.StatusCode != code {
		return errors.New("Expected status code " + strconv.Itoa(code) + ", received " + strconv.Itoa(resp.StatusCode))
	}
	return nil
}

func (t Test) CheckContent(resp *http.Response) error {
	if t.Content == "" {
		return nil
	}

	rcvdbytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	rcvdcontent := string(rcvdbytes)

	match, err := regexp.MatchString(t.Content, rcvdcontent)
	if err != nil {
		return err
	}

	if !match {
		return errors.New("Expected content '" + t.Content + "', not found in response (" + strconv.Itoa(len(rcvdbytes)) + " bytes)")
	}

	return nil
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
