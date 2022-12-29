package wp

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/hadziqm/go-svelte/logger"
)


func GetCategories(web string,body string) http.Response {
  pbody,err := json.Marshal(map[string]string{
    "query":body,
  })
  logger.Ignore(err)
  res,err := http.Post(web,"apllication/json",bytes.NewBuffer(pbody))
  return *res
}
