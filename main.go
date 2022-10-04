package main

import (
    "fmt"
    "html/template"
    "io/ioutil"
    "net/http"
    "os"
    "path/filepath"

    "github.com/gomarkdown/markdown"
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

func pathHandler(notes []NoteStruct) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "GET":
            t, _ := template.ParseFiles("/home/brun0/workspace/personal-kb/templates/base.html")
            t.Execute(w, notes)
        case "POST":
            press := r.FormValue("submit")
            html, err := mdToHtml(press)
            if err != nil {
                fmt.Println(err.Error())
            }
            fmt.Fprintf(w, string(html))
        }
    }
}

func mdToHtml(file string) (string, error) {
    content, err := ioutil.ReadFile(file)

    if err != nil {
        return "", err
    }

    html := markdown.ToHTML(content, nil, nil)

    return string(html), nil

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
    mux.HandleFunc("/notes", pathHandler(*notes)) // working on this know
    http.ListenAndServe(":"+port, mux)

}

func main1() {
    file := "/home/brun0/workspace/personal-kb/notes/test1.md"
    content, err := ioutil.ReadFile(file)

    if err != nil {
        fmt.Println(err)
    }

    html := markdown.ToHTML(content, nil, nil)
    fmt.Printf(string(html))
}
