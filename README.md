# basic_graph

This project implements a very basic graph in the [Go Language](https://golang.org).

It allows to:
* Create a new instance of a graph,
* Add edges to this graph,
* And get all connected components of the graph.

#### Example

```
import ("basic_graph",
        "fmt")

g := graph.New ()
g.Add_edge ("node1", "node2")
g.Add_edge ("node1", "node3")
g.Add_edge ("node4", "node5")

g.Set_Iterator ()
for g.Next_connected_component() {
  fmt.Println (g.Connected_component ())
}
```

> Expected output: ["node1", "node2", "node3"] , ["node4", "node5"]
