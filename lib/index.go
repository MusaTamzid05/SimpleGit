package lib

import (
    "fmt"
)


type Index struct {
    entries map[string]IndexEntry
    lockFile LockFile

}

func (i *Index) Add(filePath string, blob *Blob) error  {
    entry, err := MakeIndexEntry(filePath, blob.Oid)

    if err != nil {
        return err
    }

    i.entries[filePath] = entry
    return nil
}


func (i Index) WriteUpdate() error  {
    // create a loc file
    // pack and write the header
    // write every entry on a log file
    // generate hash for every entry and save
    // in string variable
    // save the hash of the final string variable
    // commit the log file


    err := i.lockFile.HoldForUpdate()

    if err != nil {
        return fmt.Errorf("Error holding the logfile for index ", err)
    }

    headerString, err := pack("a4N2", 2, uint32(len(i.entries)))

    if err != nil {
        return fmt.Errorf("Error getting header string in logfile for index ", err)
    }

    err = i.Write(headerString)


    if err != nil {
        return fmt.Errorf("Error writing header string in logfile for index ", err)
    }

    allEntriesData := ""

    for _, entry := range i.entries {
        entryStr, err := entry.ToString()

        if err != nil {
            return fmt.Errorf("Error getting entry string ", err)
        }

        err = i.Write(entryStr)

        if err != nil {
            return fmt.Errorf("Error writing entry string ", err)
        }
        allEntriesData += entryStr
    }

    err = i.lockFile.Write(Sha1Hasher(allEntriesData))


    if err != nil {
        return fmt.Errorf("Error writing sha1 of all entries for index ", err)
    }

    return i.lockFile.Commit()

}

func (i Index) Write(data string) error {
    err := i.lockFile.Write(data)

    if err != nil {
        return err
    }

    return nil
}

func MakeIndex(path string) Index {
    return Index{lockFile:MakeLockFile(path), entries: make(map[string]IndexEntry)}
}
