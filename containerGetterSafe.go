package di

// SafeGet retrieves an object from the Container.
// The object has to belong to the Container or one of its parents.
// If the object does not already exist, it is created and saved in the Container.
// If the object can not be created, it returns an error.
//
// There are different ways to retrieve an object.
//   - From its name: ctn.SafeGet("object-name")
//   - From its definition: ctn.SafeGet(objectDef) or ctn.SafeGet(objectDefPtr) - only with the EnhancedBuilder
//   - From its index: ctn.SafeGet(objectDef.Index()) - only with the EnhancedBuilder
//   - From its type: ctn.SafeGet(reflect.typeOf(MyObject{})) - only if objectDef.Is includes the given type
//     In case there are more than one definition matching the given type,
//     the chosen one is the last definition inserted in the builder.
func (ctn Container) SafeGet(in interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Finding the right core.

// Try to fetch an already built object as quickly as possible.

// Reset the builtList if the scope changed.

// Retrieve the definition.

// Cycle detection.

// Handle unshared objects.

// Handle shared objects.

// Check again if the object was created, with the lock this time.

// Wait for the object to be created by another call to SafeGet.
// Can not get the object without calling SafeGet again as its creation may have failed.

// Mark the object as building.
// And release the lock as it can take a while to create the object.

// Building the shared object.

// The object could not be created. Remove the building channel from the container
// and close it to allow the object to be created again.

// The container has been deleted while the object was being built.
// The newly created object needs to be closed, and it will not be returned.
