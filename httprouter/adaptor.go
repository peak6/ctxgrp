package httprouter

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/peak6/ctxgrp"
	"golang.org/x/net/context"
)

// RouterAdaptor combines ctxgrp.Group and httprouter.Router
type RouterAdaptor struct {
	*httprouter.Router
	*ctxgrp.Group
}

// New constructs a new RouterAdaptor
func New() *RouterAdaptor {
	var ret RouterAdaptor
	ret.Router = httprouter.New()
	ret.Group = ctxgrp.NewGroup(&ret, "/")
	return &ret
}

// Handle impments ctxgroup.RouterAdaptor
func (ra *RouterAdaptor) Handle(method string, path string, ctx context.Context, h ctxgrp.Handler) {
	ra.Router.Handle(
		method,
		path,
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			h.ServeHTTP(paramsContext{ctx, p}, w, r)
		})
}

type paramsContext struct {
	context.Context
	httprouter.Params
}

func (p paramsContext) Value(key interface{}) interface{} {
	if k, ok := key.(string); ok {
		if v := p.Params.ByName(k); v != "" {
			return v
		}
	}
	return p.Context.Value(key)
}
