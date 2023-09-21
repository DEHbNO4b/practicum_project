package domain

import "errors"

var (
	ErrRequired = errors.New("required value")
	ErrNotFound = errors.New("not found")

	ErrWrongLoginOrPassword = errors.New("wrong login or password")
	ErrUniqueViolation      = errors.New("unique violation")

	//order errors
	ErrAccepted = errors.New("accepted for processing")
	ErrConflict = errors.New("has already been uploaded by another user")
	ErrNilData  = errors.New("nil data")

	//balance errors
	ErrNotEnaugh            = errors.New("there are not enough funds on the account")
	ErrIncorrectOrderNumber = errors.New("incorrect order number")

	//accrual system errors
	ErrNotRegistered        = errors.New("order not registered in accrual system")
	ErrTooManyRequests      = errors.New("too many requests to accrual system")
	ErrUnexpectedRespStatus = errors.New("unexpected response status")
)
