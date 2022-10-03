package main

import (
    "fmt"
)

func main() {
    fmt.Println("test")
}

// ========================================================================================= //

// WEBSERVER STUFF

// var tpl = template.Must(template.ParseFiles("/home/brun0/workspace/personal-kb/templates/index.html"))

// func indexHandler(w http.ResponseWriter, r *http.Request) {
//     tpl.Execute(w, nil)
// }

// func main() {

//     port := "9191"

//     fs := http.FileServer(http.Dir("assets"))

//     mux := http.NewServeMux()

//     mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
//     mux.HandleFunc("/", indexHandler)
//     http.ListenAndServe(":"+port, mux)
// }

//https://freshman.tech/web-development-with-go/
//https://github.com/zorchenhimer/MovieNight/blob/master/main.go
