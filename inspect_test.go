package strutil

import "testing"

func TestContainsAll(t *testing.T) {
	tests := []struct {
		name string
		s    string
		subs []string
		want bool
	}{
		{"all present", "hello world", []string{"hello", "world"}, true},
		{"one missing", "hello world", []string{"hello", "mars"}, false},
		{"empty subs", "hello", nil, true},
		{"empty string with sub", "", []string{"hello"}, false},
		{"empty string empty subs", "", nil, true},
		{"single sub present", "hello", []string{"ell"}, true},
		{"empty sub element", "hello", []string{""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsAll(tt.s, tt.subs); got != tt.want {
				t.Errorf("ContainsAll(%q,%v) = %v, want %v", tt.s, tt.subs, got, tt.want)
			}
		})
	}
}

func TestContainsAny(t *testing.T) {
	tests := []struct {
		name string
		s    string
		subs []string
		want bool
	}{
		{"one present", "hello world", []string{"mars", "world"}, true},
		{"none present", "hello world", []string{"mars", "venus"}, false},
		{"empty subs", "hello", nil, false},
		{"empty string", "", []string{"x"}, false},
		{"all present", "hello", []string{"ell", "hel"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsAny(tt.s, tt.subs); got != tt.want {
				t.Errorf("ContainsAny(%q,%v) = %v, want %v", tt.s, tt.subs, got, tt.want)
			}
		})
	}
}
