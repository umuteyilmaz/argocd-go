package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Go Docker Tutorial")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Docker from Jenkins on ArgoCD... ")
		fmt.Fprintf(w, "Hello K8s... ")
		fmt.Fprintf(w, "Hello Container... ")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
