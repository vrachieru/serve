package main

import (
    "flag"
    "fmt"
    "log"
    "net"
    "net/http"
    "path/filepath"
)

func loggingHandler(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(respWriter http.ResponseWriter, req *http.Request) {
        host, _, err := net.SplitHostPort(req.RemoteAddr)
        if err != nil {
            host = req.RemoteAddr
        }

        log.Printf("%s - %s %s", host, req.Method, req.URL.Path)
        handler.ServeHTTP(respWriter, req)
    })
}

func main() {
    port := flag.String("p", "8080", "Port on which to serve")
    dir := flag.String("d", ".", "Directory to serve")
    flag.Parse()

    abs, err := filepath.Abs(*dir)
    if err == nil {
        *dir = abs
    }

    fmt.Printf("Serving %s on port %s\n", *dir, *port)
    http.ListenAndServe(":" + *port, loggingHandler(http.FileServer(http.Dir(*dir))))
}
