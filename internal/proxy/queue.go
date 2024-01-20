package proxy

import (
	"context"
	"github.com/imneov/modelmesh/internal/broker/config"
)

type QueueSelect func() (serviceGroup string, err error)

type PredictRequest struct {
}

type Queue struct {
	ch chan *PredictRequest
}

func NewQueue(cfg *config.Queue) (*Queue, error) {
	size := cfg.Size
	return &Queue{ch: make(chan *PredictRequest, size)}, nil
}

func (q *Queue) Push(ctx context.Context, in *PredictRequest) {
	q.ch <- in
}

func (q *Queue) Pop(ctx context.Context) *PredictRequest {
	select {
	case pr := <-q.ch:
		return pr
	default:
		return nil
	}
}
