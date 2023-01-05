package db

import (
	"fmt"
	"testing"
)

var web string = "https://masjidmoedhararifin.com/graphql"
func TestUpdate(t *testing.T){
  Update(&dbase,web)
  res := dbase.QueryRow("SELECT name FROM category WHERE slug='agenda'")
  var name string
  err := res.Scan(&name)
  if err != nil{
    t.Fatal("no records on db with err:\n",err)
  }else{
    fmt.Println("name of agenda slug is :",name)
  }
}

func TestQueries(t *testing.T) {
  GetIndex(&dbase)
  fmt.Println("index queried")
}
func TestView(t *testing.T) {
  res := GetSpecificPost(&dbase,"youtube")
  fmt.Println(res)
  row := dbase.QueryRow("SELECT views FROM post WHERE slug='youtube'")
  var view uint
  err := row.Scan(&view)
  if err!=nil{
    t.Fatal("error query")
  }
  fmt.Println("post view is now :",view)
  if view != 1{
    t.Fatal("view isnt incremented")
  }
}
