package main

import (
    "fmt"
    "html/template"
    "net/http"
    "os"
    "path/filepath"
)

type NoteStruct struct {
    Type string
    Path string
}

func discover(path string, notes *[]NoteStruct) {
    filepath.Walk(path, func(path string, info os.FileInfo, err error) error {

        if err != nil {
            fmt.Println(err)
        }

        item := NoteStruct{}

        if info.IsDir() {
            item = NoteStruct{Type: "dir", Path: path}

        } else {
            item = NoteStruct{Type: "file", Path: path}

        }
        *notes = append(*notes, item)

        return nil
    })
}

func indexHandler(template *template.Template) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        template.Execute(w, nil)
    }
}

func pathHndler(notes []NoteStruct) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        t, _ := template.ParseFiles("/home/brun0/workspace/personal-kb/templates/base.html")
        t.Execute(w, notes)
    }
}

func main() {
    notesDir := "/home/brun0/workspace/personal-kb/notes/"

    notes := &[]NoteStruct{}
    discover(notesDir, notes)

    port := "9191"
    fs := http.FileServer(http.Dir("assets"))
    mux := http.NewServeMux()

    tpl := template.Must(template.ParseFiles("/home/brun0/workspace/personal-kb/templates/index.html"))

    mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
    mux.HandleFunc("/", indexHandler(tpl))
    mux.HandleFunc("/path", pathHndler(*notes))
    http.ListenAndServe(":"+port, mux)

}
