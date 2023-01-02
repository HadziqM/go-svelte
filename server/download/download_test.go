package download

import "testing"


var url string = "https://media.discordapp.net/attachments/761607594410115182/1058033671124308028/image-1.png"
func TestDownload(test *testing.T)  {
  Download(url)
}
func TestDownload2(test *testing.T)  {
  DownloadOpen(url)
}
