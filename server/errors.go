package server

import (
	"errors"

	errors2 "emperror.dev/errors"
)

var (
	ErrIsRunning            = errors2.New("server is running")
	ErrSuspended            = errors2.New("server is currently in a suspended state")
	ErrServerIsInstalling   = errors2.New("server is currently installing")
	ErrServerIsTransferring = errors2.New("server is currently being transferred")
	ErrServerIsRestoring    = errors2.New("server is currently being restored")
)

type crashTooFrequent struct{}

func (e *crashTooFrequent) Error() string {
	return "server has crashed too soon after the last detected crash"
}

func IsTooFrequentCrashError(err error) bool {
	var crashTooFrequent *crashTooFrequent
	ok := errors.As(err, &crashTooFrequent)

	return ok
}

type serverDoesNotExist struct{}

func (e *serverDoesNotExist) Error() string {
	return "server does not exist on remote system"
}

func IsServerDoesNotExistError(err error) bool {
	var serverDoesNotExist *serverDoesNotExist
	ok := errors.As(err, &serverDoesNotExist)

	return ok
}
