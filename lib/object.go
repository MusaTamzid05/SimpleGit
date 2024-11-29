package lib

type Object interface  {
    Type() string
    ToString() string
    SetOid(oid string)

}
