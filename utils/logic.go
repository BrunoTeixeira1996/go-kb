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
    Color    string
}

// Function to get every dir and file recursv
func DiscoverFilesAndDirsRecur(path string, storage *[]Storage) {
    filepath.Walk(path, func(path string, info os.FileInfo, err error) error {

        if err != nil {
            fmt.Println(err)
        }

        item := Storage{}

        if info.IsDir() {
            item = Storage{Type: "dir", FullPath: path, Name: info.Name(), Color: "#515151"}
        } else {
            item = Storage{Type: "file", FullPath: path, Name: info.Name(), Color: "#f2dd72"}

        }
        *storage = append(*storage, item)

        return nil
    })
}


// Function to get the dirs and files inside a dir
func DiscoverFilesAndDirs(path string, storage *[]Storage) error {
    files, err := ioutil.ReadDir(path)

    if err != nil {
        return fmt.Errorf(err.Error())
    }

    item := Storage{}

    for _, f := range files {
        if f.IsDir() {
            // FIXME: The back slash works but its buggy because sometimes is used twice
            // and that's because at the first time using this a backslash already exists
            item = Storage{Type: "dir", FullPath: path + "/" + f.Name(), Name: f.Name(), Color: "#515151"}
        } else {
            item = Storage{Type: "file", FullPath: path + "/" + f.Name(), Name: f.Name(), Color: "#f2dd72"}

        }

        *storage = append(*storage, item)
    }

    return nil
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
