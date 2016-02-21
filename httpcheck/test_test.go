package httpcheck

import "testing"

func TestValidateBad(t *testing.T) {
	bad := []Test{
		Test{},
		Test{Url: ""},
		Test{Url: "/"},
		Test{Url: "/foo"},
		Test{Url: "http://example/", Method: "GET", Data: "foo"},
		Test{Url: "http://example/", Method: "GET", Type: "text/xml"},
		Test{Url: "http://example/", Method: "HEAD", Data: "foo"},
		Test{Url: "http://example/", Method: "HEAD", Type: "text/xml"},
		Test{Url: "http://example/", Method: "OPTIONS", Data: "foo"},
		Test{Url: "http://example/", Method: "OPTIONS", Type: "text/xml"},
	}
	for _, test := range bad {
		if test.Validate() == nil {
			t.Error("Test.Validate() should return error:", test)
		}
	}
}

func TestValidateGood(t *testing.T) {
	good := []Test{
		Test{Url: "http://example/"},
		Test{Url: "https://example/"},
		Test{Url: "httP://example/"},
		Test{Url: "http://example/", Method: ""},
		Test{Url: "http://example/", Method: "GET"},
		Test{Url: "http://example/", Method: "GET", Data: ""},
		Test{Url: "http://example/", Method: "GET", Headers: map[string]string{"Cookie": "foo=bar"}},
		Test{Url: "http://example/", Method: "POST"},
		Test{Url: "http://example/", Method: "POST", Data: ""},
		Test{Url: "http://example/", Method: "POST", Data: "foo=bar"},
		Test{Url: "http://example/", Method: "POST", Data: "<foo></foo>", Type: "text/xml"},
	}
	for _, test := range good {
		if test.Validate() != nil {
			t.Error("Test.Validate() should return nil:", test)
		}
	}
}

func TestMethodName(t *testing.T) {
	cases := []struct {
		input  Test
		expect string
	}{
		{Test{}, "GET"},
		{Test{Method: ""}, "GET"},
		{Test{Method: "get"}, "get"},
		{Test{Method: "GET"}, "GET"},
		{Test{Method: "POST"}, "POST"},
		{Test{Method: "COOK"}, "COOK"},
	}
	for _, c := range cases {
		result := c.input.MethodName()
		if result != c.expect {
			t.Error("Test.MethodName ", c.input.Method, "should return", c.expect, "but got", result)
		}
	}
}

func TestContentType(t *testing.T) {
	cases := []struct {
		input  Test
		expect string
	}{
		{Test{}, ""},
		{Test{Method: "GET"}, ""},
		{Test{Method: "POST"}, "application/x-www-form-urlencoded"},
		{Test{Method: "POST", Type: "text/xml"}, "text/xml"},
	}
	for _, c := range cases {
		result := c.input.ContentType()
		if result != c.expect {
			t.Error("Test.ContentType() ", c.input.Type, "should return", c.expect, "but got", result)
		}
	}
}
