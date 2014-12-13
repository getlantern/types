package types

import (
	"fmt"
	"net/mail"
	"strings"
)

var (
	domainsMap = map[string]string{
		"googlemail.com": "gmail.com",
	}
)

// InvalidEmailError indicates that an email address was not valid.
type InvalidEmailError struct {
	email string
}

func (err InvalidEmailError) Error() string {
	return fmt.Sprintf("Invalid Email %s", err.email)
}

// Email encapsulates a normalized and validated email address, obtained using
// ParseEmail(). Once normalized, funtionality equivalent email addresses will
// be equal per the == operator.  For example, ox.to.a.cart@gmail.com will equal
// oxtoacart@gmail.com.
type Email struct {
	n string
}

// ParseEmail validates and normalizes an email address, with the following
// transformations applied:
//
//   - lowercase
//   - '.' characters removed from the username part
//   - plus-extensions removed from the username
//   - domains in domainsMap are remapped (for example 'googlemail.com' is
//     remapped to 'gmail.com')
//
// If either the supplied email can't be parsed by net/mail, this function
// returns an InvalidEmailError.
//
// TODO - at the moment, this does not handle quoted strings correctly, for
// example "me@department"@company.com.
func ParseEmail(email string) (Email, error) {
	// Check supplied address
	if !isValidEmail(email) {
		return Email{}, &InvalidEmailError{email}
	}

	// Split out username and domain
	parts := strings.Split(strings.ToLower(email), "@")

	// Clean up username
	username := strings.Replace(strings.Split(parts[0], "+")[0], ".", "", -1)

	// Remap domain
	domain, remapped := domainsMap[parts[1]]
	if !remapped {
		domain = parts[1]
	}

	return Email{username + "@" + domain}, nil
}

func (e Email) String() string {
	return e.n
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
