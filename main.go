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

    } else if command == "add" {



        if len(args) != 2 {
            log.Println("Usage: mit add 'file_path'")
            return
        }

        filePath := args[1]
        workspace := lib.MakeWorkspace()
        database := lib.MakeDatabase(dbPath)
        index := lib.MakeIndex(filepath.Join(gitPath, "index"))
        data, err := workspace.ReadFile(filePath)

        if err != nil {
            log.Fatalln("Cannot read file for adding ", err)
        }

        blob := lib.NewBlob(data)
        err = database.Store(blob)

        if err != nil {
            log.Fatalln("Cannot blob file for adding ", err)
        }

        err = index.Add(filePath, blob)


        if err != nil {
            log.Fatalln("Cannot add blob file to index ", err)
        }

        err = index.WriteUpdate()


        if err != nil {
            log.Fatalln("Cannot write update with index add", err)
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
            //err = database.Store(blob)

            if err != nil {
                log.Fatalln(err)
            }
            

            entry, err := lib.MakeEntry(path, blob)

            if err != nil {
                log.Println("Error adding entry ", err)
                continue
            }

            entries = append(entries, entry)
        }

        tree := lib.NewTree(entries)
        //err = database.Store(tree)
        err = tree.Traverse(database)

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
        entries := [] lib.Entry{}
        database := lib.MakeDatabase(dbPath)

        for _, path := range paths {
            content, err := workspace.ReadFile(path)

            if err != nil {
                log.Fatalln(err)

            }
            blob := lib.NewBlob(content)
            //err = database.Store(blob)

            if err != nil {
                log.Fatalln(err)
            }

            entry, err := lib.MakeEntry(path, blob)

            if err != nil {
                log.Println("Error adding entry ", err)
                continue
            }
            entries = append(entries, entry)

        }

        tree := lib.NewTree(entries)
        //err = database.Store(tree)
        err = tree.Traverse(database)

    }



}

