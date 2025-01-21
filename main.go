package main

import (
    "os"
    "path/filepath"
    "log"
    "simple_git_clone_recording/lib"
    "fmt"
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
    dbPath := filepath.Join(gitPath, "objects")


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

        if len(args) != 3 {
            log.Println("Usage: mit commit -m 'test'")
            return
        }

        if args[1] != "-m" {
            log.Println("Usage: mit commit -m 'test'")
            return
        }

        message := args[2]


        workspace := lib.MakeWorkspace()
        filePaths, err := workspace.GetFilePathsFrom(cwd)
        database := lib.MakeDatabase(dbPath)

            entries := [] lib.Entry{}



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

            entries = append(entries, lib.MakeEntry(path, blob.Oid))
        }

        tree := lib.NewTree(entries)
        err = database.Store(tree)

        if err != nil {
            log.Fatalln(err)
        }

        refs := lib.MakeRefs(gitPath)
        parent := refs.ReadHead()


        author := lib.NewAuthor("musa", "musa@email.com")
        commit := lib.NewCommit(parent, tree.Oid, message, author)
        err = database.Store(commit)
        refs.UpdateHead(commit.Oid)

        output := ""

        if parent == "" {
            output += "root-commit "
        } 

            output += fmt.Sprintf("%s %s", commit.Oid, message)

        log.Println(output)



    } else if command == "test" {
        workspace := lib.MakeWorkspace()
        paths, _:= workspace.GetFilePathsFrom(cwd)

        for _, path := range paths {
            fmt.Println(path)

        }
    }



}

