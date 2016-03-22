package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"time"

	"github.com/golang/glog"
	"github.com/junzh0u/opendmm"
)

func search(query string) {
	metach := opendmm.Search(query)
	select {
	case meta, ok := <-metach:
		if ok {
			metajson, _ := json.MarshalIndent(meta, "", "  ")
			fmt.Println(string(metajson))
		} else {
			glog.Exit("Not found")
		}
	case <-time.After(30 * time.Second):
		glog.Fatal("Timeout")
	}
}

func main() {
	flag.Set("stderrthreshold", "FATAL")
	flag.Parse()
	switch flag.Arg(0) {
	case "search":
		search(flag.Arg(1))
	default:
		search(flag.Arg(0))
	}
}
