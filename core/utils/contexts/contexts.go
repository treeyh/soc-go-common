package contexts

import (
	"context"
	"time"
)

type valueOnlyContext struct{ context.Context }

func (valueOnlyContext) Deadline() (deadline time.Time, ok bool) { return }
func (valueOnlyContext) Done() <-chan struct{}                   { return nil }
func (valueOnlyContext) Err() error                              { return nil }

// CloneContext 克隆一个新的ctx
func CloneContext(ctx context.Context) context.Context {
	ctx2 := valueOnlyContext{ctx}
	return ctx2
}
