package app

import (
	"fmt"
	"github.com/pkg/errors"
)

var (
	ErrMovieDoesntExist = errors.New("movie doesn't exist")
)

type errorWrapper struct {
	appError  error
	originErr error
}

func Wrap(appErr error, originErr error) error {
	if originErr == nil {
		return nil
	}

	if appErr == nil {
		return originErr
	}

	return errors.WithStack(&errorWrapper{
		appError:  appErr,
		originErr: originErr,
	})
}

func (e errorWrapper) Error() string {
	return fmt.Sprintf("%s: %s", e.appError, e.originErr)
}

func (e errorWrapper) Cause() error {
	return e.appError
}

func (e errorWrapper) Unwrap() error {
	return e.appError
}
