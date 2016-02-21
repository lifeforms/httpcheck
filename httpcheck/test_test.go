package httpcheck

import "testing"

func TestValidateBad(t *testing.T) {
	bad := []Test{
		Test{},
		Test{Url: ""},
		Test{Url: "/"},
		Test{Url: "/foo"},
		Test{Url: "http://example/", Method: "GET", Data: "foo"},
	}
	for _, test := range bad {
		if test.Validate() == nil {
			t.Error("Test.Validate() should return error:", test)
		}
	}
}

func TestValidateGood(t *testing.T) {
	bad := []Test{
		Test{Url: "http://example/"},
		Test{Url: "http://example/", Method: ""},
		Test{Url: "http://example/", Method: "GET"},
		Test{Url: "http://example/", Method: "GET", Data: ""},
		Test{Url: "http://example/", Method: "POST"},
		Test{Url: "http://example/", Method: "POST", Data: "foo=bar"},
		Test{Url: "http://example/", Method: "POST", Data: "<foo></foo>", Type: "text/xml"},
	}
	for _, test := range bad {
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
