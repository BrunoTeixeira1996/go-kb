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

// TODO: Struct that represents a .md file
type NoteStruct struct {
    Type string
    Path string
}

// Struct that represents a webpage
type Page struct {
    Title string
}

// Struct that represents a webpage with buttons
type OptionsContent struct {
    Title string
    Notes []NoteStruct
}

// Handles "/"
func indexHandle(baseTemplate *template.Template) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        page := Page{
            Title: "Index",
        }
        baseTemplate.Execute(w, page)
    }
}

// Handles "/options"
func optionsHandle(options *OptionsContent, someTemplate *template.Template) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "GET":
            someTemplate.Execute(w, options)
        case "POST":
            // TODO: Make this get the same template plus rendering the markdown
            // TODO: Make this generate a 404 not found if nothing is found
            press := r.FormValue("submit")
            fmt.Printf("aqui\n")
            html, err := mdToHtml(press)
            if err != nil {
                fmt.Println(err.Error())
            }
            fmt.Fprintf(w, string(html))
        }
        
    }
}

// Function to get every dir and file recursv
// TODO: Copy this to another file
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

// Function that transforms .md syntax to .html
// TODO: Copy this to another file
func mdToHtml(file string) (string, error) {
    content, err := ioutil.ReadFile(file)

    if err != nil {
        return "", err
    }

    html := markdown.ToHTML(content, nil, nil)

    return string(html), nil

}

// TODO: Start server with more control
func main() {
    notesDir := "/home/brun0/workspace/personal-kb/notes/"

    notes := &[]NoteStruct{}
    discover(notesDir, notes)
    options := &OptionsContent{Title: "something", Notes: *notes}

    mux := http.NewServeMux()

    baseTemplate := template.Must(template.ParseFiles("/home/brun0/workspace/personal-kb/templates/base.html", "/home/brun0/workspace/personal-kb/templates/index.html"))

    someTemplate := template.Must(template.ParseFiles("/home/brun0/workspace/personal-kb/templates/base.html", "/home/brun0/workspace/personal-kb/templates/options.html"))


    fs := http.FileServer(http.Dir("assets"))
    mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

    
    mux.HandleFunc("/", indexHandle(baseTemplate))

    // TODO: Think of a better name than options
    mux.HandleFunc("/options", optionsHandle(options, someTemplate))

    http.ListenAndServe(":8080", mux)
}
