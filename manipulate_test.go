package strutil

import "testing"

func TestSquish(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"  hello   world  ", "hello world"},
		{"line1\n\tline2", "line1 line2"},
		{"", ""},
		{"   ", ""},
		{"already clean", "already clean"},
		{"\tleading tab", "leading tab"},
		{"trailing newline\n", "trailing newline"},
		{"multi\n\n\nlines", "multi lines"},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := Squish(tt.in); got != tt.want {
				t.Errorf("Squish(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestFinish(t *testing.T) {
	tests := []struct {
		s, suffix, want string
	}{
		{"path/to/dir", "/", "path/to/dir/"},
		{"path/to/dir/", "/", "path/to/dir/"},
		{"", "/", "/"},
		{"hello", "", "hello"},
		{"file.txt", ".txt", "file.txt"},
		{"file", ".txt", "file.txt"},
	}
	for _, tt := range tests {
		t.Run(tt.s+"+"+tt.suffix, func(t *testing.T) {
			if got := Finish(tt.s, tt.suffix); got != tt.want {
				t.Errorf("Finish(%q,%q) = %q, want %q", tt.s, tt.suffix, got, tt.want)
			}
		})
	}
}

func TestStart(t *testing.T) {
	tests := []struct {
		s, prefix, want string
	}{
		{"path/to/file", "/", "/path/to/file"},
		{"/path/to/file", "/", "/path/to/file"},
		{"", "/", "/"},
		{"hello", "", "hello"},
		{"://example", "http", "http://example"},
	}
	for _, tt := range tests {
		t.Run(tt.prefix+"+"+tt.s, func(t *testing.T) {
			if got := Start(tt.s, tt.prefix); got != tt.want {
				t.Errorf("Start(%q,%q) = %q, want %q", tt.s, tt.prefix, got, tt.want)
			}
		})
	}
}

func TestAfter(t *testing.T) {
	tests := []struct {
		name            string
		s, search, want string
	}{
		{"single space", "hello world", " ", "world"},
		{"first dot", "one.two.three", ".", "two.three"},
		{"not found", "hello", "z", "hello"},
		{"empty search", "hello", "", "hello"},
		{"empty string", "", "x", ""},
		{"search at start", "foo/bar", "foo", "/bar"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := After(tt.s, tt.search); got != tt.want {
				t.Errorf("After(%q,%q) = %q, want %q", tt.s, tt.search, got, tt.want)
			}
		})
	}
}

func TestBefore(t *testing.T) {
	tests := []struct {
		name            string
		s, search, want string
	}{
		{"single space", "hello world", " ", "hello"},
		{"first dot", "one.two.three", ".", "one"},
		{"not found", "hello", "z", "hello"},
		{"empty search", "hello", "", "hello"},
		{"empty string", "", "x", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Before(tt.s, tt.search); got != tt.want {
				t.Errorf("Before(%q,%q) = %q, want %q", tt.s, tt.search, got, tt.want)
			}
		})
	}
}

func TestBetween(t *testing.T) {
	tests := []struct {
		name                string
		s, start, end, want string
	}{
		{"brackets", "[hello]", "[", "]", "hello"},
		{"same delim", "a-b-c", "-", "-", "b"},
		{"missing", "no markers", "[", "]", ""},
		{"empty", "", "[", "]", ""},
		{"end before start", "]hello[", "[", "]", ""},
		{"only start", "[hello", "[", "]", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Between(tt.s, tt.start, tt.end); got != tt.want {
				t.Errorf("Between(%q,%q,%q) = %q, want %q", tt.s, tt.start, tt.end, got, tt.want)
			}
		})
	}
}
