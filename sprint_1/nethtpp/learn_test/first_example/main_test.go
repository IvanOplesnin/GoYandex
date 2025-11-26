package main

import (
	"testing"
)

func TestAbs(t *testing.T) {
	tests := []struct {
		name   string
		number float64
		want   float64
	}{
		{
			name:   "test 0",
			number: 0,
			want:   0,
		},
		{
			name:   "test negative",
			number: -20.7,
			want:   20.7,
		},
		{
			name:   "test positive",
			number: 20.7,
			want:   20.7,
		},
		{
			name: "long float",
			number: -2.000001,
			want: 2.000001,
		},
		{
			name: "integer",
			number: -3,
			want: 3,
		},
	}

	for _, test := range tests {
		if result := Abs(test.number); result != test.want {
			t.Errorf("Error Abs %s", test.name)
		} else {
			t.Logf("%s: Abs(%f) = %f",test.name, test.number, test.want)
		}
	}
}
