package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Graph struct {
	isUndirected bool
	vertecies    []*Vertex
}

func NewGraph() Graph {
	graph := Graph{}
	graph.isUndirected = false
	return graph
}

type Vertex struct {
	nodeId         string
	adjacencyList  []*Edge
	previousVertex *Vertex
	distance       float64
	explored       bool
	layer          int
}

type Edge struct {
	fromNode *Vertex
	toNode   *Vertex
	edgeId   string
	weight   float64
}

func compare(a, b *Vertex) bool {
	if a.distance < b.distance {
		return false
	} else {
		return true
	}
}

func (g *Graph) AddDirectedWeightedEdge(fromNodeId, toNodeId string, weight float64) {

	fromNode := g.GetVertex(fromNodeId)
	toNode := g.GetVertex(toNodeId)

	if fromNode != nil && toNode != nil {
		for _, vertex := range fromNode.adjacencyList {
			if vertex.toNode == toNode {
				fmt.Printf("Die Kante (%v -> %v) existiert bereits.\n", fromNodeId, toNodeId)
				return
			}
		}
		fromNode.adjacencyList = append(fromNode.adjacencyList, &Edge{fromNode: fromNode, toNode: toNode, weight: weight})

		// Check if the graph is now undirected
		if !g.isUndirected {
			for _, node := range toNode.adjacencyList {
				if node.toNode == fromNode {
					g.isUndirected = true
				}
			}
		}

	} else {
		fmt.Printf("Die Kante (%v -> %v) ist ungültig\n", fromNodeId, toNodeId)
	}

}

func (g *Graph) AddUndirectedWeightedEdge(fromNodeId, toNodeId string, weight float64) {

	g.isUndirected = true

	fromNode := g.GetVertex(fromNodeId)
	toNode := g.GetVertex(toNodeId)

	fromNodeToToNodeExists := true
	toNodeToFromNodeExists := true

	if fromNode != nil && toNode != nil {
		for _, vertex := range fromNode.adjacencyList {
			if vertex.toNode == toNode {
				fmt.Printf("Die Kante (%v -> %v) existiert bereits.\n", fromNodeId, toNodeId)
				fromNodeToToNodeExists = false
			}
		}

		for _, vertex := range toNode.adjacencyList {
			if vertex.toNode == fromNode {
				fmt.Printf("Die Kante (%v -> %v) existiert bereits.\n", fromNodeId, toNodeId)
				toNodeToFromNodeExists = false
			}
		}

		if fromNodeToToNodeExists {
			fromNode.adjacencyList = append(fromNode.adjacencyList, &Edge{fromNode: fromNode, toNode: toNode, weight: weight})
		}

		if toNodeToFromNodeExists {
			toNode.adjacencyList = append(toNode.adjacencyList, &Edge{fromNode: toNode, toNode: fromNode, weight: weight})
		}

	} else {
		fmt.Printf("Die Kante (%v -> %v) ist ungültig\n", fromNodeId, toNodeId)
	}

}

func (g Graph) GetVertex(nodeId string) *Vertex {
	for _, v := range g.vertecies {
		if v.nodeId == nodeId {
			return v
		}
	}
	return nil
}

func (g *Graph) AddVertex(nodeId string) {
	for _, vertex := range g.vertecies {
		if vertex.nodeId == nodeId {
			//fmt.Printf("This vertex already exists (%v)\n", nodeId)
			return
		}
	}
	g.vertecies = append(g.vertecies, &Vertex{nodeId: nodeId})
}

