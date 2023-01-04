package db

import (
	"fmt"
	"testing"
)

func TestQueries(t *testing.T) {
  res := GetIndex(&dbase)
  fmt.Println(res.Latest[0].Title)
}
