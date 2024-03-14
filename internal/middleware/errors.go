package middleware

import "errors"

var ErrNoAccess = errors.New("you have no access")
var ErrSomethingWrong = errors.New("something wrong")
