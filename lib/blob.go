package lib

type Blob struct {
    Oid string
    data string
}

func (b Blob) Type() string  {
    return "blob"
}


func (b Blob) ToString() string  {
    return b.data
}


func (b *Blob) SetOid(oid string) {
    b.Oid = oid
}

func NewBlob(content string) *Blob {
    return &Blob{data: content}

}
