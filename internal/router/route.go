package router

import ("net/http")

type Route struct {
    Path    string
    Method  string
    HandlerFunction http.HandlerFunc
}