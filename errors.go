package main

import "fmt"

type NewError struct {
	Message string
}

func (err *NewError) Error() string {
	return fmt.Sprintf("Error: %s", err.Message)
}
