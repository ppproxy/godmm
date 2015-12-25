package main

import (
  "encoding/json"
  "flag"
  "fmt"
  "time"

  "github.com/golang/glog"
  "github.com/junz/opendmm"
)

func main() {
  flag.Set("stderrthreshold", "FATAL")
  flag.Parse()
  for _, arg := range flag.Args() {
    metach := opendmm.Search(arg)
    select {
    case meta, ok := <-metach:
      if ok {
        metajson, _ := json.MarshalIndent(meta, "", "  ")
        fmt.Println(string(metajson))
      } else {
        glog.Exit("Not found")
      }
    case <-time.After(10 * time.Second):
      glog.Fatal("Timeout")
    }
  }
}
