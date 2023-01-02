package logger

import "log"


func Fatal(err error, on string) {
  if err!= nil{
    log.Printf(on)
    log.Fatal(err)
  }
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
func FatalNE(on string) {
  log.Fatal(on)
}
