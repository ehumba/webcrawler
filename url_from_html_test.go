package main

import (
	"reflect"
	"testing"
)

func TestGetURLFromHTML(t *testing.T) {
	cases := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{name: "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "test #2",
			inputURL: "https://developer.mozilla.org",
			inputBody: `
<html>
	<body>
		<a href="/en-US/docs/Web">
			<span>mozilla.org</span>
		</a>
		<a href="https://other.com/path/one">
			<span>mozilla.org</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://developer.mozilla.org/en-US/docs/Web", "https://other.com/path/one"},
		},
	}

	for i, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if reflect.DeepEqual(actual, tc.expected) == false {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
