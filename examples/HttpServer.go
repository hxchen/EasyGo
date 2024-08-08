package main

import (
	"log"
	"net/http"
	"time"
)

func welcome(wr http.ResponseWriter, r *http.Request) {
	wr.Write([]byte("welcome"))
}

// timeMiddleware 记录请求处理时间的中间件
func timeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()
		next.ServeHTTP(wr, r)
		timeElapsed := time.Since(timeStart)
		log.Println("time elapsed:", timeElapsed)
	})
}

type Router struct {
	middlewareChain []middleware
	mux             map[string]http.Handler
}

type middleware func(http.Handler) http.Handler

func NewRouter() *Router {
	return &Router{
		mux: make(map[string]http.Handler),
	}
}

func (r *Router) Use(m middleware) {
	r.middlewareChain = append(r.middlewareChain, m)
}

func (r *Router) Add(route string, h http.Handler) {
	var mergedHandler = h
	for i := len(r.middlewareChain) - 1; i >= 0; i-- {
		mergedHandler = r.middlewareChain[i](mergedHandler)
	}
	r.mux[route] = mergedHandler
}

// ServeHTTP 方法使 Router 结构体符合 http.Handler 接口，用于处理传入的 HTTP 请求。
func (r *Router) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	if handler, ok := r.mux[req.URL.Path]; ok {
		handler.ServeHTTP(wr, req)
		return
	}
	http.NotFound(wr, req)
}

func main() {
	r := NewRouter()
	r.Use(timeMiddleware)
	r.Add("/welcome", http.HandlerFunc(welcome))

	http.Handle("/", r)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
