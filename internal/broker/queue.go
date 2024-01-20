package broker

import (
	"context"
	"github.com/imneov/modelmesh/internal/broker/config"
	errs "github.com/imneov/modelmesh/internal/broker/error"
)

type QueueSelect func() (serviceGroup string, err error)

type QueuePool struct {
	queues map[string]*Queue
}

type Queue struct {
	ch chan *PredictRequest
}

func NewQueuePool(cfg *config.Queue, sc []*config.ServiceGroup) (*QueuePool, error) {
	size := cfg.Size
	qp := &QueuePool{
		queues: map[string]*Queue{},
	}
	for _, group := range sc {
		qp.queues[group.Name] = &Queue{ch: make(chan *PredictRequest, size)}
	}
	return qp, nil
}

func NewQueue(cfg *config.Queue) (*Queue, error) {
	size := cfg.Size
	return &Queue{ch: make(chan *PredictRequest, size)}, nil
}

func (qp *QueuePool) Select(serviceGroup string) (queue *Queue, err error) {
	queue, ok := qp.queues[serviceGroup]
	if !ok {
		return nil, errs.ErrQueueGroupIsNotExist
	}
	return queue, nil
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
