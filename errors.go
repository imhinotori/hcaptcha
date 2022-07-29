package hcaptcha

import "errors"

type ErrorCode string

const (
	missingInputSecret           = "missing-input-secret"
	invalidInputSecret           = "invalid-input-secret"
	missingInputResponse         = "missing-input-response"
	invalidInputResponse         = "invalid-input-response"
	badRequest                   = "bad-request"
	invalidOrAlreadySeenResponse = "invalid-or-already-seen-response"
	notUsingDummyPasscode        = "not-using-dummy-passcode"
	siteKeySecretMismatch        = "sitekey-secret-mismatch"
)

func (e ErrorCode) Err() error {
	switch e {
	case missingInputSecret:
		return errors.New("your secret key is missing")
	case invalidInputSecret:
		return errors.New("your secret key is invalid or malformed")
	case missingInputResponse:
		return errors.New("the response parameter (verification token) is missing")
	case invalidInputResponse:
		return errors.New("the response parameter (verification token) is invalid or malformed")
	case badRequest:
		return errors.New("the request is invalid or malformed")
	case invalidOrAlreadySeenResponse:
		return errors.New("the response parameter has already been checked, or has another issue")
	case notUsingDummyPasscode:
		return errors.New("you have used a testing sitekey but have not used its matching secret")
	case siteKeySecretMismatch:
		return errors.New("the sitekey is not registered with the provided secret")
	}

	return nil
}
