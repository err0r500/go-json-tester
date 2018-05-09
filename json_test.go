package json_tester_test

import (
	"testing"
	"github.com/err0r500/go-json-tester"
)

func Test(t *testing.T) {
	b := []byte(`{"x": ["y","42"]}`)
	c := []byte(`{"x": ["y","42."]}`)
	d := []byte(`{"x": {}}`)
	e := []byte(`{}`)

	tests := []struct {
		j1       []byte
		j2       []byte
		expected bool
	}{
		{j1: b, j2: b, expected: true},
		{j1: c, j2: c, expected: true},
		{j1: b, j2: c, expected: false},
		{j1: d, j2: d, expected: true},
		{j1: d, j2: e, expected: false},
	}
	for k, tt := range tests {
		err := json_tester.JSONEqual(tt.j1, tt.j2)

		if tt.expected && err != nil {
			t.Errorf("test #%d should return no error : %s", k, err)
		} else if !tt.expected && err == nil {
			t.Errorf("test #%d should return an error", k)
		}
	}
}
