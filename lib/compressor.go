package lib

import (
    "bytes"
    "compress/zlib"
    "os"
    "io"

)


type Compressor struct {

}

func (c Compressor) Compress(content string) ([]byte, error) {
    bytesData := []byte(content)
    var buffer bytes.Buffer
    writer := zlib.NewWriter(&buffer)

    _, err := writer.Write(bytesData)

    if err != nil {
        return nil, err
    }

    err = writer.Close()


    if err != nil {
        return nil, err
    }


    return buffer.Bytes(), nil
}

func (c Compressor) Decompress(path string) (string, error){
    file, err := os.Open(path)

    if err != nil {
        return "", err
    }

    defer file.Close()

    reader, err := zlib.NewReader(file)

    if err != nil {
        return "", err
    }

    defer reader.Close()

    decompressedData, err := io.ReadAll(reader)

    if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
        return "", err

    }


    return string(decompressedData), err
}

func MakeCompressor() Compressor {
    return Compressor{}
}
