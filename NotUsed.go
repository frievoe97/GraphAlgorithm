package main

import "fmt"

func (g Graph) AddDirectedEdge(fromNodeId, toNodeId string) {

	fromNode := g.GetVertex(fromNodeId)
	toNode := g.GetVertex(toNodeId)

	if fromNode != nil && toNode != nil {
		for _, vertex := range fromNode.adjacencyList {
			if vertex.toNode == toNode {
				err := fmt.Errorf("This already exists (%v -> %v)", fromNodeId, toNodeId)
				fmt.Println(err.Error())
				return
			}
		}
		fromNode.adjacencyList = append(fromNode.adjacencyList, &Edge{fromNode: fromNode, toNode: toNode})
	} else {
		err := fmt.Errorf("Invalid Edge (%v -> %v)", fromNodeId, toNodeId)
		fmt.Println(err.Error())
	}

}

func (g Graph) AddUndirectedEdge(fromNodeId, toNodeId string) {

	// Get the nodes
	fromNode := g.GetVertex(fromNodeId)
	toNode := g.GetVertex(toNodeId)

	fromNodeToToNodeExists := true
	toNodeToFromNodeExists := true

	// ...
	if fromNode != nil && toNode != nil {
		for _, vertex := range fromNode.adjacencyList {
			if vertex.toNode == toNode {
				err := fmt.Errorf("This already exists (%v -> %v)", fromNodeId, toNodeId)
				fmt.Println(err.Error())
				fromNodeToToNodeExists = false
			}
		}

		for _, vertex := range toNode.adjacencyList {
			if vertex.toNode == fromNode {
				err := fmt.Errorf("This already exists (%v -> %v)", toNodeId, fromNodeId)
				fmt.Println(err.Error())
				toNodeToFromNodeExists = false
			}
		}

		if fromNodeToToNodeExists {
			fromNode.adjacencyList = append(fromNode.adjacencyList, &Edge{fromNode: fromNode, toNode: toNode})
		}

		if toNodeToFromNodeExists {
			toNode.adjacencyList = append(toNode.adjacencyList, &Edge{fromNode: toNode, toNode: fromNode})
		}

	} else {
		err := fmt.Errorf("Invalid Edge (%v -> %v)", fromNodeId, toNodeId)
		fmt.Println(err.Error())
	}

}

/*
// Package graph implements an adjacency list for a graph representation
// and some selected algorithms on such graphs.
package graph

// Interface of a Graph
// - each node should have an unique nodeId (externally set) of type string
type Graph interface {

	// add an vertex. If the vertex already exists in the graph: do nothing
	AddVertex(nodeId string) //, data interface{})

	// get some generic meta data of the graph:
	NumVertices() int // number of vertices
	NumEdges() int    // number of edges
	//IsDirected() bool // is a directed or undirected graph?

	// some basic algorithms:

	// the returned map maps for each reachable
	// node the distance in layers from given nodeId to each node
	BFS(nodeId string) map[string]int

	// maps the reachable node to true
	DFS(nodeId string) map[string]bool
}

type UnDirectedGraph interface {
	Graph
	AddUndirectedEdge(nodeId1, nodeId2 string, length float64) //, data interface{})

	//Kruskal(graph *AdjacencyList) (*AdjacencyList, float64)
}

type DirectedGraph interface {
	Graph
	Dijkstra(id string) map[string]float64

	//HasVertex(nodeId string) bool
	// the vertices must exists
	AddDirectedEdge(nodeId1, nodeId2 string, length float64) //, data interface{})

	//BellmannFord(id string) ([]float64, bool)
	//FloydWarshall() ([][]float64, bool) // no negative cycles
}

// directed acyclic graph
type DAG interface {
	DirectedGraph
	// algorithms
	// TopoSort returns an map from nodeId to topological order
	TopoSort() map[string]int
}
*/
