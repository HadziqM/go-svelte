package db

import (
	"database/sql"
	"os"
	"github.com/hadziqm/go-svelte/logger"
)

type Category struct{
  Slug string `json:"slug"`
  Name string `json:"name"`
}
type Post struct{
  Slug string `json:"slug"`
  Title string `json:"title"`
  Cdate string `json:"cdate"`
  Content string `json:"content"`
  Category string `json:"category"`
  Views uint `json:"views"`
}
type Comment struct{
  Id uint `json:"id"`
  Name string `json:"string"`
  Content string `json:"content"`
  Photo string `json:"photo"`
  Post string `json:"post"`
}
type Infaq struct{
  Post string `json:"post"`
  Total uint `json:"total"`
  Now uint `json:"now"`  
}
type Donate struct{
  Post string `json:"post"`
  Cdate string `json:"cdate"`
  Name string `json:"name"`
  Amount uint `json:"amount"`
}

    
func Init(dbase *sql.DB)  {
  _,err := dbase.Exec(get_start())
  logger.Fatal(err,"error on db executing init queries","success create table")
}
func get_start() string {
  file,err := os.ReadFile("./queries/init.sql")
  logger.Fatal(err,"error on db read init.sql file","succesfully read file")
  return string(file)
}

