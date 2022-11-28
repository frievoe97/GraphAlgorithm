package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Graph struct {
	vertecies []*Vertex
}

type Vertex struct {
	nodeId         string
	adjacent       []*Vertex
	edgeAdjacent   []*Edge
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

func (g Graph) AddDirectedWeightedEdge(fromNodeId, toNodeId string, weight float64) {

	fromNode := g.GetVertex(fromNodeId)
	toNode := g.GetVertex(toNodeId)

	if fromNode != nil && toNode != nil {
		for _, vertex := range fromNode.edgeAdjacent {
			if vertex.toNode == toNode {
				err := fmt.Errorf("This already exists (%v -> %v)", fromNodeId, toNodeId)
				fmt.Println(err.Error())
				return
			}
		}
		fromNode.edgeAdjacent = append(fromNode.edgeAdjacent, &Edge{fromNode: fromNode, toNode: toNode, weight: weight})
	} else {
		err := fmt.Errorf("Invalid Edge (%v -> %v)", fromNodeId, toNodeId)
		fmt.Println(err.Error())
	}

}

func (g Graph) AddUndirectedWeightedEdge(fromNodeId, toNodeId string, weight float64) {

	fromNode := g.GetVertex(fromNodeId)
	toNode := g.GetVertex(toNodeId)

	fromNodeToToNodeExists := true
	toNodeToFromNodeExists := true

	if fromNode != nil && toNode != nil {
		for _, vertex := range fromNode.edgeAdjacent {
			if vertex.toNode == toNode {
				fmt.Printf("This edge already exists (%v -> %v)", fromNodeId, toNodeId)
				fromNodeToToNodeExists = false
			}
		}

		for _, vertex := range toNode.edgeAdjacent {
			if vertex.toNode == fromNode {
				fmt.Printf("This edge already exists (%v -> %v)", fromNodeId, toNodeId)
				toNodeToFromNodeExists = false
			}
		}

		if fromNodeToToNodeExists {
			fromNode.edgeAdjacent = append(fromNode.edgeAdjacent, &Edge{fromNode: fromNode, toNode: toNode, weight: weight})
		}

		if toNodeToFromNodeExists {
			toNode.edgeAdjacent = append(toNode.edgeAdjacent, &Edge{fromNode: toNode, toNode: fromNode, weight: weight})
		}

	} else {
		fmt.Printf("Invalid edge (%v -> %v)", fromNodeId, toNodeId)
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

func (g *Graph) Dijkstra(fromNodeId, toNodeId string) {

	resultMap := map[string]float64{}
	var result []*Vertex
	fromNode := g.GetVertex(fromNodeId)
	fromNode.distance = 0
	data := &Queue{}
	data.Enqueue(fromNode)

	for data.GetLength() > 0 {

		currentNode := data.Peek()

		for _, nextNode := range currentNode.edgeAdjacent {

			if nextNode.toNode.previousVertex == nil || nextNode.toNode.distance > (currentNode.distance+nextNode.weight) {
				nextNode.toNode.previousVertex = currentNode
				nextNode.toNode.distance = currentNode.distance + nextNode.weight
				data.Enqueue(nextNode.toNode)
			}
		}
		data.Dequeue()
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
		return
	}

	if len(result) > 0 {
		fmt.Printf("The shortest route (length: %v) from %v to %v is:\t[", g.GetVertex(toNodeId).distance, fromNodeId, toNodeId)
		for len(result) > 0 {
			fmt.Printf("%v", result[len(result)-1].nodeId)
			result = result[:len(result)-1]
			if len(result) >= 1 {
				fmt.Printf(", ")
			}
		}
		fmt.Println("]")
	}
}

func (g *Graph) NumVertices() int {
	return len(g.vertecies)

}

func (g *Graph) NumEdges() int {
	numberOfEdged := 0
	for _, v := range g.vertecies {
		numberOfEdged = numberOfEdged + len(v.edgeAdjacent)
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
		for _, c := range currentNode.edgeAdjacent {
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
			for _, node := range currentNode.edgeAdjacent {
				data.Push(node.toNode)
			}
		}

	}
	return resultMap
}

func (g *Graph) Print() {
	for _, v := range g.vertecies {
		fmt.Printf("\nVertex %v : ", v.nodeId)
		for _, w := range v.edgeAdjacent {
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

	// My own test
	fmt.Println("My own test!")
	graph := &Graph{}

	graph.AddVertex(strconv.Itoa(0))
	graph.AddVertex(strconv.Itoa(1))
	graph.AddVertex(strconv.Itoa(2))
	graph.AddVertex(strconv.Itoa(3))
	graph.AddVertex(strconv.Itoa(5))
	graph.AddVertex(strconv.Itoa(7))

	graph.AddUndirectedWeightedEdge("1", "3", 3)
	graph.AddUndirectedWeightedEdge("1", "2", 5)
	graph.AddUndirectedWeightedEdge("1", "7", 1)
	graph.AddUndirectedWeightedEdge("1", "0", 4)
	graph.AddUndirectedWeightedEdge("0", "7", 1)
	graph.AddUndirectedWeightedEdge("2", "3", 2)
	graph.AddUndirectedWeightedEdge("2", "7", 3)
	graph.AddUndirectedWeightedEdge("7", "5", 1)
	graph.AddUndirectedWeightedEdge("3", "5", 7)

	fmt.Println("Graph:")
	graph.Print()

	fmt.Printf("Number of Vertecies: %v\n", graph.NumVertices())
	fmt.Printf("Number of Edges: %v\n", graph.NumEdges())
	fmt.Printf("BFS: %v\n", graph.BFS("1"))
	fmt.Printf("DFS: %v\n", graph.DFS("1"))

	graph.Dijkstra("1", "5")
	graph.Dijkstra("1", "1")
	graph.Dijkstra("1", "7")

	// problem9.8test.txt
	fmt.Println("\n\nproblem9.8test.txt")
	graph2 := Graph{}
	graph2 = initGraph9("src/problem9.8test.txt")

	fmt.Printf("Number of Vertecies: %v\n", graph2.NumVertices())
	fmt.Printf("Number of Edges: %v\n", graph2.NumEdges())

	for i := 1; i <= 8; i++ {
		graph2.Dijkstra("1", strconv.Itoa(i))
	}

	// problem9.8.txt
	fmt.Println("\n\nproblem9.8.txt")
	graph3 := Graph{}
	graph3 = initGraph9("src/problem9.8.txt")
	//graph3.Print()

	fmt.Printf("Number of Vertecies: %v\n", graph3.NumVertices())
	fmt.Printf("Number of Edges: %v\n", graph3.NumEdges())

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
}
