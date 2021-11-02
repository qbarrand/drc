package pkg

import "context"

type Desk interface {
	GetCurrentHeight(ctx context.Context) (int, error)
}
