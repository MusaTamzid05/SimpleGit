package lib

import (
    "os"
    "bytes"
    "encoding/binary"
    "fmt"
)

func Exists(path string) bool {
    _, err := os.Stat(path)

    if os.IsNotExist(err) {
        return false
    }

    return true
}

func pack(magicStr string, number uint16, entries uint32) (string, error){
    // create buffer for the magic string, add padding if 
    // magic string is less than 4 
    // save the number as big endian in buffer
    // save the entries as big endian in buffer
    // return buffer as string


    buffer := new(bytes.Buffer)
    magicBytes := []byte(magicStr)

    for i := len(magicBytes); i < 4; i += 1 {
        magicBytes = append(magicBytes, 0)
    }


    err := binary.Write(buffer, binary.BigEndian, number)

    if err != nil {
        return  "", fmt.Errorf("Error packing number as BigEndian ", err)
    }

    err = binary.Write(buffer, binary.BigEndian, entries)

    if err != nil {
        return  "", fmt.Errorf("Error packing entries as BigEndian ", err)
    }

    bytesData := buffer.Bytes()

    return string(bytesData), nil

}
