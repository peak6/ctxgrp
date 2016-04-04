package ctxgrp

// Middleware is a collection of constructors.
type Middleware []Constructor

// Constructor is a factory method to create a middleware instance wrapping an ctxgrp.Handler
type Constructor func(handler Handler) Handler

func copyof(src Middleware) Middleware {
	dest := make(Middleware, len(src))
	copy(dest, src)
	return dest
}

// Use creates a copy of middleware stack with the supplied constructors appended
func (m Middleware) Use(cons ...Constructor) Middleware {
	return append(copyof(m), cons...)
}

// Then creates a handler with all the middleware applied
func (m Middleware) Then(hf Handler) Handler {
	for _, mf := range m {
		tmp := mf(hf)
		if tmp != nil { // allow middleware to decline
			hf = mf(hf)
		}
	}
	return hf
}
