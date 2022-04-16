package grpc

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

type IClientWrapper interface {
	Start() error
	Stop() error
}

type ClientWrapper struct {
	c              context.Context
	cancel         context.CancelFunc
	DefaultTimeout time.Duration
	gconn          grpc.ClientConn
}
