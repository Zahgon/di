package di

// ParentContainer returns the parent Container.
// If the Container does not have a parent, it returns an error.
func (ctn Container) ParentContainer() (Container, error) {
	_ = "STUB: not implemented"
	return *new(Container), nil
}

// Parent returns the parent Container.
// It works like ParentContainer but without the error.
// This method was kept to have some kind of backward compatibility.
func (ctn Container) Parent() Container { _ = "STUB: not implemented"; return *new(Container) }

// SubContainer creates a new Container in the next sub-scope
// that will have this Container as parent.
func (ctn Container) SubContainer() (Container, error) {
	_ = "STUB: not implemented"
	return *new(Container), nil
}
