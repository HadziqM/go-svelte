package wp
import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/hadziqm/go-svelte/logger"
)


func getResponse(web string,body string) http.Response {
  pbody,err := json.Marshal(map[string]string{
    "query":body,
    "variables":"",
  })
  logger.Ignore(err)
  res,err := http.Post(web,"apllication/json",bytes.NewBuffer(pbody))
  return *res
}
func readBody(res http.Response) string{
  bod,err := ioutil.ReadAll(res.Body)
  logger.Ignore(err)
  return string(bod)
}
func GetCategories(web string){
  queries := `{
  categories(where: {orderby: NAME}) {
    nodes {
      name
      slug
      posts {
        nodes {
          slug
          title
          content
          featuredImage {
            node {
              link
            }
          }
          date
        }
        }
      }
    }
  }`
  res := getResponse(web,queries)
  logger.Print(readBody(res))
}
