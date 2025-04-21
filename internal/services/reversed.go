package services

import (
	"context"
)

type Reverse struct {
}

func NewReverse() *Reverse {
	return &Reverse{
	}
}

func (l *Reverse) Do(ctx context.Context, input string) (string, error) {
	reversed := ""
	for _, r := range input {
		reversed = string(r) + reversed
	}
	return reversed, nil
}
