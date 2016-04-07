package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"time"

	"github.com/golang/glog"
	"github.com/junzh0u/opendmm"
	"github.com/syndtr/goleveldb/leveldb"
)

func search(query, cachepath string) {
	cache, err := leveldb.OpenFile(cachepath, nil)
	if err != nil {
		glog.Fatal(err)
	}
	defer cache.Close()

	metach := opendmm.Search(query, cache)
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
	movieCachePath := flag.String("moviecache", "/tmp/opendmm.movie.cache", "path to http cache")
	httpCachePath := flag.String("httpcache", "/tmp/opendmm.http.cache", "path to http cache")
	flag.Parse()

	switch flag.Arg(0) {
	case "search":
		search(flag.Arg(1), *movieCachePath)

	case "guess":
		for keyword := range opendmm.Guess(flag.Arg(1)).Iter() {
			fmt.Println(keyword)
		}

	case "crawl":
		opendmm.Crawl(*movieCachePath, *httpCachePath)

	default:
		search(flag.Arg(0), *movieCachePath)
	}
}
