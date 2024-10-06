package constants

import (
	"testing"
)

// TestTagValues checks that the Tag constants have the expected string values.
func TestTagValues(t *testing.T) {
	tests := []struct {
		name     string
		tag      Tag
		expected string
	}{
		{name: "EventTag", tag: Event, expected: "event"},
		{name: "InternalTag", tag: Internal, expected: "internal"},
		{name: "MemoryTag", tag: Memory, expected: "memory"},
		{name: "ChildTag", tag: Child, expected: "child"},
		{name: "AuthnTag", tag: Authn, expected: "authn"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if string(tt.tag) != tt.expected {
				t.Errorf("Expected %s but got %s", tt.expected, tt.tag)
			}
		})
	}
}

// TestTagType checks that the constants are of type Tag.
func TestTagType(t *testing.T) {
	var tag Tag

	tests := []struct {
		name string
		tag  Tag
	}{
		{name: "EventTag", tag: Event},
		{name: "InternalTag", tag: Internal},
		{name: "MemoryTag", tag: Memory},
		{name: "ChildTag", tag: Child},
		{name: "AuthnTag", tag: Authn},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tag = tt.tag
			if _, ok := interface{}(tag).(Tag); !ok {
				t.Errorf("Expected type Tag but got different type")
			}
		})
	}
}

