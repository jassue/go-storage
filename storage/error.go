package storage

import "errors"

var (
    FileNotFoundErr = errors.New("file not found")
    FileNoPermissionErr = errors.New("permission denied")
)
