package main

import (
	"fmt"
	"os"
)

func PanicTest() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

// Environment
func Environment() {

	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()
	/*
		for _, e := range os.Environ() {
			pair := strings.Split(e, "=")
			fmt.Println(pair[0])
		}
	*/
}

func main() {

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	// fmt.Println(arg)

	Environment()
	WritingFiles()
	ReadingFiles()
	FilePaths()
	Directories()
	TemporaryFiles()
	HTTPClients()
	// PanicTest()
}
