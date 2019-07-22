package logic

import (
	"fmt"
	"testing"
)

func TestAddWords(t *testing.T) {

	type In struct {
		firstName string
		lastName  string
	}

	testCases := []struct {
		name string
		in   In
		out  string
	}{
		{
			name: "Add first and last names",
			in: In{
				firstName: "Vardenis",
				lastName:  "Pavardenis",
			},
			out: "My name is Vardenis Pavardenis",
		},
		{
			name: "Missing first name",
			in: In{
				firstName: "",
				lastName:  "Pavardenis",
			},
			out: "My name is  Pavardenis",
		},
		{
			name: "Missing last name",
			in: In{
				firstName: "Vardenis",
				lastName:  "",
			},
			out: "My name is Vardenis ",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			expRes := AddWords(tc.in.firstName, tc.in.lastName)
			fmt.Printf("expRes " + expRes)
			if expRes != tc.out {
				t.Errorf("got %q, want %q", expRes, tc.out)
			}
		})
	}
}
