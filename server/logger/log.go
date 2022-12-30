package logger

import "log"


func Fatal(err error, on string, exp string) {
  if err!= nil{
    log.Printf(on)
    log.Fatal(err)
  }
  log.Printf(exp)
}
func Ignore(err error, on string)  {
  if err != nil{
    log.Printf(on)
    log.Println(err)
  }
}
func Print(exp string)  {
  log.Printf(exp)
}
