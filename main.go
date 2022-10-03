package main

import (
    "fmt"
    "os"
    "path/filepath"
)

type NoteStruct struct {
    Type string
    Path string
}

func discover(path string) {
    filepath.Walk(path, func(path string, info os.FileInfo, err error) error {

        if err != nil {
            fmt.Println(err)
        }

        notes := []NoteStruct{}
        item := NoteStruct{}

        if info.IsDir() {
            item = NoteStruct{Type:"dir", Path: path}

        } else {
            item = NoteStruct{Type:"file", Path: path}

        }
        notes = append(notes, item)


        for _, j := range notes {
            fmt.Printf("Type: %s Path: %s\n", j.Type, j.Path)
        }
        
        return nil
    })
}

func main() {
    notesDir := "/home/brun0/workspace/personal-kb/notes/"

    discover(notesDir)
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
