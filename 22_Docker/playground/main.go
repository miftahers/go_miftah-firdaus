package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Halo udin")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hallo ges, mip disini"))
	})

	http.ListenAndServe(":8080", nil)
}
