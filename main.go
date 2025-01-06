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

        author := lib.NewAuthor("musa", "musa@email.com")
        commit := lib.NewCommit(tree.Oid, message, author)
        err = database.Store(commit)


        if err != nil {
            log.Fatalln(err)
        }

        headPath := filepath.Join(gitPath, "HEAD")

        f, err := os.Create(headPath)

        if err != nil {
            log.Fatalln(err)
        }

        defer f.Close()

        f.WriteString(commit.Oid)
        log.Println("root-commit ", commit.Oid, " ", message)



    } else if command == "test" {
        compressor := lib.MakeCompressor()

        err := filepath.Walk(dbPath, func(path string, fileInfo os.FileInfo, err error) error {

            if err != nil {
                return err
            }

            if fileInfo.IsDir() {
                return nil
            }

            content, err := compressor.Decompress(path)

            if err != nil {
                return err
            }

            log.Println(content)

            return nil

        })


        if err != nil {
            log.Fatalln(err)
        }



    }



}

