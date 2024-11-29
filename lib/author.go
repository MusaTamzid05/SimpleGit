package lib

import (
    "time"
    "fmt"
)


type Author struct {
    Oid string
    Name string
    Email string
    Time time.Time
}


func (a Author) Type() string {
    return "author"
}


func (a Author) ToSring() string {
    timestamp := fmt.Sprintf("%d %s", a.Time.Unix(), a.Time.Format("MST"))
    return fmt.Sprintf("%s <%s> %s", a.Name, a.Email, timestamp)
}

func (a *Author) SetOid(oid string) {
    a.Oid = oid
}

func NewAuthor(name, email string) *Author {
    return &Author{Name: name, Email: email, Time: time.Now()}
}
