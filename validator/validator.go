package validator

import (
	"net/mail"
	"regexp"
	"strings"
)

type PasswordOptions struct {
	Min             uint8
	Numbers         bool
	CapitalAlphabet bool
	SmallAlphabet   bool
	Symbols         bool
	Space           bool

	CustomSymbols string
}

func DefaultPassword() PasswordOptions {
	return PasswordOptions{
		Min:             8,
		Numbers:         true,
		CapitalAlphabet: true,
		SmallAlphabet:   true,
	}
}

var (
	numbers         = "1234567890"
	capitalAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	smallAlphabet   = "abcdefghijklmnopqrstuvwxyz"
	symbols         = "!@#$%^&*()-_=+,.?/:;{}[]~"
	space           = " "
)

func Email(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func Password(password string, options PasswordOptions) bool {
	if uint8(len(password)) < options.Min {
		return false
	}

	if options.Numbers && !strings.ContainsAny(password, numbers) {
		return false
	}

	if options.SmallAlphabet && !strings.ContainsAny(password, smallAlphabet) {
		return false
	}

	if options.CapitalAlphabet && !strings.ContainsAny(password, capitalAlphabet) {
		return false
	}

	if options.Symbols {
		if options.CustomSymbols != "" && !strings.ContainsAny(password, options.CustomSymbols) {
			return false
		}

		if !strings.ContainsAny(password, symbols) {
			return false
		}
	}

	if options.Space && !strings.ContainsAny(password, space) {
		return false
	}

	return true
}

func PhoneNumber(phoneNumber string) bool {
	matched, err := regexp.MatchString(
		`^[\+]?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}[-\s\.]?[0-9]{4,6}$`,
		phoneNumber,
	)
	if err != nil || !matched {
		return false
	}

	return true
}
