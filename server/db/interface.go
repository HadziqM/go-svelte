package db

import (
	"database/sql"

	"github.com/hadziqm/go-svelte/logger"
)

func GetAllPost(dbase *sql.DB,order string) []Post {
  rows,err := dbase.Query("SELECT * FROM post ORDER BY ? DESC",order)
  logger.Fatal(err,"erron on db querying all post data")
  defer rows.Close()
  var out []Post
  for rows.Next(){
    var post Post
    err := rows.Scan(&post.Slug,&post.Title,&post.Cdate,&post.Content,&post.Image,&post.Views)
    logger.Fatal(err,"error on db parsing all post data")
    out = append(out, post)
  }
  return out
}
func GetAllCategories(dbase *sql.DB) []Category {
  rows,err := dbase.Query("SELECT * FROM category ORDER BY name ASC")
  logger.Fatal(err,"erron on db querying all Category data")
  defer rows.Close()
  var out []Category
  for rows.Next(){
    var category Category
    err := rows.Scan(&category.Slug,&category.Name)
    logger.Fatal(err,"error on db parsing all category data")
    out = append(out, category)
  }
  return out
}
func GetAllLink(dbase *sql.DB) []Category {
  rows,err := dbase.Query("SELECT * FROM linked ORDER BY category ASC")
  logger.Fatal(err,"erron on db querying all Category data")
  defer rows.Close()
  var out []Category
  for rows.Next(){
    var category Category
    err := rows.Scan(&category.Slug,&category.Name)
    logger.Fatal(err,"error on db parsing all category data")
    out = append(out, category)
  }
  return out
}
func GetSpeComments(dbase *sql.DB,post string) []Comment {
  rows,err := dbase.Query("SELECT * FROM comments WHERE post=? ORDER BY cdate DESC",post)
  logger.Fatal(err,"erron on db querying all Comment data")
  defer rows.Close()
  var out []Comment
  for rows.Next(){
    var comment Comment
    err := rows.Scan(&comment.Id,&comment.Name,&comment.Content,&comment.Photo,&comment.Post,&comment.Cdate)
    logger.Fatal(err,"error on db parsing all commentdata")
    out = append(out, comment)
  }
  return out
}
func GetSpePost(dbase *sql.DB,slug string) Post {
  row := dbase.QueryRow("SELECT * FROM post WHERE slug=?",slug)
  var post Post
  err := row.Scan(&post.Slug,&post.Title,&post.Cdate,&post.Content,&post.Image,&post.Views)
  logger.Fatal(err,"error on DB get specific post parsing")
  return post
}
