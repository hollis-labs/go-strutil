package strutil

import (
	"reflect"
	"testing"
)

func TestSplitWords(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want []string
	}{
		{"empty", "", nil},
		{"single lower", "hello", []string{"hello"}},
		{"camelCase", "helloWorld", []string{"hello", "World"}},
		{"PascalCase", "HelloWorld", []string{"Hello", "World"}},
		{"snake_case", "hello_world", []string{"hello", "world"}},
		{"kebab-case", "hello-world", []string{"hello", "world"}},
		{"space separated", "hello world", []string{"hello", "world"}},
		{"multiple separators", "hello__--  world", []string{"hello", "world"}},
		{"acronym then word", "XMLParser", []string{"XML", "Parser"}},
		{"word then acronym", "parseXML", []string{"parse", "XML"}},
		{"all caps run", "HELLO", []string{"HELLO"}},
		{"number boundary", "user2Name", []string{"user2", "Name"}},
		{"leading separator", "_hello_world", []string{"hello", "world"}},
		{"trailing separator", "hello_world_", []string{"hello", "world"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := splitWords(tt.in)
			if len(got) == 0 && len(tt.want) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitWords(%q) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestSnakeCase(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"helloWorld", "hello_world"},
		{"HelloWorld", "hello_world"},
		{"hello-world", "hello_world"},
		{"hello world", "hello_world"},
		{"hello_world", "hello_world"},
		{"XMLParser", "xml_parser"},
		{"", ""},
		{"single", "single"},
		{"UPPER", "upper"},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := SnakeCase(tt.in); got != tt.want {
				t.Errorf("SnakeCase(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestKebabCase(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"helloWorld", "hello-world"},
		{"HelloWorld", "hello-world"},
		{"hello_world", "hello-world"},
		{"hello world", "hello-world"},
		{"hello-world", "hello-world"},
		{"XMLParser", "xml-parser"},
		{"", ""},
		{"café", "café"}, // preserves multi-byte (no transliteration)
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := KebabCase(tt.in); got != tt.want {
				t.Errorf("KebabCase(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestCamelCase(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"hello_world", "helloWorld"},
		{"hello-world", "helloWorld"},
		{"Hello World", "helloWorld"},
		{"HelloWorld", "helloWorld"},
		{"helloWorld", "helloWorld"},
		{"XMLParser", "xmlParser"},
		{"", ""},
		{"single", "single"},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := CamelCase(tt.in); got != tt.want {
				t.Errorf("CamelCase(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestStudlyCase(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"hello_world", "HelloWorld"},
		{"hello-world", "HelloWorld"},
		{"hello world", "HelloWorld"},
		{"helloWorld", "HelloWorld"},
		{"HelloWorld", "HelloWorld"},
		{"XMLParser", "XmlParser"},
		{"", ""},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := StudlyCase(tt.in); got != tt.want {
				t.Errorf("StudlyCase(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestTitle(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"hello world", "Hello World"},
		{"HELLO WORLD", "Hello World"},
		{"hello_world", "Hello_world"}, // underscore not a separator
		{"", ""},
		{"a b c", "A B C"},
		{"café du monde", "Café Du Monde"},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := Title(tt.in); got != tt.want {
				t.Errorf("Title(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestUcFirst(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"hello", "Hello"},
		{"Hello", "Hello"},
		{"", ""},
		{"h", "H"},
		{"éclair", "Éclair"},
		{"HELLO", "HELLO"},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := UcFirst(tt.in); got != tt.want {
				t.Errorf("UcFirst(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestLcFirst(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"Hello", "hello"},
		{"hello", "hello"},
		{"", ""},
		{"H", "h"},
		{"Éclair", "éclair"},
		{"HELLO", "hELLO"},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := LcFirst(tt.in); got != tt.want {
				t.Errorf("LcFirst(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
