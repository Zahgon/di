package di

// buildingChan is used internally as the value of an object while it is being built.
type buildingChan chan struct{}

// buildObject wraps the Build function to recover from a panic.
func buildObject(
	buildFunc func(ctn Container) (interface{}, error),
	ctn Container,
	index int,
	defName string,
) (obj interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// formatBuiltOnClosedContainerError formats the error that happens when you try to build an object with a closed container.
func formatBuiltOnClosedContainerError(def Def, closeObjectErr error) error {
	_ = "STUB: not implemented"
	return nil
}

// formatCycleError formats the error that happens when a cycle is detected.
func formatCycleError(ctn Container, def Def) error { _ = "STUB: not implemented"; return nil }

// Fill is similar to SafeGet but it does not return the object.
// Instead it fills the provided object with the value returned by SafeGet.
// The provided object must be a pointer to the value returned by SafeGet.
// It uses reflection so it is slower than Get and SafeGet.
// But it can be convenient in some cases where performance is not a critical factor.
func (ctn Container) Fill(in interface{}, dst interface{}) error {
	_ = "STUB: not implemented"
	return nil
}
