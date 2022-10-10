package utils

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"

    "github.com/gomarkdown/markdown"
)

// Struct that stores Files and Folders paths
type Storage struct {
    Type     string
    FullPath string
    Name     string
}

// Function to get every dir and file recursv
func Discover(path string, storage *[]Storage) {
    filepath.Walk(path, func(path string, info os.FileInfo, err error) error {

        if err != nil {
            fmt.Println(err)
        }

        item := Storage{}

        if info.IsDir() {
            item = Storage{Type: "dir", FullPath: path, Name: info.Name()}
        } else {
            item = Storage{Type: "file", FullPath: path, Name: info.Name()}

        }
        *storage = append(*storage, item)

        return nil
    })
}

// Function that transforms .md syntax to .html
func MdToHtml(file string) (string, error) {
    content, err := ioutil.ReadFile(file)

    if err != nil {
        return "", err
    }

    html := markdown.ToHTML(content, nil, nil)

    return string(html), nil

}
