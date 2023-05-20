package domain

import (
	"reflect"
	"testing"
)

var (
	defaultEmail = Email{}
)

func TestNewEmail(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  Email
	}{
		{name: "CanCreateEmailType", value: "test@example.com", want: defaultEmail},
		// TODO: Add Tests
	}

	for _, tc := range tests {
		Email := NewEmail(tc.value)
		if !reflect.DeepEqual(reflect.TypeOf(Email), reflect.TypeOf(tc.want)) {
			t.Fatalf("Type is NOT CORRECT")
		}
	}
}

func TestEmailToValue(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{name: "no diff", value: "test@example.com", want: "test@example.com"},
		// TODO: Add Tests
	}

	for _, tc := range tests {
		Email := NewEmail(tc.value)
		got := Email.ToValue()
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("case: %v, expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}
