package service

import "errors"

type ErrorString string

const (
	ERR_INVALID_ADDRESS string = "ERR_INVALID_ADDRESS"
	ELECTRUM_ERROR      string = "ELECTRUM_ERROR"
)

var (
	ErrInvalidAddress = errors.New(ERR_INVALID_ADDRESS)
)

type ServiceErr struct {
	Text string
	Msg  string
}

func (e ServiceErr) Error() string {
	return e.Msg
}

func InvalidAddressError(err error) *ServiceErr {
	return &ServiceErr{
		Text: ERR_INVALID_ADDRESS,
		Msg:  err.Error(),
	}
}

func ElectrumError(err error) *ServiceErr {
	return &ServiceErr{
		Text: ELECTRUM_ERROR,
		Msg:  err.Error(),
	}
}
