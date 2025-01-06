package lib

import (
    "os"
    "path/filepath"
)

type Refs struct {
    HeadPath string

}

func MakeRefs(path string) Refs {
    return Refs{HeadPath: filepath.Join(path, "HEAD") }
}

func (r Refs) UpdateHead(oid string) error  {
    lockFile := MakeLockFile(r.HeadPath)
    err := lockFile.HoldForUpdate()

    if err != nil {
        return err
    }

    err = lockFile.Write(oid + "\n")

    if err != nil {
        return err
    }

    return lockFile.Commit()
}


func (r Refs) ReadHead() string {
    if Exists(r.HeadPath) == false {
        return ""
    }

    headData, err := os.ReadFile(r.HeadPath)

    if err != nil {
        return ""
    }

    return string(headData)

}
