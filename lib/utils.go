package lib

import (
    "os"
)

func Exists(path string) bool {
    _, err := os.Stat(path)

    if os.IsNotExist(err) {
        return false
    }

    return true
}
