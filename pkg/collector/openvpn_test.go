package collector

import "testing"

var containsTestCases = []struct {
	scenarioName string
	element      string
	list         []string
	expected     bool
}{
	{"contains element", "bar", []string{"bar", "foo"}, true},
	{"does not contain element", "baz", []string{"foo", "bar"}, false},
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
