package holder

import (
	"context"
	"k8s.io/klog/v2"
	"sync"
)

type holder struct {
	lock    sync.RWMutex
	holders map[string]chan Response
}

func New(ctx context.Context) Holder {
	return &holder{
		lock:    sync.RWMutex{},
		holders: make(map[string]chan Response),
	}
}

// Cancel cancel all waiters.
func (h *holder) Cancel() {
	h.lock.Lock()
	for _, ch := range h.holders {
		ch <- Response{
			Status:  StatusCanceled,
			ErrCode: "canceled",
		}
	}
	h.lock.Unlock()
}

func (h *holder) Wait(ctx context.Context, id string) Response {
	w := h.wait(ctx, id)
	resp := <-w.ch
	return resp
}

func (h *holder) wait(ctx context.Context, id string) *Waiter {
	h.lock.Lock()
	// chan size 为3 避免response和超时对chan的竞争.
	waitCh := make(chan Response, 2)
	h.holders[id] = waitCh
	h.lock.Unlock()

	go func() {
		select {
		case <-ctx.Done():
			klog.V(0).Infof("receive context done, id: %s", id)
			if nil != ctx.Err() {
				waitCh <- Response{
					Status:  StatusCanceled,
					ErrCode: context.Canceled.Error(),
				}
			}
		}

		// delete wait channel.
		h.lock.Lock()
		delete(h.holders, id)
		h.lock.Unlock()
	}()

	return &Waiter{
		ch: waitCh,
	}
}

func (h *holder) Notify(resp *Response) {
	klog.V(0).InfoS("received response", "respID", resp.ID, "status", resp.Status)

	h.lock.Lock()
	waitCh := h.holders[resp.ID]
	delete(h.holders, resp.ID)
	h.lock.Unlock()

	if nil == waitCh {
		klog.Warningf("holders request %s is not exists", resp.ID)
		return
	}

	waitCh <- *resp
}
