package main

import (
	"testing"
)

func TestValid(t *testing.T) {
	tables := []struct {
		input  string
		result bool
	}{
		{"hijklmmn", false},
		{"abbceffg", false},
		{"abbcegjk", false},
		{"abcdffaa", true},
		{"ghjaabcc", true},
		{"abcdeggg", false},
	}

	for _, table := range tables {
		p := CreateNewPassword(table.input)
		result := p.Valid()
		if result != table.result {
			t.Errorf("Input: %s, expected: %t, got: %t", table.input, table.result, result)
		}
	}
}

func TestNextPassword(t *testing.T) {
	tables := []struct {
		input  string
		result string
	}{
		{"abcdefgh", "abcdffaa"},
		{"ghijklmn", "ghjaabcc"},
	}

	for _, table := range tables {
		result := nextPassword(table.input)
		if result != table.result {
			t.Errorf("Input: %s, expected: %s, got: %s", table.input, table.result, result)
		}
	}
}

func TestIncrement(t *testing.T) {
	tables := []struct {
		input  string
		result string
	}{
		{"xx", "xy"},
		{"xy", "xz"},
		{"xz", "ya"},
		{"xzz", "yaa"},
		{"zzz", "aaa"},
	}

	for _, table := range tables {
		p := CreateNewPassword(table.input)
		p.Increment()
		result := p.String()
		if result != table.result {
			t.Errorf("Input: %s, expected: %s, got: %s", table.input, table.result, result)
		}
	}
}
