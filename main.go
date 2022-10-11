package main

import (
    "brunoteixeira1996/personal-kb/utils"
    "html/template"
    "net/http"
)

// TODO: Start server with more control -> create a function for starting
func main() {
    notesDir := "/home/brun0/workspace/personal-kb/notes/"

    notes := &[]utils.Storage{}
    utils.Discover(notesDir, notes)
    options := &utils.Kb{Title: "KB", Notes: *notes}

    mux := http.NewServeMux()

    baseTemplate := template.Must(template.ParseFiles("/home/brun0/workspace/personal-kb/templates/base.html", "/home/brun0/workspace/personal-kb/templates/index.html"))

    someTemplate := template.Must(template.ParseFiles("/home/brun0/workspace/personal-kb/templates/base.html", "/home/brun0/workspace/personal-kb/templates/kb.html"))

    noteTemplate := template.Must(template.ParseFiles("/home/brun0/workspace/personal-kb/templates/base.html", "/home/brun0/workspace/personal-kb/templates/notes.html"))

    fs := http.FileServer(http.Dir("assets"))
    mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

    mux.HandleFunc("/", utils.IndexHandle(baseTemplate))

    mux.HandleFunc("/kb", utils.KbHandle(options, someTemplate, noteTemplate))

    http.ListenAndServe(":8080", mux)
}

// TODO: Render folders and files if they are in the current folder
// -> Start the server and render the notes folder and pass that to /kb
// -> Inside kb it will have only the folder and files inside notes folder
// -> If I press a .md I will render that file
// -> If I press a folder I will render the contents of the file
