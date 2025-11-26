package main

import (
	"testing"
)

func TestFullName(t *testing.T) {
	tests := []struct {
		name string
		args User
		want string
	}{
		{
			name: "two empty string",
			args: User{"", ""},
			want: " ",
		},
		{
			name: "first empty",
			args: User{"", "Popov"},
			want: " Popov",
		},
		{
			name: "second empty",
			args: User{"Misha", ""},
			want: "Misha ",
		},
		{
			name: "non empty",
			args: User{"Misha", "Popov"},
			want: "Misha Popov",
		},
	}
	for _, test := range tests {
		if res := test.args.FullName(); res != test.want {
			t.Errorf("Error equal string want: %s, result: %s", test.want, res)
		} else {
			t.Logf("User{\"%s\", \"%s\"} = \"%s\"", test.args.FirstName, test.args.LastName, res)
		}
	}
}
