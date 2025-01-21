package lib

type Entry struct {
    Name string
    Oid string
    Blob *Blob

}

func MakeEntry(name string, blob *Blob) Entry {
    return Entry{Name: name, Oid: blob.Oid, Blob: blob}
}


