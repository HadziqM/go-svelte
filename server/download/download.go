package download

import (
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
)

func Download(url string) error {
  res,err := http.Get(url)
  if err != nil{
    return err
  }
  defer res.Body.Close()
  if res.StatusCode != 200{
    return errors.New("request error")
  }
  path := splitName(url)
    file,err := os.Create(path)
  if err != nil{
    return err
  }
  defer file.Close()
  _,err2 := io.Copy(file,res.Body)
  if err2 != nil{
    return err2
  }
  return nil
}
func GetName(url string)string{
  split := strings.Split(url,"/")
  return split[len(split)-1]
}
func splitName(url string)string{
  split := strings.Split(url, "/")
  return "../images/"+split[len(split)-1]
}
func DownloadOpen(url string) error{
  path := splitName(url)
  _,err := os.Open(path)
  if err != nil{
    return Download(url)
  }else{
    err := os.Remove(path)
    if err != nil{
      return err
    }
    return Download(url)
  }
}
