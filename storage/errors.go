package storage

import "errors"

var (
	ErrTableNotFound = errors.New("required table not found")
	ErrIndexNotFound = errors.New("required index not exists")
)
