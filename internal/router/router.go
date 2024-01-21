package router

import (
	"net/http"
)

type Router struct {
	routes map[string]Route
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]Route),
	}
}

func (rtr *Router) PathHandler(path string, handlerFunc http.HandlerFunc) *Route {
    var route = Route{
        Path: path,
        HandlerFunction: handlerFunc,
    }

	rtr.routes[path] = route
    return &route
}

func (route *Route) Methods(method string) *Route {
    route.Method = method
    return route
}

func (rtr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    currentRoute := rtr.routes[r.URL.Path]

    if (currentRoute.Path != "") {
        currentRoute.HandlerFunction.ServeHTTP(w, r)
        return
    }

    http.NotFound(w, r)
}