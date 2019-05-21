package main

import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	res.Header().Set("Content-Type", "text/html")
	io.WriteString(res,
		`<doctype html>
<html>
	<head>
		<meta charset="utf-8">
		<title>Hello `+name+`</title>
	</head>
	<body>
		Hello `+name+`!
	</body>
</html>`)
}

func main() {
	fs := http.FileServer(http.Dir("lesson6/static"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":80", nil)
}
