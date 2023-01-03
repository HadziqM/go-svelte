package db

import "database/sql"

type IndexOut struct{
  Highlight []Post `json:"highlight"`
  Popular []Post `json:"popular"`
  Latest []Post `json:"latest"`
}

func getConcurrentPost(f func(*sql.DB) []Post,dbase *sql.DB)<-chan []Post{
  r := make(chan []Post)
  go func(){
    defer close(r)
    r <- f(dbase)
  }()
  return r
}
