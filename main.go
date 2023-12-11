package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/beego/beego/v2/server/web"
)

func createHome(w http.ResponseWriter, r *http.Request) {
	indexdata, _ := os.ReadFile("./ui/static/index.gohtml")
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, string(indexdata))
}

func main() {

	// // API routes
	// http.HandleFunc("/", createHome)
	// // http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
	// // 	fmt.Fprintf(w, string(styledata))
	// // })
	// // http.HandleFunc("/main.js", func(w http.ResponseWriter, r *http.Request) {
	// // 	fmt.Fprintf(w, string(jsdata))
	// // })

	// fileServer := http.FileServer(http.Dir("./ui/static/"))
	// http.Handle("/static/", http.StripPrefix("/static", fileServer))

	// port := ":8080"
	// fmt.Println("Server is running on port" + port)

	// // Start server on port specified above
	// log.Fatal(http.ListenAndServe(port, nil))
	web.Run()
}
