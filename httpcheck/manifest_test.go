package httpcheck

import "testing"
import "reflect"

func TestSetBaseURL(t *testing.T) {
	cases := []struct {
		url    string
		base   string
		expect string
	}{
		{"", "http://base", "http://base"},
		{"/", "http://base", "http://base/"},
		{"", "http://base/", "http://base/"},
		{"/", "http://base/", "http://base/"},
		{"foo", "http://base/", "http://base/foo"},
		{"/foo", "http://base/", "http://base/foo"},
		{"http://example.com", "http://base/", "http://example.com"},
		{"https://example.com", "http://base/", "https://example.com"},
		{"HTTP://example.com", "http://base/", "HTTP://example.com"},
		{"HTTPS://example.com", "http://base/", "HTTPS://example.com"},
	}

	for _, c := range cases {
		manifest := Manifest{
			Server{
				Scenarios: []Scenario{
					Scenario{
						Tests: Tests{Test{Url: c.url}}},
				},
			},
		}
		manifest.SetBaseURL(c.base)
		result := manifest[0].Scenarios[0].Tests[0].Url
		if result != c.expect {
			t.Error("SetBaseURL() on", c.url, ",", c.base, "should return", c.expect, "but got", result)
		}
	}
}

func TestFromYAML(t *testing.T) {
	input := `- server: foo
  scenarios:
  - name: bar
    test: [{url: 'http://example.com/'}]
  - test: [{url: 'https://example.com/', method: 'POST', data: 'foo=bar'}]
  - test: [{url: 'https://example.com/', method: 'POST', data: '<foo></foo>', type: 'text/xml'}]
  - test: [{url: '/', method: 'OPTIONS'}]
  - test: [{url: '/', headers: {Cookie: 'foo=bar'}}]
  - test: [{url: '/', headers: {'Cookie': 'foo=bar'}}]
`
	expect := Manifest{
		Server{
			Name: "foo",
			Scenarios: []Scenario{
				Scenario{
					Name:  "bar",
					Tests: Tests{Test{Url: "http://example.com/"}}},
				Scenario{
					Tests: Tests{Test{Url: "https://example.com/", Method: "POST", Data: "foo=bar"}}},
				Scenario{
					Tests: Tests{Test{Url: "https://example.com/", Method: "POST", Data: "<foo></foo>", Type: "text/xml"}}},
				Scenario{
					Tests: Tests{Test{Url: "/", Method: "OPTIONS"}}},
				Scenario{
					Tests: Tests{Test{Url: "/", Headers: map[string]string{"Cookie": "foo=bar"}}}},
				Scenario{
					Tests: Tests{Test{Url: "/", Headers: map[string]string{"Cookie": "foo=bar"}}}},
			},
		},
	}

	manifest, err := FromYAML([]byte(input))
	if err != nil {
		t.Error("FromYAML() failed:", err)
	}
	if !reflect.DeepEqual(manifest, expect) {
		t.Error("FromYAML() returned unexpected result")
	}
}
