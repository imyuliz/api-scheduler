package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/imyuliz/api-scheduler/pkg/middleware"
	"github.com/imyuliz/api-scheduler/version"
)

func main() {
	fmt.Printf("commit: %20s\n", version.GitCommit)
	fmt.Printf("built on %20s\n", version.BuildGoVersion)
	fmt.Printf("built on %20s\n", version.BuildSystem)
	fmt.Println("Hi! I'm yulizz")
	ws := new(restful.WebService)
	c := restful.NewContainer()
	c.Filter(middleware.Log)
	c.Filter(middleware.Handler)
	c.Add(ws)
	server := &http.Server{Addr: ":8080", Handler: c}
	log.Fatal(server.ListenAndServe())

}
