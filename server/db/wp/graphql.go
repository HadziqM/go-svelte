package wp
import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/hadziqm/go-svelte/logger"
)

type CategoryList struct{
  Data struct {
    Categories struct {
      Nodes []struct {
        Name string `json:"name"`
        Slug string `json:"slug"`
        Posts struct {
          Nodes []struct {
            Title string `json:"title"`
            Slug string `json:"slug"`
            Content string `json:"content"`
            Date string `json:"date"`
            FeaturedImage struct {
              Node struct {
                Link string `json:"link"`
              } `json:"node"`
            } `json:"featuredImage"`
          } `json:"nodes"`
        } `json:"posts"`
      } `json:"nodes"`
    } `json:"categories"`
  } `json:"data"`
}

func getResponse(web string,body string) http.Response {
  pbody,err := json.Marshal(map[string]string{
    "query":body,
    "variables":"",
  })
  logger.Ignore(err)
  res,err := http.Post(web,"application/json",bytes.NewBuffer(pbody))
  return *res
}
func readBody(res http.Response) string{
  bod,err := ioutil.ReadAll(res.Body)
  logger.Ignore(err)
  return string(bod)
}
func GetCategories(web string) CategoryList{
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
  var categories CategoryList
  err := json.NewDecoder(res.Body).Decode(&categories)
  logger.Fatal(err,"successfully parsed")
  return categories
}
