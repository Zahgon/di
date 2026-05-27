package di

// DeleteWithSubContainers takes all the objects saved in this Container
// and calls the Close function of their Definition on them.
// It will also call DeleteWithSubContainers on each child and remove its reference in the parent Container.
// After deletion, the Container can no longer be used.
// The sub-containers are deleted even if they are still used in other goroutines.
// It can cause errors. You may want to use the Delete method instead.
func (ctn Container) DeleteWithSubContainers() error { _ = "STUB: not implemented"; return nil }

// Delete works like DeleteWithSubContainers if the Container does not have any child.
// But if the Container has sub-containers, it will not be deleted right away.
// The deletion only occurs when all the sub-containers have been deleted manually.
// So you have to call Delete or DeleteWithSubContainers on all the sub-containers.
func (ctn Container) Delete() error { _ = "STUB: not implemented"; return nil }

// Clean deletes the sub-container created by UnscopedSafeGet, UnscopedGet or UnscopedFill.
func (ctn Container) Clean() error { _ = "STUB: not implemented"; return nil }

// IsClosed returns true if the Container has been deleted.
func (ctn Container) IsClosed() bool { _ = "STUB: not implemented"; return false }

func deleteContainerCore(core *containerCore) error { _ = "STUB: not implemented"; return nil }

// Stop returning the already built objects.

// Delete clone.

// Remove from parent.

// Close objects in the right order.

func closeObject(obj interface{}, closeFunc func(interface{}) error, defName string) (err error) {
	_ = "STUB: not implemented"
	return nil
}
