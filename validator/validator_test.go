package validator_test

import (
	"slices"
	"testing"

	"github.com/Jamlie/vodka/validator"
)

func TestValidateEmail(t *testing.T) {
	inputs := []string{"foo@bar.com", "foo@bar", "foo"}
	expected := []bool{true, true, false}
	actual := make([]bool, len(expected))

	for i, input := range inputs {
		actual[i] = validator.Email(input)
	}

	if !slices.Equal(expected, actual) {
		t.Fatalf("%v is not %v", expected, actual)
	}
}

func TestValidatePassword(t *testing.T) {
	inputs := []string{"qwerty", "foobar123", "fooBar123", "fooBar_123"}
	expected := [][]bool{
		{false, false, false, false},
		{true, true, false, false},
		{true, true, true, false},
		{true, true, true, true},
	}

	actual := make([][]bool, len(expected))

	for i, password := range inputs {
		actual[i] = append(actual[i], validator.Password(password, validator.PasswordOptions{
			Min:     8,
			Numbers: true,
		}))
		actual[i] = append(actual[i], validator.Password(password, validator.PasswordOptions{
			Min:           8,
			Numbers:       true,
			SmallAlphabet: true,
		}))
		actual[i] = append(actual[i], validator.Password(password, validator.PasswordOptions{
			Min:             8,
			Numbers:         true,
			SmallAlphabet:   true,
			CapitalAlphabet: true,
		}))
		actual[i] = append(actual[i], validator.Password(password, validator.PasswordOptions{
			Min:             8,
			Numbers:         true,
			SmallAlphabet:   true,
			CapitalAlphabet: true,
			Symbols:         true,
		}))
	}

	for i, result := range expected {
		if !slices.Equal(result, actual[i]) {
			t.Fatalf("%v is not %v", expected, actual)
		}
	}
}
