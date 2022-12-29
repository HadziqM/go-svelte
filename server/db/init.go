package db

import (
	"database/sql"
	"os"
	"github.com/hadziqm/go-svelte/logger"
)

    
func Init(dbase *sql.DB)  {
  _,err := dbase.Exec(get_start())
    logger.Fatal(err,"success create table")
}
func get_start() string {
  file,err := os.ReadFile("./queries/init.sql")
  logger.Fatal(err,"succesfully read file")
  return string(file)
}

