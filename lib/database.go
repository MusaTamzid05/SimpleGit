package lib

import (
    "fmt"
    "crypto/sha1"
    "encoding/hex"
    "path/filepath"
    "math/rand/v2"
    "os"
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
    object.SetOid(oid)


    return d.WriteContent(oid, content)
}

func (d Database) WriteContent(oid, content string) error {
    path1 := oid[:2]
    path2 := oid[2:]

    dirPath := filepath.Join(d.path, path1)

    if Exists(dirPath) == false {
        os.Mkdir(dirPath, 0755)
    }

    compressor := MakeCompressor()
    compressedData, err := compressor.Compress(content)

    if err != nil {
        return err
    }

    tempPath := filepath.Join(dirPath, d.generateTempName())

    f, err := os.Create(tempPath)

    if err != nil {
        return err
    }

    defer f.Close()

    _, err = f.Write(compressedData)


    if err != nil {
        return err
    }

    
    objectPath := filepath.Join(dirPath, path2)
    err = os.Rename(tempPath, objectPath) 

    if err != nil  {
        return err
    }


    return nil
}

func (d Database) generateTempName() string {
    chars := "abcdefghijklmnopqrstuvwxyz"
    name := ""

    for i := 0; i < 6; i += 1 {
        index := rand.IntN(len(chars))
        name +=  string(chars[index])
    }

    return name
}

func MakeDatabase(path string) Database {
    return Database{path: path}
}


