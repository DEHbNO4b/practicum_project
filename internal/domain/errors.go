package domain

import "errors"

var (
	ErrRequired = errors.New("required value")
	ErrNotFound = errors.New("not found")

	ErrWrongLoginOrPassword = errors.New("wrong login or password")
	ErrUniqueViolation      = errors.New("unique violation")
	//order errors
	ErrAccepted         = errors.New("accepted for processing")
	ErrHasBeenUpploaded = errors.New("has already been uploaded by another user")
	ErrNilData          = errors.New("nil data")
)
