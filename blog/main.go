package main

import (
	"fmt"
	"msgo"
	"net/http"
)

func main() {
	engine := msgo.New()
	engine.Add("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s test", "kris")
	})
	engine.Run()
}
