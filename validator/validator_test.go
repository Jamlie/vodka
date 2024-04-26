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

	if !slices.EqualFunc(expected, actual, func(exp, act []bool) bool {
		return slices.Equal(exp, act)
	}) {
		t.Fatalf("%v is not %v", expected, actual)
	}
}

func TestPhoneNumber(t *testing.T) {
	inputs := []string{"12345678", "+35313441111", "123-456-7890", "1-2-3", "9874575621"}
	expected := []bool{false, true, true, false, true}
	actual := make([]bool, len(expected))

	for i, input := range inputs {
		actual[i] = validator.PhoneNumber(input)
	}

	if !slices.Equal(expected, actual) {
		t.Fatalf("%v is not %v", expected, actual)
	}
}
