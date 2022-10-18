package main

import (
    "brunoteixeira1996/go-kb/utils"
    "fmt"
    "html/template"
    "net/http"
    "os"
    "strings"
)

// TODO: Start server with more control -> create a function for starting
func main() {

    args := os.Args
    if len(args) > 2 {
        fmt.Println("Please provide only the notes directory (fullpath)")
        os.Exit(0) // FIXME: make this return an error in another function
    }
    notesDir := strings.Join(args[1:], "")
    _, err := os.Stat(notesDir)
    if os.IsNotExist(err) {
        fmt.Println("This directory does not exist")
        os.Exit(0)  // FIXME: make this return an error in another function
    }

    //notesDir := "/home/brun0/workspace/go-kb/notes/" // THIS IS JUST FOR DEBUG

    notes := &[]utils.Storage{}
    utils.DiscoverFilesAndDirs(notesDir, notes)

    kb := &utils.Kb{Title: "KB", Notes: *notes}

    mux := http.NewServeMux()

    baseTemplate := template.Must(template.ParseFiles("/home/brun0/workspace/go-kb/templates/base.html", "/home/brun0/workspace/go-kb/templates/index.html"))

    someTemplate := template.Must(template.ParseFiles("/home/brun0/workspace/go-kb/templates/base.html", "/home/brun0/workspace/go-kb/templates/kb.html"))

    noteTemplate := template.Must(template.ParseFiles("/home/brun0/workspace/go-kb/templates/base.html", "/home/brun0/workspace/go-kb/templates/notes.html"))

    fs := http.FileServer(http.Dir("assets"))
    mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

    fs_images := http.FileServer(http.Dir("images"))
    mux.Handle("/images/", http.StripPrefix("/images/", fs_images))

    mux.HandleFunc("/", utils.IndexHandle(baseTemplate))

    mux.HandleFunc("/kb", utils.KbHandle(notesDir, kb, someTemplate, noteTemplate))

    http.ListenAndServe(":8080", mux)
}
