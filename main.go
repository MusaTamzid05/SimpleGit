package main

import (
    "os"
    "path/filepath"
    "log"
    "simple_git_clone_recording/lib"
)

func main() {

    if len(os.Args) == 1 {
        return
    }

    args := os.Args[1:]
    command := args[0]

    cwd, err := os.Getwd()

    if err != nil {
        log.Fatalln(err)
    }


    gitPath := filepath.Join(cwd, ".mit")


    if command == "init" {

        if lib.Exists(gitPath) == false {
            os.Mkdir(gitPath, 0755)
        }

        folders := []string { "refs", "objects"}

        for _ , name := range folders {
            path := filepath.Join(gitPath, name)

            if lib.Exists(path) {
                log.Println("Skipping ", path)
                continue
            }

            err := os.Mkdir(path, 0755)

            if err != nil {
                log.Fatalln(err)
            }
        }

    } else if command == "commit" {
        workspace := lib.MakeWorkspace()
        filePaths, err := workspace.GetFilePathsFrom(cwd)


        dbPath := filepath.Join(gitPath, "objects")
        database := lib.MakeDatabase(dbPath)

        if err != nil {
            log.Fatalln("Dir err ", err)
        }

        for _, path := range filePaths {
            content, err := workspace.ReadFile(path)

            if err != nil {
                log.Fatalln(err)

            }
            blob := lib.NewBlob(content)
            err = database.Store(blob)

            if err != nil {
                log.Fatalln(err)
            }
        }


    }



}

