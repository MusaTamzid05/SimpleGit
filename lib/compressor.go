package lib

import (
    "bytes"
    "compress/zlib"

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

func (c Compressor) Decompress(path string) string {
    return ""
}

func MakeCompressor() Compressor {
    return Compressor{}
}
