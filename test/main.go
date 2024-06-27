package main

import (
	"fmt"
	"net/http"
)

func main() {

	server := http.Server{
		Addr: ":8081",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(r.URL.String())
			fmt.Println(r.Method)
			w.Write([]byte("hello"))

		}),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
