package lib

import (
    "os"
)

const (
    REGULAR_MODE = "100644"
    EXECUTABLE_MODE = "100755"
)

type Entry struct {
    Name string
    Oid string
    Blob *Blob
    Info os.FileInfo

}

func (e Entry) Mode() string {
    mode := e.Info.Mode()

    if mode.IsRegular() && (mode.Perm()&0111 != 0) {
        return EXECUTABLE_MODE
    }

    return REGULAR_MODE
}

func MakeEntry(name string, blob *Blob) (Entry, error){
    info, err := os.Stat(name)

    if err != nil {
        return Entry{}, err
    }

    return Entry{Name: name, Oid: blob.Oid, Blob: blob, Info: info}, nil
}


