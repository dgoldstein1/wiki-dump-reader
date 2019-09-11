package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"os"
	"sync/atomic"
)

// constants
var (
	nodesVisitedCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "golang",
			Name:      "nodes_visited",
			Help:      "Number of nodes parsed and succesfully added to the graph",
		})

	nodesAddedCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "golang",
			Name:      "nodes_added",
			Help:      "Number of nodes succesfully visited",
		})

	totalNodesAdded = asyncInt(0)
)

// resgisters and serves metrics to HTTP
func ServeMetrics() {
	http.Handle("/metrics", promhttp.Handler())
	// register metrics
	prometheus.MustRegister(nodesVisitedCounter)
	prometheus.MustRegister(nodesAddedCounter)
	// serve http
	go func() {
		logErr("%v", http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("METRICS_PORT")), nil))
	}()
}

// updates prometheus and internal metrics
func UpdateMetrics(numberOfNodesAdded int) {
	// increment number of nodes crawled
	nodesVisitedCounter.Inc()
	// increment number of nodes
	totalNodesAdded.incr(int32(numberOfNodesAdded))
	nodesAddedCounter.Add(float64(numberOfNodesAdded))
}

// increments async int by "n"
func (c *asyncInt) incr(n int32) int32 {
	return atomic.AddInt32((*int32)(c), n)
}

// decrement astnc int
func (c *asyncInt) get() int32 {
	return atomic.LoadInt32((*int32)(c))
}
