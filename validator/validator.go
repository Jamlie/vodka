package validator

import (
	"net/mail"
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
	if err != nil {
		return false
	}

	return true
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
