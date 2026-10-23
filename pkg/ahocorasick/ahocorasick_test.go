package ahocorasick

import (
	"reflect"
	"sort"
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		patterns []string
		want     []string
	}{
		{"single", "I eat and drink", []string{"eat", "drink"}, []string{"drink", "eat"}},
		{"suffix pattern", "abc", []string{"abc", "bc"}, []string{"abc", "bc"}},
		{"nested suffixes", "xabcd", []string{"abcd", "bcd", "cd"}, []string{"abcd", "bcd", "cd"}},
		{"inherited output on non-word node", "abc", []string{"abcx", "bc"}, []string{"bc"}},
		{"overlap", "ushers", []string{"he", "she", "his", "hers"}, []string{"he", "hers", "she"}},
		{"no match", "hello", []string{"xyz"}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Search(tt.text, tt.patterns)
			sort.Strings(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search(%q, %v) = %v, want %v", tt.text, tt.patterns, got, tt.want)
			}
		})
	}
}
