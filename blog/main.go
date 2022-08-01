package main

import (
	"fmt"
	"msgo"
	"net/http"
)

func main() {
	engine := msgo.New()
	g := engine.Group("user")
	g.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s test", "kris")
	})
	g.Post("/info", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s info", "htb")
	})
	g.Any("/any", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s any", "test")
	})
	engine.Run()
}
