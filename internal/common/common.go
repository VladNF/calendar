package common

import "context"

type StartStopper interface {
	Start() error
	Stop(ctx context.Context) error
}
