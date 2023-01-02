package db

import(
  "strings"
  "database/sql"
  "github.com/hadziqm/go-svelte/logger"
)


func Category_all(dbase *sql.DB){
  rows,err := dbase.Query("SELECT * FROM category")
  logger.Fatal(err,"error on db fetch queries category")
  defer rows.Close()
  for rows.Next(){
    var slug string
    var name string
    err = rows.Scan(&slug,&name)
    logger.Fatal(err,"error on db parsing result to type")
    res := []string{"result:","name is",name,"with slug",slug}
    logger.Print(strings.Join(res," "))
  }
}