func (g *Graph) Dijkstra(fromNodeId, toNodeId string) map[string]float64 {

	resultMap := map[string]float64{}

	if g.DFS(fromNodeId)[toNodeId] == false {
		fmt.Printf("Es gibt keine Verbindung zwischen dem Startknoten (%v) und dem Zielknoten (%v)\n", fromNodeId, toNodeId)
		return nil
	}

	var result []*Vertex
	fromNode := g.GetVertex(fromNodeId)
	fromNode.distance = 0
	data := &MinHeap{}
	data.Insert(fromNode)

	for len(data.Elements) > 0 {
		currentNode2 := data.ExtractMin()
		for _, nextNode := range currentNode2.adjacencyList {
			if nextNode.toNode.previousVertex == nil || nextNode.toNode.distance > (currentNode2.distance+nextNode.weight) {
				nextNode.toNode.previousVertex = currentNode2
				nextNode.toNode.distance = currentNode2.distance + nextNode.weight
				data.Insert(nextNode.toNode)
			}
		}
	}

	for _, node := range g.vertecies {
		if node.nodeId == toNodeId {
			result = append(result, node)
			resultMap[node.nodeId] = node.distance
			break
		}
	}

	if len(result) == 1 {
		for result[len(result)-1].nodeId != fromNodeId {
			result = append(result, result[len(result)-1].previousVertex)
		}
	} else {
		fmt.Printf("The dijkstra algorithm did not found a path.")
		return nil
	}

	// Print result
	if len(result) > 0 {
		fmt.Printf("Der kürzeste Weg (Länge: %v) vom Startknoten (%v) zum Zielknoten (%v) ist:\t[", g.GetVertex(toNodeId).distance, fromNodeId, toNodeId)
		for len(result) > 0 {
			fmt.Printf("%v", result[len(result)-1].nodeId)
			result = result[:len(result)-1]
			if len(result) >= 1 {
				fmt.Printf(", ")
			}
		}
		fmt.Println("]")
	}
	return resultMap
}

func (g *Graph) NumVertices() int {
	return len(g.vertecies)

}

func (g *Graph) NumEdges() int {
	numberOfEdged := 0
	for _, v := range g.vertecies {
		numberOfEdged = numberOfEdged + len(v.adjacencyList)
	}
	return numberOfEdged
}

func (g *Graph) BFS(nodeId string) map[string]int {

	resultMap := map[string]int{}
	data := &Queue{}

	for _, node := range g.vertecies {
		if node.nodeId == nodeId {
			node.explored = true
			node.layer = 0
			resultMap[node.nodeId] = 0
			data.Enqueue(node)
		} else {
			node.explored = false
		}
	}

	for len(data.Elements) > 0 {
		currentNode := data.Peek()
		for _, c := range currentNode.adjacencyList {
			if !c.toNode.explored {
				c.toNode.explored = true
				c.toNode.layer = c.fromNode.layer + 1
				resultMap[c.toNode.nodeId] = c.fromNode.layer + 1
				data.Enqueue(c.toNode)
			}
		}
		data.Dequeue()
	}
	return resultMap
}

func (g *Graph) DFS(nodeId string) map[string]bool {
	resultMap := map[string]bool{}
	data := &Stack{}

	for _, node := range g.vertecies {
		node.explored = false
		resultMap[node.nodeId] = false
	}

	data.Push(g.GetVertex(nodeId))

	for !data.Empty() {
		currentNode := data.Pop()

		if !currentNode.explored {
			currentNode.explored = true
			resultMap[currentNode.nodeId] = true
			for _, node := range currentNode.adjacencyList {
				data.Push(node.toNode)
			}
		}

	}
	return resultMap
}

func (g *Graph) Print() {
	for _, v := range g.vertecies {
		fmt.Printf("\nKnoten %v : ", v.nodeId)
		for _, w := range v.adjacencyList {
			fmt.Printf(" %v (%v)", w.toNode.nodeId, w.weight)
		}
	}
	fmt.Println("\n")
}

func initGraph9(filename string) Graph {

	graph := Graph{}

	file, _ := os.Open(filename) //
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		fields := strings.Fields(s)
		id1 := fields[0]
		graph.AddVertex(id1)
		for _, x := range fields[1:] {
			f := strings.Split(x, ",") // f[0]:id2 ,,
			var length float64         //edgeLength
			if l, err := strconv.ParseFloat(f[1], 64); err == nil {
				length = l //edgeLength{float64: l}
			} else {
				panic("convert str2float failed!")
			}
			graph.AddVertex(f[0])
			graph.AddDirectedWeightedEdge(id1, f[0], length)

		}

	}
	return graph
}

