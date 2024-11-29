package lib

import (
    "os"
    "path/filepath"
    "strings"
)

type Workspace struct {
    path string
    ignores []string

}

func (w Workspace) GetFilePaths() ([] string, error){
    paths := []string{}

    err := filepath.Walk(w.path, func(path string, fileInfo os.FileInfo, err error) error {

        if err != nil {
            return err
        }

        shouldIgnore := false

        for _, ignore := range w.ignores {
            if strings.Contains(path, ignore) {
                shouldIgnore = true
                break
            }
        }

        if shouldIgnore {
            return nil
        }


        paths = append(paths, path)
        return nil
    })

    return paths, err
}

func MakeWorkspace(path string) Workspace {
    ignores := []string{
        ".git",
        ".mit",
    }

    return Workspace{path: path, ignores: ignores}
}
