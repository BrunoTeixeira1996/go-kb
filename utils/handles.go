package utils

import (
    "html/template"
    "net/http"
    "strings"
)

// Struct that represents a webpage
type Page struct {
    Title string
}

// Struct that represents a webpage with buttons
type Kb struct {
    Title string
    Notes []Storage
}

// Struct that represents a Note
type Note struct {
    Title   string
    Content template.HTML
}

// Handles "/"
func IndexHandle(baseTemplate *template.Template) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        page := Page{
            Title: "Index",
        }
        baseTemplate.Execute(w, page)
    }
}

// Handles "/options"
// TODO: validate .Execute errors
func KbHandle(rootDir string, kb *Kb, someTemplate *template.Template, notesTemplate *template.Template) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "GET":
            someTemplate.Execute(w, kb)
        case "POST":
            submit := r.FormValue("submit")
            back := r.FormValue("back")

            // Back button was pressed
            if len(back) > 0 {
                temp := strings.Split(back, "/")
                path := strings.Join(temp[:len(temp)-1], "/")

                // Discover files and dirs from new path
                notes := &[]Storage{}
                DiscoverFilesAndDirs(path, notes)

                // Creates new page
                page := &Kb{Notes: *notes}
                // If it's the root dir, change title to KB
                if path == rootDir {
                    page.Title = "KB"
                } else {
                    page.Title = path
                }

                someTemplate.Execute(w, page)

                // Pressed another button
            } else {
                html, err := MdToHtml(submit)

                // This is a dir
                if err != nil {
                    notes := &[]Storage{}
                    DiscoverFilesAndDirs(submit, notes)

                    options := &Kb{Title: submit, Notes: *notes}
                    someTemplate.Execute(w, options)

                    // This is a file
                } else {
                    note := Note{Title: submit, Content: template.HTML(html)}
                    notesTemplate.Execute(w, note)
                }

            }

        }

    }
}
