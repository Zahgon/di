package di

import (
	"net/http"
)

// ContainerKey is a type that can be used to store a container
// in the context.Context of an http.Request.
// By default, it is used in the C function and the HTTPMiddleware.
type ContainerKey string

// HTTPMiddleware adds a container in the request context.
//
// The container injected in each request, is a new sub-container
// of the app container given as parameter.
//
// It can panic, so it should be used with another middleware
// to recover from the panic, and to log the error.
//
// It uses logFunc, a function that can log an error.
// logFunc is used to log the errors during the container deletion.
func HTTPMiddleware(h http.HandlerFunc, app Container, logFunc func(msg string)) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// create a request container from tha app container

// call the handler with a new request
// containing the container in its context

// C retrieves a Container from an interface.
// The function panics if the Container can not be retrieved.
//
// The interface can be :
//   - a Container
//   - an *http.Request containing a Container in its context.Context
//     for the ContainerKey("di") key.
//
// The function can be changed to match the needs of your application.
var C = func(i interface{}) Container {
	if c, ok := i.(Container); ok {
		return c
	}

	r, ok := i.(*http.Request)
	if !ok {
		panic("could not get the container with C()")
	}

	c, ok := r.Context().Value(ContainerKey("di")).(Container)
	if !ok {
		panic("could not get the container from the given *http.Request")
	}

	return c
}

// Get is a shortcut for C(i).Get(name).
func Get(i interface{}, name string) interface{} { _ = "STUB: not implemented"; return nil }
