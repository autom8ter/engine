package lis

import (
	"github.com/soheilhy/cmux"
	"log"
	"net"
)

type Listener struct {
	mux cmux.CMux
	lis net.Listener
}

func NewListner(lis net.Listener) *Listener {
	return &Listener{
		mux: cmux.New(lis),
		lis: lis,
	}
}

func (l *Listener) Serve() {
	log.Printf("starting %s", l.lis.Addr())
	err := l.mux.Serve()
	log.Printf("closed: %v", err)
}

func (l *Listener) GRPCListener() net.Listener {
	return l.mux.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
}

func (l *Listener) HTTPListener() net.Listener {
	return l.mux.Match(cmux.HTTP2(), cmux.HTTP1Fast())
}
