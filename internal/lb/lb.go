package lb

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudingcity/go-simple-lb/internal/server"
	log "github.com/sirupsen/logrus"
)

type LoadBalancer struct {
	controller *server.Controller
}

func New() *LoadBalancer {
	return &LoadBalancer{
		controller: server.NewController(),
	}
}

func (lb *LoadBalancer) Register(urls ...*url.URL) {
	var servers []*server.Server

	for _, u := range urls {
		servers = append(servers, server.NewServer(u))
		log.Printf("Configured server: %s", u)
	}
	lb.controller.SetServers(servers)
}

func (lb *LoadBalancer) HeathCheck(d time.Duration) {
	t := time.NewTicker(d)
	for range t.C {
		log.Println("Health check starting...")
		msgs := lb.controller.HealthCheck()
		for _, msg := range msgs {
			log.Warn(msg)
		}
		log.Println("Health check completed")
	}
}

func (lb *LoadBalancer) Listen(port int) {
	addr := fmt.Sprintf(":%d", port)
	log.Printf("Started listening on %s\n", addr)
	if err := http.ListenAndServe(addr, lb.controller.HTTPHandler()); err != nil {
		log.Fatal(err)
	}
}
