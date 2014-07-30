httpcheck
=========

Simple HTTP website monitoring library and tool in Go. It retrieves a set of specified HTTP URLs, and compares the response codes/bodies with the specification. It checks servers in parallel and understands cookies.

## Installation

Install the tool: `go get github.com/lifeforms/httpcheck/cmd/httpcheck`

Just install the library for use in Go code: `go get github.com/lifeforms/httpcheck/httpcheck`

## Concepts

The `httpcheck` library and tool operate on a *manifest*. The manifest contains a list of servers and test scenarios. It can be described as a YAML file, or it can be instantiated from Go code.

Each *server* has one or more *scenarios*. A scenario describes a list of *test* requests, which are executed in order. For example, a scenario could consist of the following tests: (1) getting a website's home page, (2) doing a search query, (3) doing a query with an incorrect parameter, et cetera. If one of the tests in a scenario fails, the scenario is abandoned. It is valid for a scenario to contain just one test.

Scenarios are grouped under a server, because `httpcheck` will test different servers concurrently. This prevents a single server or application from being swamped with test traffic. The server name is only used for administrative purposes.

## Manifest

This is an example manifest in YAML format, found in the repository as `example.yaml`:

    # A server with various simple URLs to check.
    # If code is not supplied, status code 200 is expected.

    - server: lifeforms
      scenarios:
      - test: [{url: 'http://lifeforms.nl/', content: 'Lifeforms'}] # follows redirect to HTTPS
      - test: [{url: 'https://lifeforms.nl/', content: 'Lifeforms'}]
      - test: [{url: 'https://lifeforms.nl/gfx', code: 403}]
      - test: [{url: 'https://lifeforms.nl/nonexistent', code: 404}]

    # Servers with each one scenario consisting of a user doing a search.
    # If one step in a scenario fails, that scenario is abandoned.
    # Cookies are retained, so scenario steps can depend on earlier steps.

    - server: google
      scenarios:
      - name: search
        test: [{url: 'https://google.com/', content: '<title>Google</title>'},
               {url: 'https://google.com/search?q=test', content: 'test - Google Search'}]

    - server: bing
      scenarios:
      - name: search
        test: [{url: 'https://bing.com/', content: '<title>Bing</title>'},
               {url: 'https://bing.com/search?q=test', content: '<title>test - Bing</title>'}]

The `url` field specifies the HTTP URL to call. If `content` is supplied, we check the response body for the regular expression. If `code` is supplied, we expect that HTTP status code; if not, we expect code 200.

It's optional to supply a `method` of GET (default) or POST. If `data` contains a string, it is sent as POST data. In `headers`, HTTP header name/value pairs can be specified.

## Usage

To check a manifest: `httpcheck -i example.yaml` If all test scenarios succeed, no output will be printed. If there are failures, the failure will be printed, and the command will exit with an error status.

To print the result of each individual test, use the verbose flag: `httpcheck -i example.yaml -v`

There are default timeouts per request (maximum time taken for the response to arrive and be read) and per server (maximum time for all tests). These prevent downed infrastructure from keeping the check program waiting forever. See `httpcheck -h` to view advanced options.

## Library example

A simple use of the library from Go:

    package main

    import (
    	"fmt"
    	"github.com/lifeforms/httpcheck/httpcheck"
    )

    func main() {
    	// Optionally change settings:
    	httpcheck.Verbose = true
    	httpcheck.RequestTimeout = 5
    	httpcheck.ServerTimeout = 120

    	manifest := httpcheck.Manifest{
    		httpcheck.Server{
    			Name: "tau",
    			Scenarios: []httpcheck.Scenario{
    				httpcheck.Scenario{
    					Name:  "http",
    					Tests: httpcheck.Tests{httpcheck.Test{Url: "http://lifeforms.nl/", Content: "Lifeforms"}}},
    				httpcheck.Scenario{
    					Name:  "https",
    					Tests: httpcheck.Tests{httpcheck.Test{Url: "https://lifeforms.nl/", Content: "Lifeforms"}}},
    				httpcheck.Scenario{
    					Name:  "opendir",
    					Tests: httpcheck.Tests{httpcheck.Test{Url: "https://lifeforms.nl/gfx", Code: 403}}},
    				httpcheck.Scenario{
    					Name:  "nonexistent",
    					Tests: httpcheck.Tests{httpcheck.Test{Url: "https://lifeforms.nl/nonexistent", Code: 404}}},
    			},
    		},
    	}

    	fmt.Println("Manifest:", manifest)
    	fmt.Println("Test result (nil is OK):", manifest.Test())
    }
