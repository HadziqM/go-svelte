package db

import (
	"database/sql"
	"fmt"

	"github.com/hadziqm/go-svelte/logger"
)

var noInfaq string = "SELECT a.slug,title,cdate,image FROM post a LEFT JOIN linked b ON a.slug = b.post WHERE b.category!='infaq' "
var highlightOnly string = "SELECT a.slug,title,cdate,image FROM post a LEFT JOIN linked b ON a.slug = b.post WHERE b.category = 'headline' ORDER BY cdate DESC LIMIT 5"


func postQueries(dbase *sql.DB,orderby string) []Post {
  fmt.Println("invoked")
  query := noInfaq + orderby + " LIMIT 5"
  fmt.Println(query)
  rows,err := dbase.Query(query)
  logger.Fatal(err,"error on DB at query post index")
  defer rows.Close()
  var out []Post
  for rows.Next(){
    var post Post
    rows.Scan(&post.Slug,&post.Title,&post.Cdate,&post.Image)
    fmt.Println(post)
    out = append(out, post)
  }
  fmt.Println("done")
  return out
}
func onlyHightlight(dbase *sql.DB) []Post {
  rows,err := dbase.Query(highlightOnly)
  logger.Fatal(err,"error on DB query post highlight")
  defer rows.Close()
  var out []Post
  for rows.Next(){
    var post Post
    rows.Scan(&post.Slug,&post.Title,&post.Cdate,&post.Image)
    out = append(out, post)
  }
  return out
}
func getSpecificPost(dbase *sql.DB,slug string) Post {
  row:= dbase.QueryRow("SELECT title,content from post WHERE slug=?",slug)
  var post Post
  err :=row.Scan(&post.Title,&post.Content)
  logger.Fatal(err,"error on DB query specific post")
  return post
}
func addViewCount(dbase *sql.DB,slug string){
  _,err := dbase.Exec("UPDATE post SET views = views + 1 WHERE slug=?",slug)
  logger.Fatal(err,"error on DB increment post views")
}
func getPostComment(dbase *sql.DB,slug string) []Comment{
  rows,err := dbase.Query("SELECT name,content,photo,cdate FROM comments WHERE post=?",slug)
  defer rows.Close()
  logger.Fatal(err,"error on DB getting post comments")
  var out []Comment
  for rows.Next(){
    var com Comment
    rows.Scan(&com.Name,&com.Content,&com.Photo,&com.Cdate)
    out = append(out, com)
  }
  return out
}
