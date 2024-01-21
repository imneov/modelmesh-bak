package e2e

import (
	"context"
	"testing"
)

func Test_Testsuite(t *testing.T) {
	t.Run("Testsuite", func(t *testing.T) {
		ctx := context.Background()
		Testsuite(ctx)
		select {
		case <-ctx.Done():
		}
	})
}
