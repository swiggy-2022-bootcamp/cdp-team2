package grpc

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

type IClientConn interface {
	// Start() error
	CtxWithTimeOut() context.Context
	Stop() error
}

type ClientConn struct {
	c              context.Context
	cancel         context.CancelFunc
	DefaultTimeout time.Duration
	gconn          *grpc.ClientConn
}

func (cc *ClientConn) CtxWithTimeOut() context.Context {
	ctx, _ := context.WithTimeout(cc.c, cc.DefaultTimeout)
	return ctx
}

func (cc *ClientConn) Stop() {
	cc.cancel()
	cc.gconn.Close()
	log.Println("Stopping grpc client")
}
