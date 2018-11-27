package main

import (
	"encoding/json"
	"fmt"
	"index/suffixarray"
)


type Point struct {
	X, Y int
}

type Circle struct {
	Center Point
	Radius int
}

type Wheel struct {
	Circle Circle
	Spokes int
}

func foo() {
	var w Wheel
	w.Circle.Center.X = 8
	w.Circle.Center.Y = 8
	w.Circle.Radius = 5
	w.Spokes = 20

	fmt.Println("foo(): ", w)
}

type Circle2 struct {
	Point  // anonymous fields with a named type
	Radius int
}

type Wheel2 struct {
	Circle2
	Spokes int
}

func bar() {
	var w Wheel2
	w.X = 18
	w.Y = 18
	w.Radius = 15
	w.Spokes = 120

	fmt.Println("bar(): ", w)
}

type Wheel3 struct {
	*Point
	Radius int
}

func baz() {
	var w Wheel3
	w = Wheel3{&Point{28, 28}, 25}

	json_string, err := json.Marshal(w)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("baz(): %s\n", json_string)
	}

	fmt.Printf("baz(): %#v\n", w)

}

func main() {
	arr:= []int{2,3,4,5,6}
	input :=[]int{2,4,5}
	fmt.Println(arr,input)
	index := suffixarray.New(arr)
	index.Lookup(2, -1)
	//bb(arr,input)
	//delete(arr,1)
	//arr=append(arr, 1)
	////arr[2]=0
	//fmt.Println(arr)
}

func bb(dist []int,input []int)  {
	//fmt.Println(dist)
	//var arr []int
	//for k,v := range dist {
	//	fmt.Println(k,v)
	//	a,_:=global.Contain(input,v)
	//	fmt.Println(a)
	//	if !a {
	//		arr=append(arr, v)
	//	}
	//}
	//return arr
	//fmt.Println(arr)
}