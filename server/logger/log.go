package logger

import "log"



func Fatal(err error,exp string ) {
  if err!= nil{
    log.Fatal(err)
  }
  log.Printf(exp)
}
func Ignore(err error)  {
  if err != nil{
    log.Println(err)
  }
}
func Print(exp string)  {
  log.Printf(exp)
}
