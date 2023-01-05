package db

import (
	"database/sql"
)

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
func GetIndex(dbase *sql.DB) IndexOut {
  popular := getConcurrentPost(func(d *sql.DB) []Post {
    return postQueries(d,"ORDER BY views DESC")
  },dbase)
  latest := getConcurrentPost(func(d *sql.DB) []Post {
    return postQueries(d,"ORDER BY cdate DESC")
  },dbase)
  highlight := getConcurrentPost(onlyHightlight,dbase)
  out := IndexOut{
    Popular: <- popular,
    Latest: <- latest,
    Highlight: <-highlight,
  }
  return out
}
