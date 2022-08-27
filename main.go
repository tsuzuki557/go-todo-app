package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	err := http.ListenAndServe(
		":18080",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "aaa")
		}),
	)
	if err != nil {
		fmt.Printf("エラー")
		os.Exit(1)
	}
}
