package graphs

// import "log"
import "github.com/youngzhu/algs4-go/fund"

// Does a given digraph have a directed cycle?
// Solves this problem using depth-first search

type DirectedCycle struct {
	graph Digraph
	marked []bool // marked[v]: has vertex v been marked?
	edgeTo []int // edgeTo[v]: previous vertex on path to v
	onStack []bool // onStack[v]: is vertex on the stack?
	cycle *fund.Stack // directed cycle (or nil if no such cycle)
}

func NewDirectedCycle(g Digraph) *DirectedCycle {
	marked := make([]bool, g.V())
	edgeTo := make([]int, g.V())
	onStack := make([]bool, g.V())
	dc := &DirectedCycle{g, marked, edgeTo, onStack, nil}

	// log.Println(dc.graph.String())

	for v := 0; v < g.V(); v++ {
		if !dc.marked[v] && dc.cycle == nil {
			dc.dfs(v)
		}
	}

	return dc
}

// run DFS and find a directed cycle 
// must use pointer (*), otherwise dc.cycle wouldn't change
func (dc *DirectedCycle) dfs(v int) {
	dc.onStack[v] = true
	dc.marked[v] = true

	
	for _, w := range dc.graph.Adj(v) {
		// short circuit if directed cycle found
		// log.Printf("cycle: %v", dc.cycle!= nil)
		if dc.cycle != nil {
			return
		} else if !dc.marked[w] { // found new vertex, so recur
			dc.edgeTo[w] = v
			dc.dfs(w)
		} else if dc.onStack[w] { // trace back directed cycle
			cycle := fund.NewStack()
			for x := v; x != w; x = dc.edgeTo[x] {
				cycle.Push(x)
			}
			cycle.Push(w)
			cycle.Push(v)
			// log.Println(w)
			dc.cycle = cycle
		}
	}

	dc.onStack[v] = false
}

// Does the digraph have a directed cycle?
func (dc DirectedCycle) HasCycle() bool {
	return dc.cycle != nil 
}

func (dc DirectedCycle) Cycle() []int {
	cycle := make([]int, dc.cycle.Size())

	for i, v := range dc.cycle.Iterator() {
		cycle[i] = v.(int)
	}

	return cycle
}