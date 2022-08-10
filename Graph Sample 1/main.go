package main

import "fmt"

type Graph struct{
	vertices []*vertex
}
type vertex struct{
	key int
	adjacent []*vertex
}
//add vertex
func (g *Graph) addVertex(k int){
	if contains(g.vertices,k){
		fmt.Printf("vertex %v is an existing key",k)
		fmt.Println()
	}else{
	  g.vertices=append(g.vertices, &vertex{key: k})
	}
}
//add edge
func(g *Graph)addEdge(from,to int){
	fromVertex:=g.getVertex(from)
	toVertex:=g.getVertex(to)

	if fromVertex==nil||toVertex==nil {
		fmt.Printf("Invalid edge(%v to %v)",from,to)
		fmt.Println()
	}else if contains(fromVertex.adjacent,to) {
		fmt.Printf("existing edge(%v to %v)",from,to)
		fmt.Println()
	}else{
		fromVertex.adjacent=append(fromVertex.adjacent, toVertex)
	}
}

func (g *Graph)getVertex(k int)*vertex{
	for i,v:=range g.vertices{
		if v.key==k{
			return g.vertices[i]
		}
	}
	return nil
}

//contains
func contains(s []*vertex,k int)bool{
	for _,v:=range s{
		if k==v.key {
			return true
		}
	}
	return false
}


func (g *Graph)Print(){
	for _,v:=range g.vertices{
		fmt.Printf("\n vertex %v :",v.key)
		for _,v :=range v.adjacent{
			fmt.Printf(" %v ",v.key)
		}
	}
	fmt.Println()
}



func main(){
	test:=&Graph{}
	for i := 0; i < 5; i++ {
		test.addVertex(i)
	}
	test.addVertex(1)
	test.addVertex(6)
	test.addEdge(1,4)
	test.addEdge(1,4)
	test.addEdge(7,3)
	test.Print()
}