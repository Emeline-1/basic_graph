package basic_graph

import ("container/list")

type Graph struct {
	nodes map[string]struct{}
	edges map[string]map[string]struct{}

	seen map[string]struct{}
	directed_component map[string]struct{}
	nodes_s []string
	curr_node int
}

func New () *Graph {
	return &Graph{nodes: make (map[string]struct{}), edges: make (map[string]map[string]struct{})}
}

func (g *Graph) Add_edge (n1, n2 string) bool {
	g.nodes[n1] = struct{}{}
	g.nodes[n2] = struct{}{}

	v := g._add_edge (n1, n2)
	g._add_edge (n2, n1)
	return v
}

/**
 * false for edges that were already present, or edges that are entirely new (new directed component)
 * true for a new edge with an already existing node.
 */
func (g *Graph) _add_edge (n1, n2 string) bool {
	if e, ok := g.edges[n1]; ok { // n1 was already present in graph and had edges
		if _, ok2 := e[n2]; ok2 {
			return false
		}
		e[n2] = struct{}{}
		return true // sign of a merged overlay group
	} else {
		g.edges[n1] = map[string]struct{}{n2: struct{}{}}
		return false
	}
}

func (g *Graph) Set_iterator () {
	g.seen = make (map[string]struct{})
	g.nodes_s = get_keys (g.nodes)
	g.curr_node = 0
}

func (g *Graph) Next_connected_component () bool {

	if g.curr_node >= len(g.nodes_s) {
		return false
	}
	curr_node := g.nodes_s[g.curr_node]
	if _, ok := g.seen[curr_node]; !ok {
		g.directed_component = g.bfs (curr_node)
		update (g.seen, g.directed_component)
		g.curr_node += 1
		return true

	} else { // Node already recorded in another connected component
		g.curr_node += 1
		return g.Next_connected_component ()
	}
	
}

func (g *Graph) Connected_component () []string {
	return get_keys (g.directed_component)
}

func (g *Graph) bfs (node string) map[string]struct{} {
	seen := make (map[string]struct{})

	fifo_queue := list.New ()
	fifo_queue.PushBack (node)

	for fifo_queue.Len () != 0 {
		curr := fifo_queue.Front () 
		fifo_queue.Remove (curr)
		curr_value, _ := curr.Value.(string)

		if _, ok := seen[curr_value]; !ok {
			seen[curr_value] = struct{}{}
			for neighbor, _ := range g.edges[curr_value] {
				fifo_queue.PushBack (neighbor)
			}
		}
	}
	return seen
}

func update (to_update, s map[string]struct{}) {
	for k,v := range s {
		to_update[k] = v
	}
}

func get_keys (mymap map[string]struct{}) []string {
	keys := make([]string, len(mymap))
	i := 0
	for k := range mymap {
	    keys[i] = k
	    i++
	}
	return keys
}
