package di

// graph is a Directed Acyclic Graph.
// It is used to store the dependencies inside a container.
// These dependencies are then used to determine the order
// that should be used to close the objects.
type graph struct {
	verticeSlice []int
	vertices     map[int]*graphVertex
}

// graphVertex contains the vertex data.
type graphVertex struct {
	// numIn in the number of incoming edges.
	numIn int
	// numInTmp is used by the TopologicalOrdering to avoid messing with numIn
	numInTmp int
	// out contains the name the outgoing edges.
	out []int
	// outMap is the same as "out", but in a map
	// to quickly check if a vertex is in the outgoing edges.
	outMap map[int]struct{}
}

// newGraph creates a new graph.
func newGraph() *graph { _ = "STUB: not implemented"; return nil }

// AddVertex adds a vertex to the graph.
func (g *graph) AddVertex(v int) { _ = "STUB: not implemented"; return }

// AddEdge adds an edge to the graph.
func (g *graph) AddEdge(from, to int) { _ = "STUB: not implemented"; return }

// check if the edge is already registered

// update the vertices

// TopologicalOrdering returns a valid topological sort.
// It implements Kahn's algorithm.
// If there is a cycle in the graph, an error is returned.
// The list of vertices is also returned even if it is not ordered.
func (g *graph) TopologicalOrdering() ([]int, error) { _ = "STUB: not implemented"; return nil, nil }

// multiErrBuilder can accumulate errors.
type multiErrBuilder struct {
	errs []error
}

// Add adds an error in the multiErrBuilder.
func (b *multiErrBuilder) Add(err error) { _ = "STUB: not implemented"; return }

// Build returns an errors containing all the messages
// of the accumulated errors. If there is no error
// in the builder, it returns nil.
func (b *multiErrBuilder) Build() error { _ = "STUB: not implemented"; return nil }

// fill copies src in dest. dest should be a pointer to src type.
func fill(src, dest interface{}) (err error) { _ = "STUB: not implemented"; return nil }
