package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	slice := make([][]uint8, dy)

	for i := range slice {
		slice[i] = make([]uint8, dx)
		for j := range slice[i] {
			slice[i][j] = uint8(i ^ j)
		}
	}

	return slice
}

func main() {
	pic.Show(Pic)
}

// before exercise

// import (
// 	"fmt"
// 	"strings"
// )

// type Vertex struct {
// 	X int
// 	Y int
// }

// func main() {
// 	whatIsARange()
// 	nilSlices()
// 	slicesAreReferences()
// 	sliceThemUp()
// 	arrayMethods()
// 	structLiterals()
// 	structAndPointers()
// 	structMethod()
// 	pointers()
// }

// func whatIsARange() {
// 	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

// 	for i, v := range pow {
// 		fmt.Printf("2**%d = %d\n", i, v)
// 	}
// }

// func nilSlices() {
// 	var nilSlice []int
// 	printSlice(nilSlice)
// 	if nilSlice == nil {
// 		fmt.Println("nil!")
// 		nilSlice = append(nilSlice, 1)
// 	}
// 	printSlice(nilSlice)
// }

// func printSlice(s []int) {
// 	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
// }
// func slicesAreReferences() {
// 	childrenOfThanos := [6]string{
// 		"Ebony Maw",
// 		"Proxima Midnight",
// 		"Corvus Glaive",
// 		"Cull Obsidian",
// 		"Gamora",
// 		"Nebula",
// 	}

// 	// a := childrenOfThanos[0:4]
// 	b := childrenOfThanos[4:]
// 	b[0], b[1] = b[1], "Gamora (deceased)"
// 	// fmt.Println(a, b)
// 	// greetAll(childrenOfThanos[:])

// 	s := childrenOfThanos[:0]
// 	greetAll(s)

// 	s = s[:4]
// 	greetAll(s)

// 	fmt.Println("Printing s[:4]")
// 	s = s[4:] // this wont work. why?
// 	greetAll(s)
// }

// func greetAll(people []string) {
// 	for _, name := range people {
// 		if strings.Contains(name, "deceased") {
// 			fmt.Println("RIP", name)
// 		} else {

// 			fmt.Println("Hello,", name)
// 		}
// 	}
// }

// func sliceThemUp() {
// 	primes := [6]int{2, 3, 5, 7, 11, 13}

// 	var firstSlice []int = primes[2:4]
// 	fmt.Println(firstSlice)
// }
// func arrayMethods() {
// 	var a [2]string
// 	a[1] = "World"
// 	a[0] = "Hello"

// 	fmt.Println(a[0], a[1])
// 	fmt.Println(a)

// 	primes := [6]int{2, 3, 5, 7, 11, 13}
// 	fmt.Println(primes)
// }

// func structLiterals() {
// 	var (
// 		v1 = Vertex{1, 2}  // has type Vertex
// 		v2 = Vertex{X: 1}  // Y:0 is implicit
// 		v3 = Vertex{}      // X:0 and Y:0
// 		p  = &Vertex{1, 2} // has type *Vertex
// 	)

// 	fmt.Println(v1, p, v2, v3)
// }

// func structAndPointers() {
// 	v := Vertex{1, 2}
// 	p := &v
// 	p.X = 1.1e9
// 	fmt.Println(v)
// }

// func structMethod() {
// 	fmt.Println(Vertex{1, 2})

// 	v := Vertex{1, 2}
// 	v.X = 4
// 	fmt.Println(v.X, v.Y)
// }

// func pointers() {
// 	var p *int
// 	i := 42
// 	p = &i
// 	fmt.Println(*p)
// 	*p = *p / 2
// 	fmt.Println(i)
// }
