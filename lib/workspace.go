package lib

import (
    "os"
    "path/filepath"
    "strings"
)

type Workspace struct {
    ignores []string

}

func (w Workspace) GetFilePathsFrom(path string) ([] string, error){
    paths := []string{}

    err := filepath.Walk(path, func(currentPath string, fileInfo os.FileInfo, err error) error {

        if err != nil {
            return err
        }

        if fileInfo.IsDir() {
            return nil
        }

        shouldIgnore := false

        for _, ignore := range w.ignores {
            if strings.Contains(currentPath, ignore) {
                shouldIgnore = true
                break
            }
        }

        if shouldIgnore {
            return nil
        }

        // remove extra string so that path is relative to this dir only
        paths = append(paths, currentPath[len(path) + 1 : len(currentPath)] )
        return nil
    })

    return paths, err
}

func (w Workspace) ReadFile(path string) (string, error){
    data, err := os.ReadFile(path)

    if err != nil {
        return "", err
    }

    return string(data), nil
}

func MakeWorkspace() Workspace {
    ignores := []string{
        ".git",
        ".mit",
    }

    return Workspace{ignores: ignores}
}
