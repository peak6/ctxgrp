package ctxgrp

type Middleware []MiddlewareFunc
type MiddlewareFunc func(hf Handler) Handler

func copyof(src Middleware) Middleware {
	dest := make(Middleware, len(src))
	copy(dest, dest)
	return dest
}

func (m Middleware) Use(fns ...MiddlewareFunc) Middleware {
	ret := copyof(m)
	for _, fn := range fns {
		ret = append(ret, fn)
	}
	return ret
}

// Then creates a handler with all the middleware applied
func (m Middleware) Then(hf Handler) Handler {
	for _, mf := range m {
		hf = mf(hf)
	}
	return hf
}
