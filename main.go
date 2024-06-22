package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/jlukenoff/go-portfolio/blog"
)

func index(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("client/index.html")
	if err != nil {
		log.Fatal(err)
	}

	tpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/blog", blog.RenderTemplate)

	fs := http.FileServer(http.Dir("client"))
	http.Handle("/client/", http.StripPrefix("/client/", fs))

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running on port %s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
