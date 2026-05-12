package collector

import "testing"

const testBar = "bar"

var containsTestCases = []struct {
	scenarioName string
	list         []string
	element      string
	expected     bool
}{
	{"contains element", []string{testBar, "foo"}, testBar, true},
	{"does not contain element", []string{"foo", testBar}, "baz", false},
}

func TestContainsFunction(t *testing.T) {
	for _, tt := range containsTestCases {
		t.Run(tt.scenarioName, func(t *testing.T) {
			if contains(tt.list, tt.element) != tt.expected {
				t.Errorf("Unexpected result")
			}
		})
	}
}
