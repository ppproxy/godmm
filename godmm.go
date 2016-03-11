package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"time"

	"github.com/golang/glog"
	"github.com/junzh0u/opendmm"
)

func search(query string, dbpath string) {
	db, err := opendmm.NewDB(dbpath)
	if err != nil {
		glog.Fatal(err)
	}
	defer db.Close()
	metach := opendmm.Search(query, db)
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

func crawl(dbpath string) {
	db, err := opendmm.NewDB(dbpath)
	if err != nil {
		glog.Fatal(err)
	}
	opendmm.Crawl(db)
}

func main() {
	flag.Set("stderrthreshold", "FATAL")
	dbpath := flag.String("db", "/tmp/opendmm.boltdb", "path for crawler db")
	flag.Parse()
	switch flag.Arg(0) {
	case "search":
		search(flag.Arg(1), *dbpath)
	case "crawl":
		crawl(*dbpath)
	default:
		search(flag.Arg(0), *dbpath)
	}
}
