package github

import (
	"testing"
)

func TestParseLinks(t *testing.T) {
	var testcases = []struct {
		data     string
		expected map[string]string
	}{
		{
			``,
			map[string]string{},
		},
		{
			`<u>; rel="r"`,
			map[string]string{"r": "u"},
		},
		{
			`<u>; rel="r", <u2>; rel="r2"`,
			map[string]string{"r": "u", "r2": "u2"},
		},
		{
			`<u>;rel="r"`,
			map[string]string{"r": "u"},
		},
		{
			`<u>; rel="r"; other="o"`,
			map[string]string{"r": "u"},
		},
		{
			`<u>; rel="r",,`,
			map[string]string{"r": "u"},
		},
	}
	for _, tc := range testcases {
		actual := parseLinks(tc.data)
		if len(actual) != len(tc.expected) {
			t.Errorf("lengths mismatch: %v and %v", actual, tc.expected)
		}
		for k, v := range tc.expected {
			if av, ok := actual[k]; !ok {
				t.Errorf("link not found: %s", k)
			} else if av != v {
				t.Errorf("link %s was %s but should have been %s", k, av, v)
			}
		}
	}
}
