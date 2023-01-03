package db

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)
//crud test and queries file test

func getDB() sql.DB{
  dbase,err := sql.Open("sqlite3","./test.db")
  if err != nil{
    fmt.Println("fatal error on init")
  }
  defer dbase.Close()
  return *dbase
}
var dbase sql.DB = getDB()
func TestInit(t *testing.T){
  file,err := os.ReadFile("../queries/init.sql")
  if err != nil{
    t.Fatal("cant open file with err:\n",err)
  }
  query := string(file)
  _,err2 := dbase.Exec(query)
  if err2 != nil{
    t.Fatal("cant create table with err:\n",err2)
  }
}
func TestQuery(t *testing.T) {
  fm,err :=dbase.Prepare("INSERT INTO category (slug,name) VALUES (?,?)")
  if err != nil{
    t.Fatal("cant query database with error: ",err)
  }
  _,err2 := fm.Exec("test","name")
  if err2 != nil{
    t.Fatal("cant exect queries with err:\n",err2)
  }
  defer fm.Close()
}
func TestGetQueries(t *testing.T){
  res:= dbase.QueryRow("SELECT name FROM category WHERE slug='test'")
  var name string
  err2 := res.Scan(&name)
  if err2 != nil{
    t.Fatal("cant parse result")
  }
  fmt.Println("the result name is:\n",name)
}
func TestDelete(t *testing.T) {
  fm,err :=dbase.Prepare("DELETE FROM category WHERE slug=?")
  if err != nil{
    t.Fatal("cant query database with error: ",err)
  }
  _,err2 := fm.Exec("test")
  if err2 != nil{
    t.Fatal("cant exect queries")
  }
  defer fm.Close()
}

func TestCheckQueries(t *testing.T){
  res:= dbase.QueryRow("SELECT name FROM category WHERE slug='test'")
  var name string
  err2 := res.Scan(&name)
  if err2 == nil{
    t.Fatal("record is not deleted with value:\n",name)
  }
}

