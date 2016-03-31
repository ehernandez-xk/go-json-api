package main

import "errors"

type jsonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

// ok represents types capable of validating
// themselves.
type ok interface {
	OK() error
}

//ErrorMissingField Returns error
func ErrorMissingField() error {
	return errors.New("Missing field")
}
