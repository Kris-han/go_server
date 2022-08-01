package msgo

import (
	"fmt"
	"log"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type RouterGroup struct {
	name             string
	HandlerFuncMap   map[string]HandlerFunc
	HandlerMethodMap map[string][]string
}

//func (r *RouterGroup) Add(name string, handlerFunc HandlerFunc) {
//	r.HandlerFuncMap[name] = handlerFunc
//}

func (r *RouterGroup) Any(name string, handlerFunc HandlerFunc) {
	r.HandlerFuncMap[name] = handlerFunc
	r.HandlerMethodMap["ANY"] = append(r.HandlerMethodMap["ANY"], name)
}
func (r *RouterGroup) Get(name string, handlerFunc HandlerFunc) {
	r.HandlerFuncMap[name] = handlerFunc
	r.HandlerMethodMap[http.MethodGet] = append(r.HandlerMethodMap[http.MethodGet], name)
}
func (r *RouterGroup) Post(name string, handlerFunc HandlerFunc) {
	r.HandlerFuncMap[name] = handlerFunc
	r.HandlerMethodMap[http.MethodPost] = append(r.HandlerMethodMap[http.MethodPost], name)
}

type router struct {
	RouterGroups []*RouterGroup
}

func (r *router) Group(name string) *RouterGroup {
	RouterGroup := &RouterGroup{
		name:             name,
		HandlerFuncMap:   make(map[string]HandlerFunc),
		HandlerMethodMap: make(map[string][]string),
	}
	r.RouterGroups = append(r.RouterGroups, RouterGroup)
	return RouterGroup
}

type Engine struct {
	router
}

func New() *Engine {
	return &Engine{
		router: router{},
	}
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	for _, group := range e.RouterGroups {
		for name, MethodHandle := range group.HandlerFuncMap {
			url := "/" + group.name + name
			if r.RequestURI == url {
				routers, ok := group.HandlerMethodMap["ANY"]
				if ok {
					for _, routerName := range routers {
						if routerName == name {
							MethodHandle(w, r)
							return
						}
					}
				}
				// 按照method进行匹配
				routers, ok = group.HandlerMethodMap[method]
				if ok {
					for _, routerName := range routers {
						if routerName == name {
							MethodHandle(w, r)
							return
						}
					}
				}
				w.WriteHeader(http.StatusMethodNotAllowed)
				fmt.Fprint(w, r.RequestURI+method, "not allowed  ")
				return
			}
		}
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, r.RequestURI+method, "not found  ")
}

func (e *Engine) Run() {
	http.Handle("/", e)

	err := http.ListenAndServe(":8881", nil)
	if err != nil {
		log.Fatal(err)
	}
}
