package main

import (
	"net"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var ingestLatency = prometheus.NewHistogram(prometheus.HistogramOpts{
	Name:    "ingest_latency_ms",
	Help:    "Latency of UDP ingestion",
	Buckets: prometheus.ExponentialBuckets(0.1, 2, 12),
})

func main() {
	prometheus.MustRegister(ingestLatency)

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":2112", nil)
	}()

	addr, _ := net.ResolveUDPAddr("udp", ":9000")
	conn, _ := net.ListenUDP("udp", addr)
	defer conn.Close()

	buffer := make([]byte, 2048)

	for {
		start := time.Now()
		n, _, _ := conn.ReadFromUDP(buffer)
		_ = buffer[:n]
		ingestLatency.Observe(float64(time.Since(start).Milliseconds()))
	}
}
