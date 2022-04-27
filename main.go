package main

import (
	"fmt"

	"github.com/imyuliz/template-go/version"
)

func main() {
	fmt.Printf("commit: %20s\n", version.GitCommit)
	fmt.Printf("built on %20s\n", version.BuildGoVersion)
	fmt.Printf("built on %20s\n", version.BuildSystem)
	fmt.Println("Hi! I'm yulizz")

}

// HelloWorld define hello world
func HelloWorld() {
	fmt.Println("hello world")
}
