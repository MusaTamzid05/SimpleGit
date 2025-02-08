package lib

import (
    "os"
    "bytes"
    "encoding/binary"
    "encoding/hex"
    "fmt"
    "log"
    "crypto/sha1"
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

func decodeHex(hexString string)([]byte, error) {
    // hex string must be of even length
    // create a slice of bytes where
    // every two hex will be converted
    // to bytes and saved in a slice (LOL)


    if len(hexString) % 2 != 0 {
        return nil, fmt.Errorf("HexString must be of even length")
    }

    byteSlices := make([]byte, len(hexString) / 2)


    for i := 0; i < len(hexString); i += 2 {
        byteValue, err := hexToByte(hexString[i : i + 2])

        if err != nil {
            log.Fatalln("Error converting hex to byte ", err)
        }

        byteSlices[i / 2] = byteValue

    }

    return byteSlices, nil


}

func hexToByte(hexString string)(byte, error) {
    var result byte
    _, err := fmt.Sscanf(hexString, "%2x", &result)

    if err != nil {
        return 0, err
    }

    return result, nil
}

func Sha1Hasher(input string) string {
    algo := sha1.New()
    algo.Write([]byte(input))
    byteSlice := algo.Sum(nil)

    return hex.EncodeToString(byteSlice)



}
