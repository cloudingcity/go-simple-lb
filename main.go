package main

import (
	"flag"

	"github.com/cloudingcity/go-simple-lb/internal/proxy"
	"github.com/cloudingcity/go-simple-lb/internal/utils"
	log "github.com/sirupsen/logrus"
)

var (
	flagURL utils.FlagURL
	port    int
)

func main() {
	flag.Var(&flagURL, "servers", "Use commas to separate")
	flag.IntVar(&port, "port", 8080, "Port to serve")
	flag.Parse()

	if len(flagURL.URLs) == 0 {
		log.Fatal("Please provide one or more servers to load balance")
	}

	lb := proxy.NewLB()
	for _, serverURL := range flagURL.URLs {
		lb.Add(serverURL)
	}
	lb.Listen(port)
}
