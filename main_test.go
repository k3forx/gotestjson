package main

import "testing"

func TestGet(t *testing.T) {
	cases := map[string]struct {
		input string
		want  string
	}{
		"ok": {
			input: "case1",
			want:  "case1",
		},
		"ng": {
			input: "case2",
			want:  "case2dfdfd",
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			got := Get(tt.input)
			if got != tt.want {
				t.Attr("input", tt.input)
				t.Errorf("expected '%s', but got '%s'", tt.want, got)
			}
		})
	}
}
