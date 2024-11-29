package lib

type Commit struct {
    Oid string
    treeOid string
    author *Author
    Message string
}


func (c Commit) Type() string {
    return "commit"
}


func (c Commit) ToString() string {
    lines := ""
    lines += "tree " + c.treeOid + "\n"
    lines += "author " + c.author.ToSring() +  "\n"
    lines += "commiter " + c.author.ToSring() +  "\n"
    lines += "\n"
    lines +=  c.Message +  "\n"

    return lines
}

func (c *Commit) SetOid(oid string) {
    c.Oid = oid
}



func NewCommit(treeOid, message string, author *Author) *Commit {
    return &Commit{treeOid: treeOid, author: author, Message: message}
}
