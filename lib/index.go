package lib


type Index struct {
    entries map[string]IndexEntry

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
    return nil
}

func MakeIndex() Index {
    return Index{entries: make(map[string]IndexEntry)}
}
