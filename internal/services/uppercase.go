package services

import (
	"context"
	"strings"
)

type Uppercase struct {
	trimSvc *Trim
}

func NewUppercase(trimSvc *Trim) *Uppercase {
	return &Uppercase{
		trimSvc: trimSvc,
	}
}

func (l *Uppercase) Do(ctx context.Context, input string) (string, error) {
	trimed, err := l.trimSvc.Do(ctx, input)
	if err != nil {
		return "", err
	}
	lowered := strings.ToUpper(trimed)
	return lowered, nil
}
