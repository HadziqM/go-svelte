package download

import "testing"

func TestDownload(test *testing.T)  {
  url := "https://media.discordapp.net/attachments/761607594410115182/1058033671124308028/image-1.png"
  err := Download(url)
  if err != nil{
    test.Errorf("fail on progress")
  }
}
