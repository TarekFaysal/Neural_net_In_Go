package main

import (
	"fmt"
)

func main() {

	var x float64
	var y float64
	fmt.Println("Enter x: ")
	fmt.Scanln(&x)
	fmt.Println("Enter y: ")
	fmt.Scanln(&y)
	var noOfWeights int64
	fmt.Println("Enter No. of Weights: ")
	fmt.Scanln(&noOfWeights)
	var lr float64 = 0.0001
	var weights [100]float64
	var i int64 = 0
	for ; i < noOfWeights; i++ {
		fmt.Println("Enter weight")
		fmt.Scanln(&weights[i])
	}
	// Create the graph
	g := &Graph{}
	var count int64 = 1
	for ; count < 20; count++ {
		n1 := Node{x, 0, 0, nil, nil}
		n1.multiply(weights[0])
		g.Append(&n1)
		n2 := Node{x, 0, 0, nil, nil}
		n2.multiply(weights[1])
		g.Append(&n2)
		n3 := Node{n2.i, 0, 0, nil, nil}
		n3.max()
		g.Append(&n3)
		n4 := Node{n3.i, 0, 0, nil, nil}
		n4.plus(n1.i)
		g.Append(&n4)
		n5 := Node{n4.i, 0, 0, nil, nil}
		n5.minus(y)
		g.Append(&n5)
		n6 := Node{n5.i, 0, 0, nil, nil}
		n6.square(n5.i)
		g.Append(&n6)

		// Print the values of every nodes
		fmt.Println(count, "th Forward Pass")
		var z int64 = 1
		for g.start != nil {
			fmt.Println("Node i", z, "=", g.start.i)
			g.start = g.start.next
			z++
		}

		// <<<calculate derivative of loss by weights>>>

		//calculate derivative of loss by last weight
		var derweights [100]float64
		var memoder [100]float64
		for k := 0; k < 100; k++ {
			derweights[k] = 1
		}
		for g.end.deriw == 0 {
			derweights[0] = derweights[0] * g.end.deri
			g.end = g.end.previous
		}
		memoder[0] = derweights[0]
		derweights[0] = derweights[0] * g.end.deriw
		//Calculate remaining derivative of loss by weights
		var l int64
		for l = 1; l < noOfWeights; l++ {
			if memoder[l-1] != 0 {
				derweights[l] = memoder[l-1]
				g.end = g.end.previous
			}
			for g.end.deriw == 0 {
				derweights[l] = derweights[l] * g.end.deri
				g.end = g.end.previous
			}
			memoder[l] = derweights[l]
			derweights[l] = derweights[l] * g.end.deriw
		}

		fmt.Println(count, "th Updated weights")
		var m int64 = noOfWeights - 1
		var n int64 = 1
		//var w float64
		for p := 0; m >= 0; m-- {
			//w = weights[p]
			weights[p] = weights[p] - (lr * derweights[m])
			fmt.Println("bias", n, "=", weights[m])
			p++
			n++
		}

		g.Delete()
		//fmt.Println("w ", w, "weight0", weights[0])

		//if (w-weights[0])*(w-weights[0]) > 0.0009 {
		//break
		//}

	}
	fmt.Println(count)
}

type Node struct {
	i     float64
	deri  float64
	deriw float64

	next     *Node // link to the next node
	previous *Node // link to the Previous node
}

func (n *Node) plus(weight float64) {
	n.i = n.i + weight
	n.deri = 1
	n.deriw = 1
}
func (n *Node) minus(y float64) {
	n.i = n.i - y
	n.deri = 1
	n.deriw = 0
}
func (n *Node) multiply(weight float64) {
	n.i = n.i * weight
	n.deri = weight
	n.deriw = n.i
}
func (n *Node) square(previousi float64) {
	n.i = n.i * n.i
	n.deri = 2 * previousi
}
func (n *Node) max() {

	if n.i > 0 {
		n.i = n.i
		n.deri = 1
	} else {
		n.i = 0
		n.deri = 0
	}

}

type Graph struct {
	length int
	start  *Node
	end    *Node
}

func (g *Graph) Append(newNode *Node) {
	if g.length == 0 {
		g.start = newNode
		g.end = newNode
	} else {
		lastNode := g.end
		lastNode.next = newNode
		newNode.previous = g.end
		g.end = newNode
	}
	g.length++
}
func (g *Graph) Delete() {
	g.start = nil
	g.end = nil
	g.length = 0
}
