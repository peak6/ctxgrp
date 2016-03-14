package httprouter

import (
	"github.com/julienschmidt/httprouter"
	"github.com/peak6/ctxgrp"
	"golang.org/x/net/context"
	"net/http"
)

type RouterAdaptor struct {
	*httprouter.Router
}

func New(r *httprouter.Router) ctxgrp.RouterAdaptor {
	return &RouterAdaptor{r}
}

func (ra *RouterAdaptor) Handle(method string, path string, ctx context.Context, h ctxgrp.Handler) {
	ra.Router.Handle(
		method,
		path,
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			h.ServeHTTP(ParamsContext{ctx, p}, w, r)
		})
}

type ParamsContext struct {
	context.Context
	httprouter.Params
}

func (p ParamsContext) Value(key interface{}) interface{} {
	if k, ok := key.(string); ok {
		if v := p.Params.ByName(k); v != "" {
			return v
		}
	}
	return p.Context.Value(key)
}
