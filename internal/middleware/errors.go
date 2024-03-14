package middleware

import "errors"

var ErrNoAccess = errors.New("you have no access")
var ErrInternal = errors.New("interal error")
