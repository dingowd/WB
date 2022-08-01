package provider

import "context"

type MessageProvider interface {
	Start(ctx context.Context)
	Stop(ctx context.Context)
}
