package db

import (
	"database/sql"

	"github.com/hadziqm/go-svelte/db/wp"
	"github.com/hadziqm/go-svelte/download"
	"github.com/hadziqm/go-svelte/logger"
)

func Update(dbase *sql.DB,web string)  {
  raw := wp.GetCategories(web)
  category := getCategory(dbase)
  post := getPost(dbase)
  //truncate link table here
  Truncate(dbase,"linked")
  for _,j := range raw.Data.Categories.Nodes{
    var is_new bool = true
    for _,k := range category{
      if k.Slug == j.Slug{
        //update categories here
        updateCategories(dbase,k)
        is_new = false
        break
      }
    }
    if is_new{
      //create new categories here
      cat := Category{j.Slug,j.Name}
      newCategories(dbase,cat)
    }
    for _,l := range j.Posts.Nodes{
      var are_new bool = true
      for _,i := range post{
        if i.Slug == l.Slug{
          //update post here
          updatePost(dbase,i)
          are_new = false
          break
        }
      }
      if are_new{
        //create new post here
        po := Post{l.Slug,l.Title,l.Date,l.Content,l.FeaturedImage.Node.Link,0}
        newPost(dbase,po)
      }
      //create new link row here
      li := Linked{j.Slug,l.Slug}
      newLink(dbase,li)
    }
  }
}
func getCategory(dbase *sql.DB) []Category{
  rows,err := dbase.Query("SELECT * FROM category")
  logger.Fatal(err,"erron on DB fetch category")
  var table []Category
  defer rows.Close()
  var cat Category
  for rows.Next(){
    err = rows.Scan(&cat.Slug,&cat.Name)
    logger.Fatal(err,"error on DB parsing category")
    table = append(table, cat)
  }
  return table
}
func getPost(dbase *sql.DB) []Post{
  rows,err := dbase.Query("SELECT * FROM post")
  logger.Fatal(err,"erron on DB fetch post")
  var table []Post
  defer rows.Close()
  var cat Post
  for rows.Next(){
    err = rows.Scan(&cat.Slug,&cat.Title,&cat.Cdate,&cat.Content,&cat.Views)
    logger.Fatal(err,"error on DB parsing post")
  }
  return table
}
func updateCategories(dbase *sql.DB,category Category){
  fm,err := dbase.Prepare("UPDATE category SET name=? WHERE slug=?")
  logger.Fatal(err,"error on DB prepare update categories")
  fm.Exec(category.Name,category.Slug)
  defer fm.Close()
}
func updatePost(dbase *sql.DB,post Post){
  fm,err := dbase.Prepare("UPDATE post SET title=?,content=?,cdate=?,image=? WHERE slug=?")
  logger.Fatal(err,"error on DB prepare update post")
  err2 := download.DownloadOpen(post.Image)
  logger.Fatal(err2,"error on DB download post image")
  fm.Exec(post.Title,post.Content,post.Cdate,download.GetName(post.Image),post.Slug)
  defer fm.Close()
}
func newCategories(dbase *sql.DB, category Category)  {
  fm,err := dbase.Prepare("INSERT INTO category (slug,name) VALUES (?,?)")
  logger.Fatal(err,"error on DB prepare insert category")
  fm.Exec(category.Slug,category.Name)
  defer fm.Close()
}
func newPost(dbase *sql.DB,post Post)  {
  fm,err := dbase.Prepare("INSERT INTO post (slug,title,cdate,content,image) VALUES (?,?,?,?,?)")
  logger.Fatal(err,"error on DB prepare insert post")
  err2 := download.Download(post.Image)
  logger.Fatal(err2,"error on DB create post image")
  fm.Exec(post.Slug,post.Title,post.Cdate,post.Content,download.GetName(post.Image))
  defer fm.Close()
}
func newLink(dbase *sql.DB,link Linked)  {
  fm,err := dbase.Prepare("INSERT INTO linked (post,category) VALUES (?,?)")
  logger.Fatal(err,"error on DB prepare insert linked")
  fm.Exec(link.Post,link.Category)
  defer fm.Close()
}
func Truncate(dbase *sql.DB,table string)  {
  fm,err := dbase.Prepare("DELETE FROM ?")
  logger.Fatal(err,"error on DB prepare truncate")
  fm.Exec(table)
  defer fm.Close()
}
