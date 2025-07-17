package storage

import "errors"

var ErrDataNotFound = errors.New("data not found")
var ErrDataExists = errors.New("data already exists")
