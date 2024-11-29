package lib

import (
    "bytes"
    "fmt"
    "encoding/hex"
    "sort"
    "log"
)

type Tree struct {
    Entries []Entry
    Oid string
}

func (t Tree) ToString() string {
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

func NewTree(entries []Entry) *Tree {
    sort.Slice(entries, func(i, j int) bool {
        return entries[i].Name < entries[j].Name
    })

    return &Tree{Entries: entries}
}
