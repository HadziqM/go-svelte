package db

import (
	"database/sql"
)

type IndexOut struct{
  Highlight []Post `json:"highlight"`
  Popular []Post `json:"popular"`
  Latest []Post `json:"latest"`
}
type PostOut struct{
  Post Post `json:"post"`
  Comment []Comment `json:"comment"`
}

func getConcurrentPost(f func(*sql.DB) []Post,dbase *sql.DB)<-chan []Post{
  r := make(chan []Post)
  go func(){
    defer close(r)
    r <- f(dbase)
  }()
  return r
}
func getConcurrentComment(f func(*sql.DB,string) []Comment,dbase *sql.DB,slug string) <-chan []Comment {
  r:= make (chan []Comment)
  go func(){
    defer close(r)
    r <- f(dbase,slug)
  }()
  return r
}
func getConcurrentPage(f func(*sql.DB,string) Post,dbase *sql.DB,slug string) <-chan Post {
  r:= make (chan Post)
  go func(){
    defer close(r)
    r <- f(dbase,slug)
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
func GetSpecificPost(dabase *sql.DB,slug string) PostOut{
  post := getConcurrentPage(getSpecificPost,dabase,slug)
  comment := getConcurrentComment(getPostComment,dabase,slug)
  go addViewCount(dabase,slug)
  out := PostOut{
    Post: <-post,
    Comment: <-comment,
  }
  return out
}