func main() {

	// First Graph
	fmt.Println("Graph:")
	graph := NewGraph()

	// Add vertecies
	graph.AddVertex(strconv.Itoa(0))
	graph.AddVertex(strconv.Itoa(1))
	graph.AddVertex(strconv.Itoa(2))
	graph.AddVertex(strconv.Itoa(3))
	graph.AddVertex(strconv.Itoa(5))
	graph.AddVertex(strconv.Itoa(7))

	// Add edges
	graph.AddUndirectedWeightedEdge("1", "3", 3)
	graph.AddUndirectedWeightedEdge("1", "2", 5)
	graph.AddUndirectedWeightedEdge("1", "7", 1)
	graph.AddUndirectedWeightedEdge("1", "0", 4)
	graph.AddUndirectedWeightedEdge("0", "7", 1)
	graph.AddUndirectedWeightedEdge("2", "3", 2)
	graph.AddUndirectedWeightedEdge("2", "7", 3)
	graph.AddUndirectedWeightedEdge("7", "5", 1)
	graph.AddUndirectedWeightedEdge("3", "5", 7)

	// Print graph
	graph.Print()

	// Graph stats
	fmt.Printf("Anzahl an Knoten: %v\n", graph.NumVertices())
	fmt.Printf("Anzahl an Kanten: %v\n", graph.NumEdges())
	fmt.Printf("BFS (Startknoten: 1): %v\n", graph.BFS("1"))
	fmt.Printf("DFS (Startknoten: 1): %v\n", graph.DFS("1"))
	fmt.Printf("Gerichtet: %v\n", !graph.isUndirected)

	graph.Dijkstra("1", "5")
	graph.Dijkstra("1", "1")
	graph.Dijkstra("1", "7")

	// Second graph
	fmt.Println("\nGraph: problem9.8test.txt")
	graph2 := NewGraph()

	// Init vertecies and edges
	graph2 = initGraph9("src/problem9.8test.txt")

	// Print graph
	graph2.Print()

	// Graph stats
	fmt.Printf("Anzahl an Knoten: %v\n", graph2.NumVertices())
	fmt.Printf("Anzahl an Kanten: %v\n", graph2.NumEdges())
	fmt.Printf("BFS (Startknoten: 1): %v\n", graph2.BFS("1"))
	fmt.Printf("DFS (Startknoten: 1): %v\n", graph2.DFS("1"))
	fmt.Printf("Gerichtet: %v\n", !graph2.isUndirected)

	for i := 1; i <= 8; i++ {
		graph2.Dijkstra("1", strconv.Itoa(i))
	}

	// Third graph
	fmt.Println("\nGraph: problem9.8.txt")
	graph3 := NewGraph()

	// Init vertecies and edges
	graph3 = initGraph9("src/problem9.8.txt")

	// Print graph
	//graph3.Print()

	fmt.Printf("Anzahl an Knoten: %v\n", graph3.NumVertices())
	fmt.Printf("Anzahl an Kanten: %v\n", graph3.NumEdges())
	//fmt.Printf("BFS (Startknoten: 1): %v\n", graph3.BFS("1"))
	//fmt.Printf("DFS (Startknoten: 1): %v\n", graph3.DFS("1"))
	fmt.Printf("Gerichtet: %v\n", !graph3.isUndirected)

	graph3.Dijkstra("1", "7")
	graph3.Dijkstra("1", "37")
	graph3.Dijkstra("1", "59")
	graph3.Dijkstra("1", "82")
	graph3.Dijkstra("1", "99")
	graph3.Dijkstra("1", "115")
	graph3.Dijkstra("1", "133")
	graph3.Dijkstra("1", "165")
	graph3.Dijkstra("1", "188")
	graph3.Dijkstra("1", "197")

	/*
		heap := &MinHeap{}

		heap.Insert(&Vertex{distance: 12})
		heap.Insert(&Vertex{distance: 13})
		heap.Insert(&Vertex{distance: 16})
		heap.Insert(&Vertex{distance: 34})
		heap.Insert(&Vertex{distance: 45})
		heap.Insert(&Vertex{distance: 100})

		heap.PrintHeap()

		heap.ExtractMin()

		heap.PrintHeap()
		heap2 := &MaxHeap{}

		heap2.Insert(&Vertex{distance: 100})
		heap2.Insert(&Vertex{distance: 19})
		heap2.Insert(&Vertex{distance: 36})
		heap2.Insert(&Vertex{distance: 17})
		heap2.Insert(&Vertex{distance: 3})
		heap2.Insert(&Vertex{distance: 25})
		heap2.Insert(&Vertex{distance: 1})
		heap2.Insert(&Vertex{distance: 2})
		heap2.Insert(&Vertex{distance: 7})

		heap2.PrintHeap()

		heap2.ExtractMax()

		heap2.PrintHeap()
	*/

}
