package wp

import "github.com/hadziqm/go-svelte/logger"


func GetTestCategories(web string){
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
