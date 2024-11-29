package lib

type Entry struct {
    Name string
    Oid string

}

func MakeEntry(name, oid string) Entry {
    return Entry{Name: name, Oid: oid}
}


