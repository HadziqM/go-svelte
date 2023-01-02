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
  split := strings.Split(url, "/")
  path := "../images/"+split[len(split)-1]
  file,err := os.Create(path)
  if err != nil{
    return err
  }
  defer file.Close()
  _,err = io.Copy(file,res.Body)
  if err != nil{
    return err
  }
  return nil
}
