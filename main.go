package main

import (
	"flag"
	"time"

	"github.com/cloudingcity/go-simple-lb/internal/proxy"
	"github.com/cloudingcity/go-simple-lb/internal/utils"
	log "github.com/sirupsen/logrus"
)

var (
	flagURL utils.FlagURL
	port    int
)

func init() {
	flag.Var(&flagURL, "servers", "Use commas to separate")
	flag.IntVar(&port, "port", 8080, "Port to serve")
	flag.Parse()

	if len(flagURL.URLs) == 0 {
		log.Fatal("Please provide one or more servers to load balance")
	}

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 03:04:05",
	})
}

func main() {
	lb := proxy.NewLB()
	for _, serverURL := range flagURL.URLs {
		lb.Add(serverURL)
	}
	go lb.HeathCheck(1 * time.Minute)
	lb.Listen(port)
}
