package download

import "testing"


var url string = "https://media.discordapp.net/attachments/761607594410115182/1058033671124308028/image-1.png"
func TestDownload(test *testing.T)  {
  err := Download(url)
  if err != nil{
    test.Errorf("fail on progress %s",err)
  }
}
func TestDownload2(test *testing.T)  {
  err := DownloadOpen(url)
  if err != nil{
    test.Errorf("fail on progress %s",err)
  }
}
