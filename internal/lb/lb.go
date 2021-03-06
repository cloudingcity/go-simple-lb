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
	for _, u := range urls {
		log.Printf("Configured server: %s", u)
	}
	lb.controller.SetupServers(urls...)
}

func (lb *LoadBalancer) HeathCheck(d time.Duration) {
	t := time.NewTicker(d)
	for range t.C {
		log.Println("Health check starting...")
		lb.controller.HealthCheck()
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
