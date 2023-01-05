package db

import (
	"database/sql"
	"fmt"

	"github.com/hadziqm/go-svelte/logger"
)

var noInfaq string = "SELECT post.slug AS Post,title,cdate,image FROM linked INNER JOIN post ON post.slug = linked.post INNER JOIN category ON category.slug = linked.category WHERE category.slug != 'infaq' "
var highlightOnly string = "SELECT post.slug AS Post,title,cdate,image FROM linked INNER JOIN post ON post.slug = linked.post INNER JOIN category ON category.slug = linked.category WHERE category.slug = 'highlight' ORDER BY cdate DESC LIMIT 5"

func postQueries(dbase *sql.DB,orderby string) []Post {
  fmt.Println("invoked")
  query := noInfaq + orderby + " LIMIT 5"
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
