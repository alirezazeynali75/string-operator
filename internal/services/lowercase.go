package services

import (
	"context"
	"strings"
)

type Lowercase struct {
	trimSvc *Trim
}

func NewLowercase(trimSvc *Trim) *Lowercase {
	return &Lowercase{
		trimSvc: trimSvc,
	}
}

func (l *Lowercase) Do(ctx context.Context, input string) (string, error) {
	trimed, err := l.trimSvc.Do(ctx, input)
	if err != nil {
		return "", err
	}
	lowered := strings.ToLower(trimed)
	return lowered, nil
}
