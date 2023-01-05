package wp

import (
	"fmt"
	"reflect"
	"testing"
)

var web string = "https://masjidmoedhararifin.com/graphql"
func TestGetRaw(t *testing.T)  {
  queried := readBody(getCategoriesRaw(web))
  if reflect.TypeOf(queried).String() != "string"{
    t.Fatal("response invalid")
  }
}
func TestGetStruct(t *testing.T){
  structured := GetCategory(web)
  name := structured.Data.Categories.Nodes[0].Name
  if reflect.TypeOf(name).String() != "string"{
    t.Fatal("result invalid")
  }
  fmt.Printf("first category item is %s",name)
}
