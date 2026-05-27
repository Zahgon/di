package di

// Builder can be used to create a Container.
// The Builder should be created with NewBuilder.
// Then you can add definitions with the Add method,
// and finally build the Container with the Build method.
//
// Consider using the EnhancedBuilder that provides more features.
type Builder struct {
	definitions    DefMap
	scopes         ScopeList
	insertionOrder map[string]int
	numAdded       int
}

// NewBuilder is the only way to create a working Builder.
// It initializes a Builder with a list of scopes.
// The scopes are ordered from the most generic to the most specific.
// If no scope is provided, the default scopes are used:
// [App, Request, SubRequest]
// It can return an error if the scopes are not valid.
func NewBuilder(scopes ...string) (*Builder, error) { _ = "STUB: not implemented"; return nil, nil }

func checkScopes(scopes []string) error { _ = "STUB: not implemented"; return nil }

// Scopes returns the list of available scopes.
func (b *Builder) Scopes() ScopeList { _ = "STUB: not implemented"; return *new(ScopeList) }

// Definitions returns a map with the all the objects definitions
// registered with the Add method.
// The key of the map is the name of the Definition.
func (b *Builder) Definitions() DefMap { _ = "STUB: not implemented"; return *new(DefMap) }

// IsDefined returns true if there is a definition with the given name.
func (b *Builder) IsDefined(name string) bool { _ = "STUB: not implemented"; return false }

// Add adds one or more definitions in the Builder.
// It returns an error if a definition can not be added.
// If a definition with the same name has already been added,
// it will be replaced by the new one, as if the first one never existed.
func (b *Builder) Add(defs ...Def) error { _ = "STUB: not implemented"; return nil }

func (b *Builder) add(def Def) error { _ = "STUB: not implemented"; return nil }

// note that an empty scope is allowed
// it will be replaced in the Build method by the most generic scope

// Set is a shortcut to add a definition for an already built object.
func (b *Builder) Set(name string, obj interface{}) error { _ = "STUB: not implemented"; return nil }

// Build creates a Container in the most generic scope
// with all the definitions registered in the Builder.
func (b *Builder) Build() Container { _ = "STUB: not implemented"; return *new(Container) }

// Update definition scopes.

// Put definitions in a slice and sort them by insertion order.

// Generate the indexes based on the definitions.

// Update the definition bound fields.

// Update indexes and definitionScopeLevels slices.
