package ctxgrp

import (
	"net/http"

	"golang.org/x/net/context"
)

type RouterAdaptor interface {
	// Handle is passed a base context to be used in creating your own context values
	Handle(method string, path string, ctx context.Context, h Handler)
}

type Handler interface {
	ServeHTTP(ctx context.Context, w http.ResponseWriter, r *http.Request)
}
type HandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request)

func (hf HandlerFunc) ServeHTTP(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	hf(ctx, w, r)
}

type Group struct {
	r  RouterAdaptor
	p  string
	mw Middleware
}

func NewGroup(r RouterAdaptor, path string) *Group {
	return &Group{r, path, nil}
}

func (g *Group) Use(cons ...Constructor) *Group {
	return &Group{g.r, g.p, g.mw.Use(cons...)}
}

func (g *Group) NewGroup(path string) *Group {
	return &Group{g.r, mkpath(g.p, path), g.mw}
}

func (g *Group) Handle(method string, path string, hf Handler) {
	g.r.Handle(method, mkpath(g.p, path), context.Background(), g.mw.Then(hf))
}

func (g *Group) GET(path string, hf HandlerFunc) {
	g.Handle("GET", path, hf)
}
func (g *Group) PUT(path string, hf HandlerFunc) {
	g.Handle("PUT", path, hf)
}
func (g *Group) POST(path string, hf HandlerFunc) {
	g.Handle("POST", path, hf)
}
func (g *Group) DELETE(path string, hf HandlerFunc) {
	g.Handle("DELETE", path, hf)
}
func (g *Group) OPTIONS(path string, hf HandlerFunc) {
	g.Handle("OPTIONS", path, hf)
}
func (g *Group) HEAD(path string, hf HandlerFunc) {
	g.Handle("HEAD", path, hf)
}

func mkpath(pre string, post string) string {
	if pre == "" {
		if post == "" {
			panic("Cannot map empty path")
		}
		if post[0] != '/' {
			panic("Path must start with /")
		}
		return post
	}
	if pre[0] != '/' {
		panic("Path must start with /")
	}
	if post == "" {
		return pre
	}
	if post[0] == '/' && pre[len(pre)-1] == '/' {
		return pre + post[1:]
	} else {
		return pre + post
	}
}

func HttpHandler(h http.Handler) HandlerFunc {
	return func(_ context.Context, w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}
