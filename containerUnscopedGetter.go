package di

// UnscopedSafeGet retrieves an object from the Container, like SafeGet.
// The difference is that the object can be retrieved
// even if it belongs to a more specific scope.
// To do so, UnscopedSafeGet creates a sub-container.
// When the created object is no longer needed,
// it is important to use the Clean method to delete this sub-container.
//
// /!\ Do not use unscope functions inside a `Build` function.
// In this case, circular definitions are not detected. If you do this,
// you take the risk of having an infinite loop in your code when building an object.
func (ctn Container) UnscopedSafeGet(in interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// There was no need to call UnscopedSafeGet, SafeGet was enough.

// UnscopedGet is similar to UnscopedSafeGet but it does not return the error.
// Instead it panics.
func (ctn Container) UnscopedGet(in interface{}) interface{} { _ = "STUB: not implemented"; return nil }

// UnscopedFill is similar to UnscopedSafeGet but copies the object in dst instead of returning it.
func (ctn Container) UnscopedFill(in interface{}, dst interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (ctn Container) getUnscopedChild() (Container, error) {
	_ = "STUB: not implemented"
	return *new(Container), nil
}

func (ctn Container) addUnscopedChild() (Container, error) {
	_ = "STUB: not implemented"
	return *new(Container), nil
}
