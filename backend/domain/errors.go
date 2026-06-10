package domain

import "errors"

var ErrEmailInUse = errors.New("email already in use")
