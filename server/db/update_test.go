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
  res := GetIndex(&dbase)
  fmt.Println(res)
}
