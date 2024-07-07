package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/pic"
)

func main() {
	c1 := counter(10)
	c2 := counter(20)

	fmt.Println(c1()) // 11
	fmt.Println(c1()) // 12
	fmt.Println(c2()) // 21
	fmt.Println(c1()) // 13

	sumOfTwoNumber := func(number1, number2 int) int {
		return number1 + number2
	}

	fmt.Println(sumOfTwoNumber(1, 2))

	fmt.Println(exerciseSolutionMaps("I am learning Go"))

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	mapMutation()
	mappingFirstLook()
	pic.Show(exerciseSolutionPic)
	whatIsARange()
	nilSlices()
	slicesAreReferences()
	sliceThemUp()
	arrayMethods()
	structLiterals()
	structAndPointers()
	structMethod()
	pointers()
}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func counter(start int) func() int {
	count := start
	return func() int {
		count++
		return count
	}
}
func exerciseSolutionMaps(s string) map[string]int {
	result := make(map[string]int)

	arrayOfStrings := strings.Split(s, " ")

	for _, value := range arrayOfStrings {
		_, isPresent := result[value]

		if isPresent {
			result[value]++
		} else {
			result[value] = 1
		}

	}

	return result
}

type Vertex struct {
	Lat  float64
	Long float64
}

func mapMutation() {
	m := make(map[string]string)

	m["Answer"] = "42"
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = "48 is a number"
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func mappingFirstLook() {
	m := make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	m = map[string]Vertex{
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}
	fmt.Println(m["Bell Labs"])

	fmt.Println(m)
}

func exerciseSolutionPic(dx, dy int) [][]uint8 {
	slice := make([][]uint8, dy)

	for i := range slice {
		slice[i] = make([]uint8, dx)
		for j := range slice[i] {
			slice[i][j] = uint8(i ^ j)
		}
	}

	return slice
}

func whatIsARange() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func nilSlices() {
	var nilSlice []int
	printSlice(nilSlice)
	if nilSlice == nil {
		fmt.Println("nil!")
		nilSlice = append(nilSlice, 1)
	}
	printSlice(nilSlice)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
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

	// a := childrenOfThanos[0:4]
	b := childrenOfThanos[4:]
	b[0], b[1] = b[1], "Gamora (deceased)"
	// fmt.Println(a, b)
	// greetAll(childrenOfThanos[:])

	s := childrenOfThanos[:0]
	greetAll(s)

	s = s[:4]
	greetAll(s)

	fmt.Println("Printing s[:4]")
	s = s[4:] // this wont work. why?
	greetAll(s)
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
		v1 = Vertex{1, 2}
		v2 = Vertex{Lat: 1}
		v3 = Vertex{}
		p  = &Vertex{1, 2}
	)

	fmt.Println(v1, p, v2, v3)
}

func structAndPointers() {
	v := Vertex{1, 2}
	p := &v
	p.Lat = 1.1e9
	fmt.Println(v)
}

func structMethod() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.Lat = 4
	fmt.Println(v.Lat, v.Long)
}

func pointers() {
	var p *int
	i := 42
	p = &i
	fmt.Println(*p)
	*p = *p / 2
	fmt.Println(i)
}
