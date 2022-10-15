package utils

import (
    "html/template"
    "net/http"
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
func KbHandle(options *Kb, someTemplate *template.Template, notesTemplate *template.Template) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "GET":
            someTemplate.Execute(w, options)
        case "POST":
            press := r.FormValue("submit")
            html, err := MdToHtml(press)

            // This is a dir
            if err != nil {
                notes := &[]Storage{}
                DiscoverFilesAndDirs(press, notes)
                // FIXME: change this title from KB to the original one
                options := &Kb{Title: "KB", Notes: *notes}
                someTemplate.Execute(w, options)

                // This is a file
            } else {
                note := Note{Title: press, Content: template.HTML(html)}
                notesTemplate.Execute(w, note)
            }

        }

    }
}
