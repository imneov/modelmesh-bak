package holder

import (
	"context"
	"k8s.io/klog/v2"
	"time"
)

func ExampleHolder() {
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)
	holder := New(ctx, 10)

	w := holder.Wait(ctx, "requestID-test")

	go func() {
		time.Sleep(5 * time.Second)
		holder.OnRespond(&Response{
			ID:     "requestID-test",
			Status: StatusOK,
		})
	}()
	for {
		klog.V(0).Info("wait")
		resp := w.Wait()
		reqID := resp.ID
		//if resp.Status != StatusOK {
		//	klog.Fatal("resp status is not ok", "resp", resp, "reqID", reqID)
		//}

		klog.V(0).Info("resp status is ok", "resp", resp, "reqID", reqID)
		time.Sleep(1 * time.Second)
	}

	// Output:
	//
}
