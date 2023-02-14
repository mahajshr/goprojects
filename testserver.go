package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func calcTable(w http.ResponseWriter, r *http.Request) {
	for i := 1; i < 11; i++ {
		res := 2 * i
		res2 := strconv.Itoa(res) + "\n"
		fmt.Fprintf(w, res2)
	}
}

func main() {

	http.HandleFunc("/", HandleHome)

	http.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "buy from golang server")
	})

	http.HandleFunc("/table", calcTable)
	http.HandleFunc("/file", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Disposition", "attachment; filename=abc.txt")
	})

	fmt.Println("Server started on localhost:8989")
	http.ListenAndServe(":8989", nil)

}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from golang server, use /table to see table of 2")
}
