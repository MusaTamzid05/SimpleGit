package lib

type Index struct {

}

func (i Index) Add(filePath string, blob *Blob) error  {
    return nil
}


func (i Index) WriteUpdate() error  {
    return nil
}

func MakeIndex() Index {
    return Index{}
}
