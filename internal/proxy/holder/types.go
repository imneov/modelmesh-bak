package holder

import (
	"context"
)

type Status string

var (
	StatusCanceled Status = "canceled"
	StatusOK       Status = "ok"
)

type Holder interface {
	Cancel()
	Wait(ctx context.Context, id string) *Waiter
	OnRespond(*Response)
}

type Response struct {
	ID       string
	Status   Status
	ErrCode  string
	Metadata map[string]string
	Data     []byte
}

type Waiter struct {
	ch chan Response
}

func (w *Waiter) Wait() Response {
	resp := <-w.ch
	return resp
}
