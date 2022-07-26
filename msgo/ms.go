package msgo

import (
	"log"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type Router struct {
	HandlerFuncMap map[string]HandlerFunc
}

func (r *Router) Add(name string, handlerFunc HandlerFunc) {
	r.HandlerFuncMap[name] = handlerFunc
}

type Engine struct {
	Router
}

func New() *Engine {
	return &Engine{
		Router{HandlerFuncMap: make(map[string]HandlerFunc)},
	}
}

func (e *Engine) Run() {
	for key, value := range e.HandlerFuncMap {
		http.HandleFunc(key, value)
	}
	err := http.ListenAndServe(":8881", nil)
	if err != nil {
		log.Fatal(err)
	}
}
