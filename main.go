package main

import (
    "brunoteixeira1996/go-kb/utils"
    "html/template"
    "net/http"
)

// TODO: Start server with more control -> create a function for starting
func main() {
    notesDir := "/home/brun0/workspace/go-kb/notes/"

    notes := &[]utils.Storage{}
    utils.DiscoverFilesAndDirs(notesDir, notes)

    kb := &utils.Kb{Title: "KB", Notes: *notes}

    mux := http.NewServeMux()

    baseTemplate := template.Must(template.ParseFiles("/home/brun0/workspace/go-kb/templates/base.html", "/home/brun0/workspace/go-kb/templates/index.html"))

    someTemplate := template.Must(template.ParseFiles("/home/brun0/workspace/go-kb/templates/base.html", "/home/brun0/workspace/go-kb/templates/kb.html"))

    noteTemplate := template.Must(template.ParseFiles("/home/brun0/workspace/go-kb/templates/base.html", "/home/brun0/workspace/go-kb/templates/notes.html"))

    fs := http.FileServer(http.Dir("assets"))
    mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

    mux.HandleFunc("/", utils.IndexHandle(baseTemplate))

    mux.HandleFunc("/kb", utils.KbHandle(notesDir, kb, someTemplate, noteTemplate))

    http.ListenAndServe(":8080", mux)
}

// TODO: make back button work on markdown files
// TODO: fix bug of double back slashes
// TODO: make back button better looking
