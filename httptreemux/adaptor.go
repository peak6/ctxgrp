package httptreemux

import (
	"net/http"

	"github.com/dimfeld/httptreemux"
	"github.com/peak6/ctxgrp"
	"golang.org/x/net/context"
)

// TreeMuxAdaptor combines ctxgrp.Group and httptreemux.TreeMux
type TreeMuxAdaptor struct {
	*httptreemux.TreeMux
	*ctxgrp.Group
}

// New constructs a new TreeMuxAdaptor
func New() *TreeMuxAdaptor {
	var ret TreeMuxAdaptor
	ret.TreeMux = httptreemux.New()
	ret.Group = ctxgrp.NewGroup(&ret, "/")
	return &ret
}

// Handle impments ctxgroup.RouterAdaptor
func (tma *TreeMuxAdaptor) Handle(method string, path string, ctx context.Context, h ctxgrp.Handler) {
	tma.TreeMux.Handle(
		method,
		path,
		func(w http.ResponseWriter, r *http.Request, p map[string]string) {
			h.ServeHTTP(mapContext{ctx, p}, w, r)
		})
}

type mapContext struct {
	context.Context
	m map[string]string
}

func (m mapContext) Value(key interface{}) interface{} {
	if k, ok := key.(string); ok {
		if v, ok := m.m[k]; ok {
			return v
		}
	}
	return m.Context.Value(key)
}
