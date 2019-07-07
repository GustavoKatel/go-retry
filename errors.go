package retry

import "errors"

var (
	ErrMaxAttemptsReached = errors.New("Max attempts reached")
)
