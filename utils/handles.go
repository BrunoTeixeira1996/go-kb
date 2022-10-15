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
func KbHandle(kb *Kb, someTemplate *template.Template, notesTemplate *template.Template) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "GET":
            someTemplate.Execute(w, kb)
        case "POST":
            press := r.FormValue("submit")
            back := r.FormValue("back")

            // Pressed the back button
            if len(back) > 0 {
                // TODO: make here the back button logic
                temp := strings.Split(back, "/")
                temp = temp[:len(temp)-1]
                new := strings.Join(temp, "/")

                notes := &[]Storage{}
                DiscoverFilesAndDirs(new, notes)

                // BUG HERE
                // NEED TO USE THE RIGHT PATH WHEN GOING ONCE BACK AND THAT'S NOT HAPPENING
                options := &Kb{Title: "TEMP", Notes: *notes}
                someTemplate.Execute(w, options)

                // Pressed another button
            } else {
                html, err := MdToHtml(press)

                // This is a dir
                if err != nil {
                    notes := &[]Storage{}
                    DiscoverFilesAndDirs(press, notes)

                    options := &Kb{Title: press, Notes: *notes}
                    someTemplate.Execute(w, options)

                    // This is a file
                } else {
                    note := Note{Title: press, Content: template.HTML(html)}
                    notesTemplate.Execute(w, note)
                }

            }

        }

    }
}
