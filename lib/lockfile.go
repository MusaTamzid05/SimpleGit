package lib

import (
    "os"
    "errors"
)

type LockFile struct {
    path string
    lockPath string
    lockPtr *os.File
}

func MakeLockFile(path string) LockFile  {
    return LockFile{path: path, lockPath: path + ".lock" }
}

func (l *LockFile) HoldForUpdate() error {
    var err error
    l.lockPtr, err = os.OpenFile(l.lockPath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0644)

    if err != nil {
        return err
    }

    return nil
}

func (l *LockFile) Write(content string) error {
    _, err := l.lockPtr.WriteString(content)
    return err
}



func (l *LockFile) Commit() error {
    if l.lockPtr == nil {
        return errors.New("Not holding any lockfile")
    }

    l.lockPtr.Close()

    err := os.Rename(l.lockPath, l.path)
    return err
}
