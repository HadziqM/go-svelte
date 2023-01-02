package download

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/hadziqm/go-svelte/logger"
)

func Download(url string) {
  fmt.Println("downloading from:",url)
  res,err := http.Get(url)
  logger.Fatal(err,"error on download http request")
  defer res.Body.Close()
  if res.StatusCode != 200{
    logger.FatalNE("error on download bad status code")
  }
  path := splitName(url)
  file,err := os.Create(path)
  if err != nil{
    err2 := os.MkdirAll("../images",os.ModePerm)
    res,_ := os.Create(path)
    file = res
    logger.Fatal(err2,"error on download cant create dir")
  }
  defer file.Close()
  _,err2 := io.Copy(file,res.Body)
  logger.Fatal(err2,"error on download cant copy to file")
}
func GetName(url string)string{
  split := strings.Split(url,"/")
  return split[len(split)-1]
}
func splitName(url string)string{
  split := strings.Split(url, "/")
  return "../images/"+split[len(split)-1]
}
func DownloadOpen(url string){
  path := splitName(url)
  _,err := os.Open(path)
  if err != nil{
    Download(url)
  }else{
    err := os.Remove(path)
    logger.Fatal(err,"error on download cant delete file")
    Download(url)
  }
}
