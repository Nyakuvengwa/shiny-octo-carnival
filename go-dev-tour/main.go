package main

import (
	"fmt"
	"strings"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	slicesAreReferences()
	sliceThemUp()
	arrayMethods()
	structLiterals()
	structAndPointers()
	structMethod()
	pointers()
}

func slicesAreReferences() {
	childrenOfThanos := [6]string{
		"Ebony Maw",
		"Proxima Midnight",
		"Corvus Glaive",
		"Cull Obsidian",
		"Gamora",
		"Nebula",
	}

	a := childrenOfThanos[0:4]
	b := childrenOfThanos[4:]
	b[0], b[1] = b[1], "Gamora (deceased)"
	fmt.Println(a, b)
	greetAll(childrenOfThanos[:])
}

func greetAll(people []string) {
	for _, name := range people {
		if strings.Contains(name, "deceased") {
			fmt.Println("RIP", name)
		} else {

			fmt.Println("Hello,", name)
		}
	}
}

func sliceThemUp() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var firstSlice []int = primes[2:4]
	fmt.Println(firstSlice)
}
func arrayMethods() {
	var a [2]string
	a[1] = "World"
	a[0] = "Hello"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func structLiterals() {
	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		p  = &Vertex{1, 2} // has type *Vertex
	)

	fmt.Println(v1, p, v2, v3)
}

func structAndPointers() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1.1e9
	fmt.Println(v)
}

func structMethod() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X, v.Y)
}

func pointers() {
	var p *int
	i := 42
	p = &i
	fmt.Println(*p)
	*p = *p / 2
	fmt.Println(i)
}
