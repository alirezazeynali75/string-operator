package services

import (
	"context"

)


var counter int

type Increase struct {}


func NewIncrease() *Increase {
	return &Increase{}
}

func (i *Increase) Inc(ctx context.Context) error {
	counter = counter + 1
	return nil
}
