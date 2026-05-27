package di

// App is the name of the application scope.
const App = "app"

// Request is the name of the request scope.
const Request = "request"

// SubRequest is the name of the subrequest scope.
const SubRequest = "subrequest"

// ScopeList is a slice of scope.
type ScopeList []string

// Copy returns a copy of the ScopeList.
func (l ScopeList) Copy() ScopeList { _ = "STUB: not implemented"; return *new(ScopeList) }

// ParentScopes returns the scopes before the one given as parameter.
func (l ScopeList) ParentScopes(scope string) ScopeList {
	_ = "STUB: not implemented"
	return *new(ScopeList)
}

// SubScopes returns the scopes after the one given as parameter.
func (l ScopeList) SubScopes(scope string) ScopeList {
	_ = "STUB: not implemented"
	return *new(ScopeList)
}

// Contains returns true if the ScopeList contains the given scope.
func (l ScopeList) Contains(scope string) bool { _ = "STUB: not implemented"; return false }
