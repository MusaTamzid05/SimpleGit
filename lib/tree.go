package lib

import (
    "bytes"
    "fmt"
    "encoding/hex"
    "sort"
    "log"
    "strings"
)

type Tree struct {
    Entries []Entry
    Oid string
    Nodes map[string]interface{} // datatype are basically tree and entry
}

func (t Tree) ToString() string {
    // @TODO Update this with Nodes
    MODE := "100644"
    var buffer bytes.Buffer

    for _, entry := range t.Entries {
        formattedEntry := fmt.Sprintf("%s %s", MODE, entry.Name)
        buffer.WriteString(formattedEntry + "\x00")

        oidBytes, err := hex.DecodeString(entry.Oid)

        if err != nil {
            log.Println(err)
            return ""
        }

        buffer.Write(oidBytes)

    }
    return buffer.String() 
}

func (t Tree) Type() string  {
    return "tree"
}


func (t *Tree) SetOid(oid string) {
    t.Oid = oid
}

func (t Tree) Traverse(database Database) error  {
    for _, value := range t.Nodes {
        tree, ok := value.(*Tree)

        if ok {
            tree.Traverse(database)
            continue
        } 

        entry, ok := value.(*Entry)

        if ok {
            err := database.Store(entry.Blob)

            if err != nil {
                return err
            }
        } 

    }

    return database.Store(&t)

}

func buildTree(entries []Entry) *Tree {
    nodes := make(map[string]interface{})

    // if an entry is a dir path, we tree it as a tree
    // and save it in the nodes as tree
    // files are considered child


    for _, entry := range entries {
        path := entry.Name

        if strings.Contains(path, "/") {
            pathData := strings.Split(path, "/")
            parentDirName := pathData[0]

            _, keyExists := nodes[parentDirName]

            if keyExists {
                continue
            }

            childEntries := []Entry{}
            tempEntries := make([]Entry, len(entries))
            copy(tempEntries, entries)


            for _, tempEntry := range tempEntries {
                // @TODO : it will fail if parent and child have same dir name
                // fix this

                if strings.Contains(tempEntry.Name, parentDirName) == false {
                    continue
                }


                tempEntry.Name = tempEntry.Name[len(parentDirName) + 1 : len(tempEntry.Name)]
                childEntries = append(childEntries, tempEntry)

            }

            nodes[parentDirName] = buildTree(childEntries)
            continue
        }

        fmt.Println("Storing ", path, " as entry")

        nodes[path] = &entry

    }

    return &Tree{Nodes: nodes}
}

func NewTree(entries []Entry) *Tree {
    sort.Slice(entries, func(i, j int) bool {
        return entries[i].Name < entries[j].Name
    })

    return buildTree(entries)
}
