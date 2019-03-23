package repository

import "context"

// Reserve ...
type Reserve interface {
	Put(ctx context.Context)
	SetStarted(ctx context.Context)
	SetFinished(ctx context.Context)
	SetRevoked(ctx context.Context)
}
