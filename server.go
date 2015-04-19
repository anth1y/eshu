package main
import (
  "gopkg.in/alecthomas/kingpin.v1"
  "fmt"
   log "github.com/Sirupsen/logrus"
  "net"
)

var (
  app = kingpin.New("eshu", "A proxy server for sqldb.")
  srchost = app.Flag("srchost", "Originating host.").Default("127.0.0.1").IPIP()
  dsthost = app.Flag("dsthost", "Destination host.").IP() 
)

func main() {
  kingpin.Version("0.0.1")
  Kingpin.Parse()
  fmt.Printf("Proxying %s to %s", *srchost, *dsthost)
  // Set up our listening server
  server, err := net.Listen("tcp", *dsthost)
  if err != nil {
     log.Fatal(err)
  }
}
