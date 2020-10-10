package flags

import (
	"reflect"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	var tests = []struct {
		args []string
		conf Config
	}{
		{[]string{"--url", "http://www.google.com"},
			Config{Service: "shorten", LongURL: "http://www.google.com", ID: "", args: []string{}}},

		{[]string{"--service", "shorten", "--url", "http://www.google.com"},
			Config{Service: "shorten", LongURL: "http://www.google.com", ID: "", args: []string{}}},

		{[]string{"--service", "decode", "--id", "RICKROLLED"},
			Config{Service: "decode", LongURL: "", ID: "RICKROLLED", args: []string{}}},
	}

	for _, tt := range tests {
		t.Run(strings.Join(tt.args, " "), func(t *testing.T) {
			conf, output, err := Parse("shorty", tt.args)
			if err != nil {
				t.Errorf("err got %v, want nil", err)
			}
			if output != "" {
				t.Errorf("output got %q, want empty", output)
			}
			if !reflect.DeepEqual(*conf, tt.conf) {
				t.Errorf("conf got %+v, want %+v", *conf, tt.conf)
			}
		})
	}
}

func TestErrors(t *testing.T) {
	var tests = []struct {
		args    []string
		conf    Config
		message string
	}{
		{[]string{},
			Config{Service: "shorten", LongURL: "http://www.google.com", ID: "", args: []string{}},
			"usage: cli (--url|-u)=<url>"},

		{[]string{"--service", "shorten"},
			Config{Service: "shorten", LongURL: "http://www.google.com", ID: "", args: []string{}},
			"usage: cli (--url|-u)=<url>"},

		{[]string{"--service", "decode"},
			Config{Service: "decode", LongURL: "", ID: "RICKROLLED", args: []string{}},
			"usage: cli (--id)=<id>"},
	}

	for _, tt := range tests {
		t.Run(strings.Join(tt.args, " "), func(t *testing.T) {
			_, _, err := Parse("shorty", tt.args)
			if err == nil {
				t.Errorf("expected error, got nil")
			}
			if err.Error() != tt.message {
				t.Errorf("expected %s, got %s", tt.message, err.Error())

			}
		})
	}
}
