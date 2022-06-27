package main

import (
	"fmt"
	"log"

	"github.com/imyuliz/api-scheduler/frame"
	"github.com/imyuliz/api-scheduler/version"
)

func main() {
	fmt.Printf("commit: %20s\n", version.GitCommit)
	fmt.Printf("built on %20s\n", version.BuildGoVersion)
	fmt.Printf("built on %20s\n", version.BuildSystem)
	fmt.Println("Hi! I'm yulizz")
	server := frame.NewServer()
	server.GET("/richardyu", func(c *frame.Context) {
		c.HTML(200, "<h1>Hello RichardYu</h1>")
	})
	log.Fatal(server.Run(":8080"))

}
