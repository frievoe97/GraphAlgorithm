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
