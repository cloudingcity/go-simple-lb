package proxy

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudingcity/go-simple-lb/internal/server"
	log "github.com/sirupsen/logrus"
)

type LoadBalancer struct {
	serverPool *server.Pool
}

func NewLB() *LoadBalancer {
	return &LoadBalancer{
		serverPool: server.NewPool(),
	}
}

func (lb *LoadBalancer) Add(u *url.URL) {
	lb.serverPool.Put(server.NewServer(u))
	log.Printf("Configured server: %s", u)
}

func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if next := lb.serverPool.GetNext(); next != nil {
		next.ServeHTTP(w, req)
		return
	}
	http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
}

func (lb *LoadBalancer) HeathCheck(d time.Duration) {
	t := time.NewTicker(d)
	for range t.C {
		log.Println("Health check starting...")
		msgs := lb.serverPool.HealthCheck()
		for _, msg := range msgs {
			log.Warn(msg)
		}
		log.Println("Health check completed")
	}
}

func (lb *LoadBalancer) Listen(port int) {
	addr := fmt.Sprintf(":%d", port)
	log.Printf("Started listening on %s\n", addr)
	if err := http.ListenAndServe(addr, lb); err != nil {
		log.Fatal(err)
	}
}
