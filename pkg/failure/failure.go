package failure

import (
	"errors"
	"fmt"

	"github.com/mikestefanello/pagoda/pkg/tracing"
)

const traceLevel = 4

// ErrWithTrace wrap useful information for debugging such as stack stace and function name
func ErrWithTrace(err error) error {
	return fmt.Errorf("%w \n at %s", err, tracing.Trace(traceLevel))
}

// IsError check if error type match or wrap the input type, similar to catching custom exception
func IsError[T error](originError error) bool {
	var errTypeValue T
	return errors.As(originError, &errTypeValue)
}

// AsError similar to IsError but also return the original error value
func AsError[T error](originError error) (T, bool) {
	var errTypeValue T
	ok := errors.As(originError, &errTypeValue)
	return errTypeValue, ok
}
