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

type PageData struct {
    PageTitle string
    Nav       template.HTML
    Content   []NoteStruct
}

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
func indexHandler(page PageData, navData string, baseTemplate *template.Template) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        data := PageData{
            PageTitle: "Index",
            Nav:       template.HTML(navData),
            Content:   []NoteStruct{},
        }
        baseTemplate.Execute(w, data)
    }
}

func pathHandler(notes []NoteStruct, page PageData, navData string, baseTemplate *template.Template) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "GET":
            data := PageData{
                PageTitle: "Notes",
                Nav:       template.HTML(navData),
                Content:   notes,
            }
            baseTemplate.ExecuteTemplate(w, "buttons.html", data)
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

func main1() {
    notesDir := "/home/brun0/workspace/personal-kb/notes/"

    notes := &[]NoteStruct{}
    discover(notesDir, notes)

    port := "9191"

    mux := http.NewServeMux()

    baseTemplate := template.Must(template.ParseFiles("/home/brun0/workspace/personal-kb/templates/base.html", "/home/brun0/workspace/personal-kb/templates/buttons.html"))

    fs := http.FileServer(http.Dir("assets"))
    mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

    navData, _ := ioutil.ReadFile("/home/brun0/workspace/personal-kb/templates/nav.html")
    mux.HandleFunc("/", indexHandler(PageData{}, string(navData), baseTemplate))

    mux.HandleFunc("/notes", pathHandler(*notes, PageData{}, string(navData), baseTemplate))
    http.ListenAndServe(":"+port, mux)

}

type Page struct {
    Title string
}

type OptionsContent struct {
    Title string
    Notes []NoteStruct
}

func indexTest(baseTemplate *template.Template) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        page := Page{
            Title: "Index",
        }
        baseTemplate.Execute(w, page)
    }
}

func optionsHandle(options *OptionsContent, someTemplate *template.Template) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        someTemplate.Execute(w, options)
    }
}

func main() {
    notesDir := "/home/brun0/workspace/personal-kb/notes/"

    notes := &[]NoteStruct{}
    discover(notesDir, notes)
    options := &OptionsContent{Title: "something", Notes: *notes}

    mux := http.NewServeMux()

    baseTemplate := template.Must(template.ParseFiles("/home/brun0/workspace/personal-kb/draft/base.html", "/home/brun0/workspace/personal-kb/draft/index.html"))

    someTemplate := template.Must(template.ParseFiles("/home/brun0/workspace/personal-kb/draft/base.html", "/home/brun0/workspace/personal-kb/draft/options.html"))

    //mux.Handle("/assets/", http.FileServer(http.Dir("assets")))
    mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
    http.HandleFunc("/index", indexTest(baseTemplate))
    http.HandleFunc("/options", optionsHandle(options, someTemplate))

    //http.ListenAndServe(":8080", mux)
    http.ListenAndServe(":8080", nil)
}
