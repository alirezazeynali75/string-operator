package services

import (
	"context"
	"fmt"
	"strings"
)


type Trim struct {}

func NewTrim() *Trim {
	return &Trim{}
}

func (t *Trim) Do(ctx context.Context, input string) (string, error) {
	fmt.Println("Trim step")
	trimed := strings.TrimSpace(input)
	return trimed, nil
}