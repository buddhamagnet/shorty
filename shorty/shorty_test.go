package shorty

import "testing"

func TestIsValidURL(t *testing.T) {
	for _, tt := range []struct {
		url      string
		expected bool
	}{
		{"https://www.google.com", true},
		{"www.google.com", false},
		{"google.com", false},
		{"bogus", false},
		{"http://", false},
	} {
		actual := IsValidURL(tt.url)
		if actual != tt.expected {
			t.Errorf("Fib(%s): expected %t, actual %t", tt.url, tt.expected, actual)
		}
	}
}
