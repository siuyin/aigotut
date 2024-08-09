package main

import "testing"

func TestJsonify(t *testing.T) {
	m := map[string]any{
		"a":      "apple",
		"i":      42,
		"o":      map[string]any{"cat": 24.0, "dog_name": "woofer"},
		"fruits": []interface{}{"bananas", "pineapple", "oranges"},
		"f2":     map[string]any{"": []interface{}{"bananas", "pineapple", "oranges"}},
	}
	s := jsonify(m)
	if s != "gerbau" {
		t.Errorf("unexpected: %q", s)
	}
}
