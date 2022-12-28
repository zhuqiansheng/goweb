package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func fooHandler(writer http.ResponseWriter, request *http.Request)  {
	fmt.Fprintf(writer,"foo,%q",html.EscapeString(request.URL.Path))
}

func main()  {
	http.Handle("/foo",http.HandlerFunc(fooHandler))
	http.HandleFunc("/bar", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer,"Hello,%q",html.EscapeString(request.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
