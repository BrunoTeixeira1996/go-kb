package main

import (
    "brunoteixeira1996/go-kb/utils"
    "fmt"
    "html/template"
    "net/http"
    "os"
    "strings"
)

// Function that starts the web server
func startServer(notesDir string, currentPath string) error{
        
    notes := &[]utils.Storage{}
    err := utils.DiscoverFilesAndDirs(notesDir, notes)

    if err != nil {
        return err
    }

    kb := &utils.Kb{Title: "KB", Notes: *notes}

    mux := http.NewServeMux()


    baseTemplate := template.Must(template.ParseFiles(currentPath + "/templates/base.html", currentPath + "/templates/index.html"))

    someTemplate := template.Must(template.ParseFiles(currentPath + "/templates/base.html", currentPath + "/templates/kb.html"))

    noteTemplate := template.Must(template.ParseFiles(currentPath + "/templates/base.html", currentPath + "/templates/notes.html"))

    fs := http.FileServer(http.Dir("assets"))
    mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

    fs_images := http.FileServer(http.Dir("images"))
    mux.Handle("/images/", http.StripPrefix("/images/", fs_images))

    mux.HandleFunc("/", utils.IndexHandle(baseTemplate))

    mux.HandleFunc("/kb", utils.KbHandle(notesDir, kb, someTemplate, noteTemplate))

    http.ListenAndServe(":8080", mux)

    return nil
}

// Function that handles the errors
func run() error {
    args := os.Args
    
    if len(args) > 2 {
        return fmt.Errorf("[ERROR] Please provide only the notes directory (fullpath)\n")
    }

    notesDir := strings.Join(args[1:],"")

    _, err := os.Stat(notesDir)
    if os.IsNotExist(err) {
        return fmt.Errorf("[ERROR] This directory does not exist\n")
    }

    currentPath, err := os.Getwd()
    if err != nil {
        return fmt.Errorf("[ERROR] Couldn't get the current path\n") 
    }
    
   err = startServer(notesDir, currentPath)
    if err != nil {
        return err
    }

    return nil
}


func main() {
    if err := run(); err != nil {
        fmt.Fprint(os.Stderr, err)
        os.Exit(1)
    }
}
