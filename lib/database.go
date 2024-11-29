package lib

import (
    "fmt"
    "crypto/sha1"
    "encoding/hex"
)

type Database struct {
    path string
}


func (d Database) Store(object Object) error {
    str := object.ToString()
    content := fmt.Sprintf("%s %d%s%s",
                        object.Type(),
                        len(str),
                        "\x00",
                        str,
                    )

    hash := sha1.New()
    hash.Write([]byte(content))
    hashValue := hash.Sum(nil)
    oid := hex.EncodeToString(hashValue)
    fmt.Println(oid)


    return nil
}

func MakeDatabase(path string) Database {
    return Database{path: path}
}


