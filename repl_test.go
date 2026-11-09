package main

import (
	"testing"
	"github.com/google/go-cmp/cmp"
)


func TestCleanInput(t *testing.T){
	cases := map[string]struct {
		     input     string
		expected []    string
	}{
		"Doublespaced" : {input: "  hello  world  ", expected: []string{"hello", "world"}},
		"All Cap" : {input: " HELLO WORLD ", expected: []string{"hello", "world"}},
		"empty string": {input: "", expected: []string{}},
		"one word": {input: " hello ", expected: []string{"hello"}},
		"whitespace": {input: "   ", expected: []string{}},
	}

for name, c := range cases {
	t.Run(name, func(t *testing.T) {
	actual := cleanInput(c.input)
	diff := cmp.Diff(c.expected, actual)
	if diff != "" {
		t.Fatalf("%s", diff)
		}	
	 })
   }
}
