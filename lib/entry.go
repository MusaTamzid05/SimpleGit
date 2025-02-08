package lib

import (
    "os"
    "fmt"
    "bytes"
    "encoding/binary"
)

const (
    REGULAR_MODE = "100644"
    EXECUTABLE_MODE = "100755"

    INDEX_REGULAR_MODE = 100644
    INDEX_EXECUTABLE_MODE = 100755
    MAX_PATH_SIZE = 0xfff
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

type IndexEntry struct {
    CTime int64
    CTimeNSec int32

    MTime int64
    MTimeNSec int32

    Mode os.FileMode
    Size int64
    Oid string
    Flags int
    Path string

}

func (i IndexEntry) ToString() (string, error) {
    // object id must be 40 bits long
    // write all the stuct data as BigEndian in buffer
    // return the buffer

    if len(i.Oid) != 40 {
        return "", fmt.Errorf("IndexEntry OID must 40 bits")
    }

    buffer := new(bytes.Buffer)

    if err := binary.Write(buffer, binary.BigEndian, i.CTime); err != nil {
        return "", fmt.Errorf("Error converting IndexEntry variable ", err)
    }


    if err := binary.Write(buffer, binary.BigEndian, i.CTimeNSec); err != nil {
        return "", fmt.Errorf("Error converting IndexEntry variable ", err)
    }


    if err := binary.Write(buffer, binary.BigEndian, i.MTime); err != nil {
        return "", fmt.Errorf("Error converting IndexEntry variable ", err)
    }


    if err := binary.Write(buffer, binary.BigEndian, i.MTimeNSec); err != nil {
        return "", fmt.Errorf("Error converting IndexEntry variable ", err)
    }

    if err := binary.Write(buffer, binary.BigEndian, int32(i.Mode)); err != nil {
        return "", fmt.Errorf("Error converting IndexEntry variable ", err)
    }

    if err := binary.Write(buffer, binary.BigEndian, i.Size); err != nil {
        return "", fmt.Errorf("Error converting IndexEntry variable ", err)
    }

    oidBytes, err := decodeHex(i.Oid)

    if err != nil {
        return "", fmt.Errorf("Error decoding hex for oid ", err)
    }

    if err := binary.Write(buffer, binary.BigEndian, oidBytes); err != nil {
        return "", fmt.Errorf("Error converting IndexEntry variable ", err)
    }


    if err := binary.Write(buffer, binary.BigEndian, int32(i.Flags)); err != nil {
        return "", fmt.Errorf("Error converting IndexEntry variable ", err)
    }

    pathBytes := []byte(i.Path)

    if err := binary.Write(buffer, binary.BigEndian, pathBytes); err != nil {
        return "", fmt.Errorf("Error converting IndexEntry variable ", err)
    }

    err = buffer.WriteByte(0)

    if err != nil {
        return "", fmt.Errorf("Error creating buffer to IndexEntry string ", err)
    }

    return buffer.String(), nil

}

func MakeIndexEntry(path, oid string) (IndexEntry, error){
    // set file mode , executable or readable
    // set length
    // create entry with the value

    stat, err := os.Stat(path)

    if err != nil {
        return IndexEntry{}, err
    }

    var mode os.FileMode

    if stat.Mode()&0111 != 0 {
        mode =  INDEX_EXECUTABLE_MODE
    } else {
        mode = INDEX_REGULAR_MODE
    }

    flags := len(path)

    if flags > MAX_PATH_SIZE {
        flags = MAX_PATH_SIZE
    }

    return IndexEntry {
        CTime:  stat.ModTime().Unix(),
        CTimeNSec:  int32(stat.ModTime().Nanosecond()),

        MTime:  stat.ModTime().Unix(),
        MTimeNSec:  int32(stat.ModTime().Nanosecond()),
        Mode: mode,
        Size: stat.Size(),
        Oid: oid,
        Flags: flags,
        Path: path,
    }, nil


}


