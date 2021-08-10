package main

import "testing"

func TestGenPassword(t *testing.T) {
	chars := map[string]string{
		"lower":   "abcdefghijklmnopqrstuvwxyz",
		"upper":   "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"numbers": "0123456789",
		"extra":   "!@#$%^&*?",
	}

	tc := []struct {
		chars    map[string]string
		config   *Config
		expected string
	}{
		{chars, &Config{[]string{"lower", "upper", "numbers", "extra"}, 1, 8}, "V#GAc19w"},
		{chars, &Config{[]string{"lower", "upper", "numbers", "extra"}, 1, 0}, ""},
		{chars, &Config{[]string{"lower"}, 1, 8}, "vbgacrjw"},
		{chars, &Config{[]string{}, 1, 8}, ""},
		{map[string]string{}, &Config{[]string{"lower", "upper", "numbers", "extra"}, 1, 8}, ""},
	}

	for _, c := range tc {
		pass, _ := GenPassword(c.chars, c.config)

		if pass != c.expected {
			t.Errorf("expected a %s; got %s", c.expected, pass)
		}
	}

}
