package helpers

import "errors"

var ErrorInvalidToken = errors.New("invalid group id")

var ErrorInvalidObject = errors.New("invalid object")

var ErrorKeyboardContainTooMuchRows = errors.New("buttons contain too much rows")
var ErrorKeyboardContainsTooMuchButtons = errors.New("keyboard contains too much buttons")
