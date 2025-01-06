package lib

type Commit struct {
    parent string
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

    if c.parent != "" {
        lines += "parent " + c.parent + "\n"
    }

    lines += "author " + c.author.ToSring() +  "\n"
    lines += "commiter " + c.author.ToSring() +  "\n"
    lines += "\n"
    lines +=  c.Message +  "\n"

    return lines
}

func (c *Commit) SetOid(oid string) {
    c.Oid = oid
}



func NewCommit(parent, treeOid, message string, author *Author) *Commit {
    return &Commit{parent: parent, treeOid: treeOid, author: author, Message: message}
}
