package exception

import "fmt"

var (
	ErrRepositoryNotFound = fmt.Errorf("record not found in repository")
)
