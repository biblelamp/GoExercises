package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
        //fmt.Fprintf(w, "Hello World!")
        for name, headers := range req.Header {
            for _, h := range headers {
                fmt.Fprintf(w, "%v: %v\n", name, h)
            }
        }
    })

    http.ListenAndServe(":80", nil)
}