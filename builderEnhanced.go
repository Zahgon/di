package di

const generatedNamePrefix = "_di_generated_"

// EnhancedBuilder can be used to create a Container.
// The EnhancedBuilder should be created with NewEnhancedBuilder.
// Then you can add definitions with the Add method.
// Once all the definitions have been added to the builder,
// you can generate the Container with the Build method.
//
// It works the same way as the basic Builder. But the definitions given to the Add method are pointers.
// The definitions are updated when the Build method is called.
// That allows to retrieve objects by their definitions which is faster than retrieving them by name.
type EnhancedBuilder struct {
	definitions    DefMap
	bindings       map[string]*Def
	insertionOrder map[string]int
	numAdded       int
	scopes         ScopeList
}

// NewEnhancedBuilder is the only way to create a working EnhancedBuilder.
// It initializes an EnhancedBuilder with a list of scopes.
// The scopes are ordered from the most generic to the most specific.
// If no scope is provided, the default scopes are used:
// [App, Request, SubRequest]
// It can return an error if the scopes are not valid.
func NewEnhancedBuilder(scopes ...string) (*EnhancedBuilder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func checkBuilderScopes(scopes []string) error { _ = "STUB: not implemented"; return nil }

// Scopes returns the list of available scopes.
func (b *EnhancedBuilder) Scopes() ScopeList { _ = "STUB: not implemented"; return *new(ScopeList) }

// Definitions returns a map with the all objects definitions registered at this point.
// The key of the map is the name of the definition.
func (b *EnhancedBuilder) Definitions() DefMap { _ = "STUB: not implemented"; return *new(DefMap) }

// NameIsDefined returns true if there is a definition registered with the given name.
func (b *EnhancedBuilder) NameIsDefined(name string) bool { _ = "STUB: not implemented"; return false }

// Add adds one definition to the Builder.
// It returns an error if the definition can not be added.
//
// The name must be unique. If a definition with the same name has already been added,
// it will be replaced by the new one, as if the first one never was added.
// If an empty name is provided, a name starting with "_di_generated_" is generated.
// You can not add a definition with a name starting with "_di_generated_" as it is reserved for auto-genrated ones.
// Providing a name is recommended as it makes errors much easier to understand.
//
// The input definition is a pointer.
// It will be updated when the container is generated with the Build method.
// It binds the definition to the generated Container.
// That allows to build an object not only from its name
// but also from its definition which happens to be faster.
func (b *EnhancedBuilder) Add(def *Def) error { _ = "STUB: not implemented"; return nil }

// Build creates a Container in the most generic scope
// with all the definitions registered in the builder.
//
// The definition provided in the Add method
// are updated to match their state when they were added to the builder.
//
// A definition can only belong to one container.
// That means you can only call Build once.
func (b *EnhancedBuilder) Build() (Container, error) {
	_ = "STUB: not implemented"
	return *new(Container), nil
}

// Update definition scopes.

// Put definitions in a slice and sort them by insertion order.

// Generate the indexes based on the definitions.

// Update the bound fields of the definition.

// Update indexes and definitionScopeLevels slices.

// Update the bound definition.
