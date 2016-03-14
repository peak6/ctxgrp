package ctxgrp

type Middleware []MiddlewareFunc
type MiddlewareFunc func(hf Handler) HandlerFunc

func copyof(src Middleware) Middleware {
	dest := make(Middleware, len(src))
	copy(dest, src)
	return dest
}

func (m Middleware) Use(fns ...MiddlewareFunc) Middleware {
	return append(copyof(m), fns...)
}

// Then creates a handler with all the middleware applied
func (m Middleware) Then(hf Handler) Handler {
	for _, mf := range m {
		hf = mf(hf)
	}
	return hf
}
